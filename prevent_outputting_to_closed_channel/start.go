package prevent_outputting_to_closed_channel

import "fmt"

func ErrorProneOutputtingToClosedChannel() {
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
	fmt.Println("Test")
}
