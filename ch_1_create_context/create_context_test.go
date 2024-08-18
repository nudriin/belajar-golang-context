package ch_1_create_context

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateContext(t *testing.T) {
	background := context.Background() // * Membuat context kosong
	fmt.Println(background)

	todo := context.TODO() // * Membuat context kosong
	fmt.Println(todo)
}
