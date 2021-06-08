package entities

import (
	"github.com/JoseEvangelistaCucho/mod/models"
)

type ArrayEntities models.ArrayNum
type ResponseVista models.ResponseVista
type Response models.Response

func (c *ArrayEntities) Rotate9() (response Response) {
	//Declarar Variables
	var nArray [][]int
	var arrayIn = c.NumList
	var numArray int = len(arrayIn)

	//Recorre los datos ingresados y verifica errores
	for i := range arrayIn {
		var TamArray int = len(arrayIn[i])

		if TamArray != numArray {
			if TamArray > numArray {
				response.ErrorCode = 0
				response.Message = "Cantidad de Digitos Mayor que Arrays (Verifique lo Digitado ejemp [[1,2,3],[4,5,6],[7,8,9]])"
				return response
			}
			if TamArray < numArray {
				response.ErrorCode = 1
				response.Message = "Cantidad de Arrays Mayor que Digitos (Verifique lo Digitado ejemp [[1,2],[3,4]])"
				return response
			}

			response.ErrorCode = 2
			response.Message = "Fatal Error"
			return response

		}

	}
	// Creacion de nueva matriz
	nArray = make([][]int, numArray)
	for i := range nArray {
		//creamos nuevas filas
		nArray[i] = make([]int, numArray)
	}
	//intercambiando los índices [h -> i], [i -> h] desde el final)
	for i := range arrayIn {
		for h := range arrayIn[i] {
			nArray[h][i] = arrayIn[i][numArray-h-1]
		}
	}

	var vista models.ResponseVista
	vista.ArrayView = append(vista.ArrayView, nArray...)
	response.ErrorCode = 0
	response.Message = "Successful Response"
	response.Response = vista

	return response
}
