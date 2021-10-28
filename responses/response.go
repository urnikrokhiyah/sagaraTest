package responses

type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Failed struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SuccessResponse(status, message string, data interface{}) Success {
	response := Success{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return response
}

func FailedResponse(status, message string) Failed {
	response := Failed{
		Status:  status,
		Message: message,
	}
	return response
}
