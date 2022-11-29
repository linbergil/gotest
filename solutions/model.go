package solutions

type UserName string

type TaskName string

type Result int

type Payload string

type solutionResult struct {
	Payload []Payload `json:"payload"`
	Results []Result  `json:"results"`
}

type Solution struct {
	UserName UserName       `json:"user_name"`
	TaskName TaskName       `json:"task_name"`
	Results  solutionResult `json:"results"`
}
