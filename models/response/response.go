package response

type JSONResponse struct {
	ResponseCode    string      `json:"http_code"`
	ResponseMessage string      `json:"message"`
	Data            interface{} `json:"data" binding:"omitempty"`
}

const (
	FailedHttpCode  = "400"
	SuccessHttpCode = "200"

	SuccessHttpStatus = "SUCCESS"
	FailedHttpStatus  = "FAILED"
)
