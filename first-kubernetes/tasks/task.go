package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Title string 		`json:"title"`
	Description string 	`json:"description"`
}

type TaskCreator struct {
	wtr http.ResponseWriter
	req *http.Request
	Input TaskCreatorInput
}

type TaskCreatorInput struct {
	Data struct {
		Task *Task `json:"task"`
	} `json:"data"`
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

func (t *TaskCreator) CreateTask() {
	w, req := t.wtr, t.req
	dec := json.NewDecoder(req.Body)
	input := &t.Input
	err := dec.Decode(input)
	if err != nil {
		log.Fatalf("error decoding JSON: %v", err)
	}
	task := input.Data.Task
	err = createTask(task)
	enc := json.NewEncoder(w)
	output := NewTaskCreatorOutput("successfully created new task record", "success")
	if err != nil {
		output = NewTaskCreatorOutput(fmt.Sprintf("%s", err), "fail")
	}
	_ = enc.Encode(output) // discard error
}

func createTask(task *Task) error {
	log.Printf("created task with title:%q and description:%q\n", task.Title, task.Description)
	return nil
}

func GetTasks(w http.ResponseWriter, req *http.Request) {


}

func UpdateTask(w http.ResponseWriter, req *http.Request) {

}