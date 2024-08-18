package ch_5_context_with_deadline

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destinantion := make(chan int)

	go func() {
		defer close(destinantion)
		counter := 1
		for {
			select {
			// Done() mereturn channel
			case <-ctx.Done(): // jika ada dikirim sinyal cancel maka akan di return
				return
			default:
				destinantion <- counter
				counter++
				time.Sleep(1 * time.Second) // membuat simulasi lemot
			}
		}
	}()
	return destinantion
}
func TestContextWithDeadline(t *testing.T) {
	// Deadline membatalkan context sesuai dengan jam atau waktu yang sudah ditentukan
	fmt.Println(runtime.NumGoroutine())

	parent := context.Background() // * membuat parent context
	// Menggunakan (time)
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // * membuat context dengan Deadline, cancel akan otomatis di eksekusi setelah 5 detik
	defer cancel()

	// * Menggunakan context cancel sebagai parameternya
	destinantion := CreateCounter(ctx)
	for n := range destinantion {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
