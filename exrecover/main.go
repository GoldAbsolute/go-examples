package exrecover

import "fmt"

func mayPanic() {
	panic("PANIKA!!!")
}
func Main() {

	// закончит панику
	defer func() {
		recover()
	}()
	// закончит панику и выведет сообщение
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("hey, i dont work")
}
