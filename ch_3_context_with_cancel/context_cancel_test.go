package ch_3_context_with_cancel

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
			}
		}
	}()
	return destinantion
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine()) // * mengitung jumlah goroutine yang sedang running

	parent := context.Background()            // * membuat parent context
	ctx, cancel := context.WithCancel(parent) // * membuat context dengan cancel, returnya ada context itu sendiri dan func cancel

	// * Menggunakan context cancel sebagai parameternya
	destinantion := CreateCounter(ctx)
	for n := range destinantion {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	// * goroutine akan langsung berhenti ketika mendapatkan sinya cancel
	cancel() // * memanggil function cancel yang berfungsi sebagai sinya cancel

	time.Sleep(2 * time.Second)         // * menunggu cancel selesai, karena cancel async
	fmt.Println(runtime.NumGoroutine()) // * mengitung jumlah goroutine yang sedang running
}
