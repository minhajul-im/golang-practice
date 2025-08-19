package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

type SigninToken struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type Signout struct {
	Message string `json:"message"`
}
