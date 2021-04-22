package main

type APIResponse struct {
	Status int32 `json:"status"`
	Message string `json:"message"`
	Count int64 `json:"count"`
	Data map[string]interface{} `json:"data"`
}

type Task struct {
	Title string `json:"title"`
	Text string `json:"text"`
}