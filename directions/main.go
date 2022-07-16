package directions

import "fmt"

func ping(chanTaker chan<- string, message string) {
	chanTaker <- message
}
func pong(chanTaker chan<- string, chanGiver <-chan string) {
	msg := <-chanGiver
	chanTaker <- msg
}
func Main() {
	oneChan := make(chan string, 1)
	twoChan := make(chan string, 1)
	ping(oneChan, "Channel pipes message")
	pong(twoChan, oneChan)
	fmt.Println(<-twoChan)

	threeChan := make(chan string, 1)
	fourChan := make(chan string, 1)
	threeChan <- "Прямая передача через каналы"
	doubleMSG := <-threeChan
	fourChan <- doubleMSG
	fmt.Println(<-fourChan)
}
