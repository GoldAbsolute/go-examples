package extimer

import (
	"fmt"
	"time"
)

func Main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// Можешь перенести sleep чтобы таймер успел остановиться
	time.Sleep(2 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}
