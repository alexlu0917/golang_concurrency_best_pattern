package update_consumer_error_handle_pattern

import "fmt"

type Result struct {
	data int
	err  error
}

func StartPoint() {
	fmt.Println("Updated Consumer error handle pattern")
	// input := []int{1, 2, 3, 4}
	// resultCh := make(chan Result)

	// inputCh := generator(input)
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
