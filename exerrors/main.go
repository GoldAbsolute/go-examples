package exerrors

import (
	"errors"
	"fmt"
)

func f1(i int) (int, error) {
	if i == 42 {
		return -1, errors.New("42 не подходит")
	}
	return i + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(i int) (int, error) {
	if i == 42 {
		return -1, argError{i, "cant work with it"}
	}
	return i + 3, nil
}

func Main() {
	fmt.Println(f2(10))
	fmt.Println(f2(15))
	fmt.Println(f2(42))

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
