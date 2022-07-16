package exselect

import (
	"fmt"
	"time"
)

func Main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)

		}
	}

	chan1 := make(chan int, 1)
	//chan2 := make(chan int)
	//chan3 := make(chan int)
	//chan4 := make(chan int)
	//chan5 := make(chan int)
	go func() {
		chan1 <- 1
	}()
	fmt.Println("Длина chan1: ", len(chan1))
	go func() {
		if len(chan1) == 1 {
			fmt.Println("Хуя, напечатается ли эта строка - дикий рандом")
			chan1 <- int(byte(255))
		}
	}()

	fmt.Println("Длина chan1: ", len(chan1))
	for i := 0; i < 1; i++ {
		select {
		case message1 := <-chan1:
			fmt.Println("Передача между каналами работает исправно!", message1)
		}
	}
}
