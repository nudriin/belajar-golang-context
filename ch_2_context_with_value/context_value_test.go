package ch_2_context_with_value

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {

	contextA := context.Background() // * membuat context kosong

	contextB := context.WithValue(contextA, "b", "CTX_B") // * membuat child context dengan value
	contextC := context.WithValue(contextA, "c", "CTX_C") // * membuat child context dengan value

	contextD := context.WithValue(contextB, "d", "CTX_D") // * membuat child context dengan value
	contextE := context.WithValue(contextB, "e", "CTX_E") // * membuat child context dengan value

	contextF := context.WithValue(contextC, "f", "CTX_F") // * membuat child context dengan value

	fmt.Println("================ CONTEXT ====================")
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println("================ CONTEXT ====================")
	fmt.Println()
	fmt.Println("================ VALUE ====================")
	fmt.Println("CTX B :", contextB.Value("b")) // * mengambil nilai dari context b
	fmt.Println("CTX D :", contextD.Value("b")) // * mengambil nilai dari context b melalui context d karana childnya
	fmt.Println("CTX D :", contextD.Value("c")) // * tidak dapat karena bukan parentnya
	fmt.Println("CTX D :", contextF.Value("c")) // * dapat karena parentnya

}
