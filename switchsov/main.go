package switchsov

import (
	"fmt"
	"time"
)

func Main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	MyPractice := func(i int) {
		if i%2 == 0 {
			fmt.Print("i = ", i)
			fmt.Printf(",i is odd(maybe) number \n")
		} else {
			fmt.Println(i, "is even number or not number type")
		}
	}
	MyPractice(5)
	MyPractice(2)
	MyPractice(4)

	nasd := 123
	switch nasd {
	case 1:
		fmt.Println("nasd = 1")
	case 2:
		fmt.Println("nasd = 2")
	case 123:
		fmt.Println("nasd = 123")
	}
	stro4ka := "custom string"
	lenstr := len(stro4ka)
	switch lenstr {
	case 10:
		fmt.Println("Длина Строчки меньше 10")
	case 9:
		fmt.Println("Длина Строчки больше 10")
	default:
		fmt.Println("Хуевая строчка")

	}
	opredilitel := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		}
	}
	opredilitel("Мда пиздц какой то")

}
