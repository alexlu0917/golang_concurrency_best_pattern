package main

import (
	"concurrency-pattern/prevent_outputting_to_closed_channel"
	"fmt"
)

func main() {
	fmt.Println("Golang concurrency best practices!!!")

	//Error prone example
	// prevent_outputting_to_closed_channel.ErrorProneOutputtingToClosedChannel()

	//Preventing error from sending data to the closed channel
	prevent_outputting_to_closed_channel.PreventClosedChannelOutputting()
}
