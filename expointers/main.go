package expointers

import "fmt"

var int1, int2 = 50, 255

func showAdress() {
	fmt.Println("int1: ", &int1, " int2: ", &int2)
}

func Main() {
	showAdress()
	testInt := 5050
	FuckPtr(&testInt)
	fmt.Println("Сама переменная: ", testInt)
	fmt.Println("Адресс в памяти: ", &testInt)
	PtrIntresting(&int1, &int2)

	println("Значение int1 после: ", int1, " Значение int2 после: ", int2, " После изменения")
}
func FuckPtr(ptr *int) {
	*ptr *= 2
}
func PtrIntresting(onePtr, twoPtr *int) {
	println("Первая область памяти: ", &onePtr, onePtr, " Вторая память: ", &twoPtr, twoPtr, " До изменения")
	//varPtr := onePtr
	//*onePtr = *twoPtr
	onePtr = twoPtr // значение для первого поинтера начинает указывать на область памяти второй переменной
	*onePtr = 1377  // изменение хранящегося значения по первому указателю меняет значение второй переменной
	//twoPtr = varPtr
	//*onePtr = 1377
	println("Первая область памяти: ", &onePtr, onePtr, " Вторая память: ", &twoPtr, twoPtr, " После изменения")
}
