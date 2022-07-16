package embedding

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func Main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	type y struct {
		co int
	}
	type x struct {
		co int
	}
	type xyz struct {
		y
		x
	}
	abc := xyz{
		y{co: 5},
		x{co: 10},
	}
	fmt.Println(abc.x.co)
}
