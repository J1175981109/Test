package main

import (
	"fmt"
	"math/rand"
)

func main() {

	type food struct {
		number int
		name string
		cost int
	}

	xm := &food{
	1,
	"小面",
	6,
}
	ft := &food{
		2,
		"饭团",
		7,
	}

	xghj := &food{
		3,
		"香菇滑鸡",
		12,
	}

	xcr := &food{
		4,
		"小炒肉",
		15,
	}

	hmj := &food{
		5,
		"黄焖鸡",
		16,
	}

	mc := &food{
		6,
		"冒菜",
		18,
	}

	rand.Seed(2)

	for i := 0; i <= 2; i++{
		j := rand.Intn(6)
		fmt.Printf("food.name", food.number == j)
	}

}

func (food *food)number()int{
	return food.number
}