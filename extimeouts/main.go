package extimeouts

import (
	"fmt"
	"time"
)

func Main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		chan1 <- "hop hop"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		chan2 <- "Tyan 2, i love you<3"
	}()
	select {
	case one := <-chan1:
		fmt.Println(one)
	case <-time.After(time.Second * 3):
		fmt.Println("1: Time out")
	}
	select {
	case <-time.After(time.Second * 4):
		fmt.Println("2: Time out")
	case two := <-chan2:
		fmt.Println(two)
	}
}
