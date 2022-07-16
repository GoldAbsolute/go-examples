package recursion

import "fmt"

func recurFunc(n int) int {
	if n == 1 {
		return 1
	}
	return n * recurFunc(n-1)
}

func Main() {
	fmt.Println("Результат рекурсии: ", recurFunc(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("Результат рекурсивного замыкания: ", fib(3))
}

// Замыкания тоже могут быть рекурсивны, но должны быть обьявлены через var перед обьявлением/описанием
