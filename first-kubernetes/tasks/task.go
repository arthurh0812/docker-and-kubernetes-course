package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"log"
	"net/http"
)

type Task struct {
	Title string 		`json:"title" validate:"required,max=40"`
	Description string 	`json:"description" validate:"max=200"`
}

type TaskCreator struct {
	wtr http.ResponseWriter
	req *http.Request
	Input TaskCreatorInput
}

type TaskCreatorInput struct {
	Data struct {
		Task *Task `json:"task" validate:"required"`
	} `json:"data" validate:"required"`
}

func (t *TaskCreatorInput) validate() error {
	val := validator.New()
	return val.Struct(t.Data)
}

type TaskCreatorOutput struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func NewTaskCreatorOutput(msg, status string) *TaskCreatorOutput {
	return &TaskCreatorOutput{
		Message: msg,
		Status: status,
	}
}

func NewTaskCreator(w http.ResponseWriter, req *http.Request) *TaskCreator {
	return &TaskCreator{
		wtr: w,
		req: req,
	}
}

func RenderJSON(w http.ResponseWriter, msg, status string) {
	enc := json.NewEncoder(w)
	output := NewTaskCreatorOutput(msg, status)
	_ = enc.Encode(output)
}

func (t *TaskCreator) CreateTask() {
	w, req := t.wtr, t.req
	dec := json.NewDecoder(req.Body)
	input := &t.Input
	err := dec.Decode(input)
	if err != nil {
		log.Fatalf("error decoding JSON: %v", err)
	}
	err = input.validate()
	if err != nil {
		RenderJSON(w, fmt.Sprintf("%s", err), "fail")
	}
	task := input.Data.Task
	err = createTask(task)
	if err != nil {
		RenderJSON(w, fmt.Sprintf("error while creating task: %s", err), "fail")
	}
	RenderJSON(w, "successfully created new task record", "success")
}

func createTask(task *Task) error {
	log.Printf("created task with title:%q and description:%q\n", task.Title, task.Description)
	return nil
}

func GetTasks(w http.ResponseWriter, req *http.Request) {


}

func UpdateTask(w http.ResponseWriter, req *http.Request) {

}