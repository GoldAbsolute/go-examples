package exchannels

import "fmt"

func Main() {
	oneChan := make(chan string)
	//twoChan := make(chan string)
	//threeChan := make(chan string)

	go func() {
		//for i := 0; i < 15; i++ {
		//	oneChan <- string(i)
		//}
		oneChan <- "Good morning"
		oneChan <- "Good evening"
		oneChan <- "Good night"
	}()
	//msg := <-oneChan
	//fmt.Println(msg)

	fmt.Println(<-oneChan)

	//go func() {
	//	msg := <-oneChan
	//	twoChan <- msg
	//}()
	//go func() {
	//	msg := <-twoChan
	//	threeChan <- msg
	//}()
	//messages := <-threeChan
	//fmt.Println(messages)
}
