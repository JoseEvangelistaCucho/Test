package entities

import (
	"github.com/JoseEvangelistaCucho/mod/models"
)

type ArrayEntities models.ArrayNum
type ResponseVista models.ResponseVista
type Response models.Response

func (c *ArrayEntities) Rotate9() (r Response) {
	var a [][]int
	var m = c.Valor1
	var n int = len(m)

	//Recorre  los datos ingresados
	for i := range m {
		var TamArray int = len(m[i])

		if TamArray != n {
			if TamArray > n {
				r.Code = 1
				r.Message = "Cantidad de Digitos Mayor que Arrays"
				return r
			}
			if TamArray < n {
				r.Code = 2
				r.Message = "Cantidad de Arrays Mayor que Digitos"
				return r
			}
			r.Code = 3
			r.Message = "Fatal Error"
			return r

		}

	}

	a = make([][]int, n) // Crea la nueva "matriz"///
	for i := range a {
		a[i] = make([]int, n) // y sus filas
	}
	//Intercambia las filas a columnas (intercambiando los índices j -> i, i -> j desde el final)
	for i := range m {
		for j := range m[i] {
			a[j][i] = m[i][n-j-1]
		}
	}

	var vista models.ResponseVista
	vista.ArrayVista = append(vista.ArrayVista, a...)
	r.Code = 0
	r.Message = "Respuesta Con Exito"
	r.Response = vista

	return r
}
