package explicit_cancellation_pattern

import "fmt"

func ExplicitCancellationPattern() {
	fmt.Println("Explicit Cancellation Patter")
}

func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for data := range input {
			inputCh <- data
		}
	}()

	return inputCh
}
