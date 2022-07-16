package closing_channels

import "fmt"

func Main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			fmt.Println("more: ", more)
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	chan1 := make(chan int, 10)
	for i := 3; i < 8; i++ {
		chan1 <- i
	}
	close(chan1)
	fmt.Println("chan1 CLOSED")
	for i := 0; i < 10; i++ {
		fmt.Println("chan1 elem: ", <-chan1)
	}
	//fmt.Println("chan1: ", <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1, <-chan1)
}
