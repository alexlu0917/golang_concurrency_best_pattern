package generator_pattern

import "fmt"

func MainFunc() {
	fmt.Println("Generator Pattern")

	input := []int{1, 2, 3, 4, 5}

	inputCh := generator(input)

	consumer(inputCh)
}

func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			inputCh <- data
		}
	}()

	return inputCh
}

func consumer(inputCh chan int) {
	for data := range inputCh {
		fmt.Println(data)
	}
}
