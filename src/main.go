package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"

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
	toCap := func(str string) string { return strings.ToUpper(str) }
	res := liftedStrs.Apply(toCap)
	results := res.Get().([]string)
	fmt.Println(results)

	// map to different type
	makeFloat := func(in int) float64 { return float64(in) * 1.5 }
	res1 := lifted.Apply(makeFloat)
	// res3 := res1.Get().([]float64)
	// fmt.Println(res3)
	for r := range res1.Get().([]float64) {
		fmt.Println(r)
	}

	name := apply.Lift("saddam")
	checkName := func(s string) bool {
		if s == "saddam" {
			return true
		}
		return false
	}
	res2 := name.Apply(checkName)
	fmt.Println(res2.Get())

	number := apply.Lift(1)
	multiply := func(s int) float64 { return float64(s) * 10.3 }
	m := number.Apply(multiply)
	fmt.Println(m.Get().(float64))

	var numbersLst []int
	_ = faker.SetRandomMapAndSliceSize(1000000)
	_ = faker.FakeData(&numbersLst)

	nLstLifted := apply.Lift(numbersLst)
	start := time.Now()
	output := nLstLifted.Apply(multiply)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(output.Get().([]float64)[100])

}
