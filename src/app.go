package main

import (
	"fmt"

	"github.com/sghaida/fp/src/Monads"
)

func main() {
	option := Monads.Option("test")
	res := option.Some("saddam")
	fmt.Println(res.GetValue())

	o := option.Some(10)
	fmt.Println(option.HasValue())
	fmt.Println(o.GetValue())
}
