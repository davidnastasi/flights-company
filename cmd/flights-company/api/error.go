package api


type apiError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
}