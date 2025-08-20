package img_processing

type FailedImg struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ResData struct {
	SavedFiles  []string    `json:"savedFiles"`
	FailedFiles []FailedImg `json:"failedFiles"`
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
