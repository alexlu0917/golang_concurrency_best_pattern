package update_consumer_error_handle_pattern

import "fmt"

type Result struct {
	data int
	err  error
}

func StartPoint() {
	fmt.Println("Updated Consumer error handle pattern")
	input := []int{1, 2, 3, 4}
	resultCh := make(chan Result)

	inputCh := generator(input)
	
	consumer(inputCh, resultCh)
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

func consumer(inputCh chan int, resultCh chan Result) {
	fmt.Println("consumer")
	defer close(resultCh)
	
	for data := range inputCh {
		res, error = loadDataFromDatabase(data)
		result := Result{
			data: res,
			err: error
		}
		resultCh <- data
	}
}
