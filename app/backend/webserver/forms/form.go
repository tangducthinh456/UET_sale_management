package forms

type ResponseError struct {
	Message    string `json:"error"`
	StatusCode int    `json:"status_code"`
}
