package explicit_cancellation_pattern

import "fmt"

func ExplicitCancellationPattern() {
	fmt.Println("Explicit Cancellation Patter")

	inputArray := []int{1, 2, 3, 4, 5}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(inputArray)
	resultCh := add(doneCh, inputCh)

	for data := range resultCh {
		fmt.Println(data)
	}
}

func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for data := range inputCh {
			result := data + 1
			select {
			case <-doneCh:
				return
			case resultCh <- result:
			}
		}
	}()

	return resultCh
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
