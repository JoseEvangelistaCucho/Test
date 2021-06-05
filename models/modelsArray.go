package models

type ArrayNum struct {
	Valor1 [][]int `json:"valor1"`
}

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"string"`
	Response interface{} `json:"interface,omitempty"`
}

type ResponseVista struct {
	ArrayVista [][]int `json:"arrayVista"`
}
