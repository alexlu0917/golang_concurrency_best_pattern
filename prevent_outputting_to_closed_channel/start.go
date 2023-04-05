package prevent_outputting_to_closed_channel

import "fmt"

func ErrorProneOutputtingToClosedChannel() {
	fmt.Println("Error prone example")
	ch := make(chan int)
	close(ch)

	//it is ok to read data from closed channel.
	read, open := <-ch

	fmt.Println("read =>", read) //elemet's zero-type value
	fmt.Println("open =>", open) // indicating if the channel is open

	//Writing to the closed channel is panic.
	ch <- 1
}

func PreventClosedChannelOutputting() {
	fmt.Println("Pattern for preventing error from sending data to the closed channel.")
	channel := make(chan int)

	go sender(channel)

	recipient(channel)
}

func sender(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}

	fmt.Println("Sender closes a channel after sending data.")
	close(channel)
}

func recipient(channel chan int) {
	fmt.Println("Recipient present data into console.")
	for data := range channel {
		fmt.Println(data)
	}
}
