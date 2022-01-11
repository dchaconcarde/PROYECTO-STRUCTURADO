package calculadora

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirOK(t *testing.T) {
	num1 := 10
	num2 := 5
	esperado := 2
	resultado, err := Dividir(num1, num2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DIVIDIR OK:", "Esperado: ", esperado, " Resultado: ", resultado)
	assert.Equal(t, esperado, resultado)
}

func TestDividirErr(t *testing.T) {
	num1 := 10
	num2 := 0
	var esperado error = errors.New("El denominador no puede ser 0")
	_, err := Dividir(num1, num2)
	fmt.Println("DIVIDIR / 0:", "Esperado: ", esperado, " Resultado: ", err)
	assert.Equal(t, esperado, err)
}
