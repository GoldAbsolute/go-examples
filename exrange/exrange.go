package exrange

import "fmt"

func Main() {
	mySlice := []int{2, 3, 4}
	var sum int
	for _, iterator := range mySlice {
		sum += iterator
	}
	println(sum)

	myTwoSlice := []string{"My", "little", "slice"}
	for index, element := range myTwoSlice {
		if index == 1 {
			println(element)
		}
	}

	key := map[string]string{"a": "hui", "b": "pizda"}
	for key, value := range key {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for _, values := range key {
		print("just a values: ", values, " ")
	}

	for keys := range key {
		print("Just a keys: ", keys, " ")
	}

	for index, rune := range "go" {
		print(" Over strings range its index of rune and rune itself: ", index, " ", rune, " ")
	}
}
