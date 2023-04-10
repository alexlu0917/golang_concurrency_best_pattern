package main

import (
	"concurrency-pattern/generator_pattern"
	"concurrency-pattern/prevent_outputting_to_closed_channel"
	update_consumer_error_handle_pattern "concurrency-pattern/updated_consumer_error_handle_pattern"
	"fmt"
)

func main() {
	fmt.Println("Golang concurrency best practices!!!")

	//Error prone example
	// prevent_outputting_to_closed_channel.ErrorProneOutputtingToClosedChannel()

	//Preventing error from sending data to the closed channel
	prevent_outputting_to_closed_channel.PreventClosedChannelOutputting()

	generator_pattern.MainFunc()

	update_consumer_error_handle_pattern.StartPoint()
}
