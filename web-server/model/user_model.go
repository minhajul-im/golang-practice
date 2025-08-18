package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Status bool     `json:"status"`
	Code   int      `json:"code"`
	Errors []string `json:"errors"`
}

type SuccessResponse struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}
