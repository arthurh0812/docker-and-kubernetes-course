package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var dirname string
var authAPIService string
var tasksDirectory string

func init() {
	authAPIService = os.Getenv("AUTH_API_SERVICE_HOST")
	tasksDirectory = os.Getenv("TASKS_DIRECTORY")
	var err error
	dirname, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tasksFilePath := filepath.Join(dirname, tasksDirectory, "tasks.txt")
	tasksSep := "TASK_SEP"

	app := iris.Default()

	logger := app.Logger()

	app.Use(func(ctx *context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTION")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		ctx.Next()
	})

	app.Get( "/tasks", func(ctx *context.Context) {
		logger.Logf(golog.InfoLevel, "tasksFilePath: %v\n", tasksFilePath)

		_, err := extractAndVerifyToken(ctx.Request().Header)
		if err != nil {
			ctx.StopWithError(500, err)
			return
		}
		bs, err := os.ReadFile(tasksFilePath)
		if err != nil {
			ctx.StopWithError(500, fmt.Errorf("failed to load the tasks: %v", err))
			return
		}
		logger.Logf(golog.InfoLevel, "tasksFilePath bytes: %v\n", bs)
		tasks := bytes.Split(bs, []byte(tasksSep))
		if len(tasks) > 0 {
			tasks = tasks[:len(tasks)-1]
		}
		logger.Logf(golog.InfoLevel, "tasks: %#v\n", tasks)

		goTasks := make([]*Task, 0, len(tasks))
		for _, t := range tasks {
			goTask := &Task{}
			err := json.Unmarshal(t, goTask)
			if err != nil {
				ctx.StopWithError(500, err)
				return
			}
			goTasks = append(goTasks, goTask)
		}
		logger.Logf(golog.InfoLevel, "goTasks: %#v", goTasks)

		ctx.StopWithJSON(http.StatusOK, APIResponse{
			Message: "Tasks loaded.",
			Status: http.StatusOK,
			Count: int64(len(goTasks)),
			Data: iris.Map{
				"tasks": goTasks,
			},
		})
	})

	app.Post( "/tasks", func(ctx *context.Context) {
		newTask := Task{}
		err := ctx.ReadJSON(&newTask)
		if err != nil {
			ctx.StopWithError(500, err)
			return
		}

		f, err := os.OpenFile(tasksFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0750)
		if err != nil {
			ctx.StopWithError(500, err)
			return
		}
		enc := json.NewEncoder(f)
		err = enc.Encode(newTask)
		if err != nil {
			ctx.StopWithError(500, err)
			return
		}
		_, err = f.WriteString(tasksSep) // insert a task separator
		if err != nil {
			ctx.StopWithError(500, err)
			return
		}
		ctx.StopWithJSON(http.StatusCreated, APIResponse{
			Message: "Task stored.",
			Status: http.StatusCreated,
			Count: 1,
			Data: iris.Map{
				"createdTask": &newTask,
			},
		})
	})

	err := app.Run(iris.Addr(":8000"))
	if err != nil {
		log.Fatal(err)
	}
}

func extractAndVerifyToken(header http.Header) (string, error) {
	auth := header.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("no authorization header provided")
	}
	var token string
	if parts := strings.Split(auth, " "); parts[0] != "Bearer" {
		return "", fmt.Errorf("authorization requires %q key", "Bearer")
	} else if len(parts) < 2 {
		return "", fmt.Errorf("no bearer token provided")
	} else {
		token = parts[1]
	}
	log.Printf("URL: http://%s/verify-token/%s\n", authAPIService, token)
	res, err := http.Get(fmt.Sprintf("http://%s/verify-token/%s", authAPIService, token))
	if err != nil {
		return "", fmt.Errorf("request to authorization service failed: %v", err)
	}
	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()
	apiRes := APIResponse{}
	err = dec.Decode(&apiRes)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON API response from the authorization service: %v", err)
	}
	if apiRes.Status % 100 > 2 {
		return "", fmt.Errorf("request to authorization service returned error: %s", apiRes.Message)
	}
	return apiRes.Data["uid"].(string), nil
}