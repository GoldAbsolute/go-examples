package arrayssov

import (
	"fmt"
	"math"
)

func Main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	pipbob := [4]string{"hui", "pizda", "mocha", "ujas"}
	pipbob[3] = "stap this"

	var BigBoob [100][5]float64
	for i := 0; i < len(BigBoob); i++ {
		for j := 0; j < len(BigBoob[i]); j++ {
			BigBoob[i][j] = math.Pow(float64(i), float64(j))
		}
	}
	fmt.Println("BigBoob: ", BigBoob)
}
