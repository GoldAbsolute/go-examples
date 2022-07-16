package mapex

import (
	"fmt"
	"go/types"
)

func Main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	//	Тут ты остановился в 3:30

	userSlise := make([]string, 5)
	userSlise = append(userSlise, "str")
	userMap := make(map[string]int)
	delete(userMap, "k1")

	var unicfunc = func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Printf("", i)
			fmt.Println()
		case string:
			fmt.Printf("%t", i)
			fmt.Println()
		case bool:
			fmt.Printf("%t", i)
			fmt.Println()
		case types.Slice:
			fmt.Printf("%t", i)
			fmt.Println()
		case types.Array:
			fmt.Printf("%t", i)
			fmt.Println()
		case types.Func:
			fmt.Printf("%t", i)
			fmt.Println()
		case types.Map:
			fmt.Printf("%t", i)
			fmt.Println()

		}
	}
	//var hui = func(i int) (int, int) {
	//	return 1 + i, i * 2
	//}
	unicfunc(3)
}
