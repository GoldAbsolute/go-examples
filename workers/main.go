package workers

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "start ", j, "job")
		time.Sleep(1 * time.Second)
		fmt.Println("Worker", id, "end ", j, "job")
		results <- j * 2

	}
}

func Main() {
	arr := make([]int, 0)
	const jobsInt = 5
	jobs := make(chan int, jobsInt)
	results := make(chan int, jobsInt)

	for w := 1; w < 10; w++ {
		go worker(w, jobs, results)
	}
	//go worker(1, jobs, results)
	//go worker(2, jobs, results)
	//go worker(3, jobs, results)
	//go worker(4, jobs, results)
	//go worker(5, jobs, results)
	//go worker(6, jobs, results)
	//go worker(7, jobs, results)
	for i := 1; i <= jobsInt; i++ {
		jobs <- i
	}
	close(jobs)
	fmt.Println("jobs closed")

	for i := 1; i <= jobsInt; i++ {
		//a := <-results
		//arr = append(arr, a)
		// Это гарантирует что работа программы завершится только после завершения работы ассинхронных го рутин(воркеров)
		<-results
	}
	//time.Sleep(2 * time.Second)
	fmt.Println(arr)
}
