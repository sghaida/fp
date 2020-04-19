package main

import (
	"fmt"
	"strings"

	"github.com/sghaida/fp/src/apply"
)

func main() {
	// check map function for integers
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	lifted := apply.Lift(data)
	addOne := func(item int) int { return item + 1 }
	newData := lifted.Apply(addOne)
	fmt.Println(newData.Get())
	fmt.Println(data)
	// check map function for strings
	lst := []string{"a", "b", "c"}
	liftedStrs := apply.Lift(lst)
	cap := func(str string) string { return strings.ToUpper(str) }
	res := liftedStrs.Apply(cap)
	fmt.Println(lst)
	fmt.Println(res.Get())
	// map to different type
	makeFloat := func(in int) float64 { return float64(in) * 1.5 }
	res1 := lifted.Apply(makeFloat)
	fmt.Println(res1)

}
