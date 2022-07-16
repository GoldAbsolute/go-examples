package variadic

import "fmt"

var sum int

func Main() {
	final := CalcSum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("Final sum from variadic: ", final)
	finalSlice := []int{98, 65, 32, 10, 21}
	fmt.Println("Final sum fron variadic, args = slice: ", CalcSum(finalSlice...))
}

func CalcSum(nums ...int) int {
	for _, value := range nums {
		sum += value
	}
	return sum
}
