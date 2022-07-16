package buffering

import "fmt"

func Main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	oneChan := make(chan string, 2)
	oneChan <- "1"
	msg := <-oneChan
	fmt.Println(msg)
	oneChan <- "2"
	msg = <-oneChan
	fmt.Println(msg)
}
