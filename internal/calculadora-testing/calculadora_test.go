package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 10
	num2 := 7
	esperado := 3
	resultado := Restar(num1, num2)
	fmt.Println("RESTAR:", "Esperado: ", esperado, " Resultado: ", resultado)
	assert.Equal(t, esperado, resultado)
}
