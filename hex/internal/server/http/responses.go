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

type RegisterUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}