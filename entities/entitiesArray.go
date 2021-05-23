package entities

import (
	"github.com/JoseEvangelistaCucho/mod/models"
)

type ArrayEntities models.ArrayNum

func (c *ArrayEntities) Rotate9() (a [][]int) {

	var m = c.Valor1
	var n int = len(m)

	a = make([][]int, n) // Crea la nueva "matriz"
	for i := range a {
		a[i] = make([]int, n) // y sus filas
	}
	//Intercambia las filas a columnas (intercambiando los índices j -> i, i -> j desde el final)
	for i := range m {
		for j := range m[i] {
			a[j][i] = m[i][n-j-1]
		}
	}
	return
}
