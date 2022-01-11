package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortNumbers(t *testing.T) {
	var num []int
	num = append(num, 10, 2, 3, 9, 8, 4, 5, 1, 7, 6)
	var esperado []int
	esperado = append(esperado, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	resultado := SortNumbers(num)
	fmt.Println("SORT NUMBERS:", "Esperado: ", esperado, " Resultado: ", resultado)
	assert.Equal(t, esperado, resultado)
}
