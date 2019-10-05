package http

type CreateCounterRequest struct {
	Name string `json:"name" binding:"required"`
}

type IncrementCounterRequest struct {
	ID string `json:"id" binding:"required"`
}