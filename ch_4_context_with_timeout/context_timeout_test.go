package ch_3_context_with_timeout

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
func TestContextWithTimeot(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // * mengitung jumlah goroutine yang sedang running

	parent := context.Background()                            // * membuat parent context
	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // * membuat context dengan timeout, cancel akan otomatis di eksekusi setelah 5 detik
	defer cancel()

	// * Menggunakan context cancel sebagai parameternya
	destinantion := CreateCounter(ctx)
	for n := range destinantion {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)         // * menunggu cancel selesai, karena cancel async
	fmt.Println(runtime.NumGoroutine()) // * mengitung jumlah goroutine yang sedang running
}
