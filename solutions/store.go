package solutions

func NewPayloadStore() solutionResult {
	return solutionResult{Payload: []Payload{}}
}

func (c *solutionResult) Add(payload Payload) {
	c.Payload = append(c.Payload, payload)
}
