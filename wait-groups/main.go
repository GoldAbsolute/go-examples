package wait_groups

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func Main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Обязательное условие если хочешь чтобы снный код дико не обогнал ассинхронный
		// или хуй знает, может и не убежит,
		// defer заставляет ждать пока не понизиться число горутин в WG на 1
		// Оператор defer откладывает выполнение функции до возврата из окружающей функции.
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
