package models

type ArrayNum struct {
	NumList [][]int `json:"numList"`
}

type Response struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Response  interface{} `json:"interface,omitempty"`
}

type ResponseVista struct {
	ArrayView [][]int `json:"arrayView"`
}
