package graphql

type Input struct {
	Query string `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

