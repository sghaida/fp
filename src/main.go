package main

import (
	"fmt"

	"github.com/sghaida/fp/src/apply"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	lifted := apply.Lift(data)
	addOne := func(item int) int { return item + 1 }
	newData := lifted.Apply(addOne)
	fmt.Println(newData.Get())
	fmt.Println(data)
}
