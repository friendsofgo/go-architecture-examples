package http

type CreateCounterRequest struct {
	Name string `json:"name" binding:"required"`
}

type IncrementCounterRequest struct {
	ID string `json:"id" binding:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}