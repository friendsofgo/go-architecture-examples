package http

type CreateCounterResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value uint   `json:"value"`
}

type GetCounterResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value uint   `json:"value"`
}

