package helpers_testData

import (
	"github.com/sghaida/fp/src/Monads"
)

type TTOptionsData struct {
	Name    string
	Data    Monads.Optional
	IsEmpty bool
}

func CreateTTOptionsData() []TTOptionsData {
	return []TTOptionsData{
		{
			Name:    "some string",
			Data:    Monads.Option("this is not empty"),
			IsEmpty: false,
		}, {
			Name:    "some number",
			Data:    Monads.Option(1.32345),
			IsEmpty: false,
		}, {
			Name: "some struct",
			Data: Monads.Option(
				struct {
					Name string
					Age  uint16
				}{"Saddam Abu Ghaida", 40}),
			IsEmpty: false,
		}, {
			Name:    "some primitive slice",
			Data:    Monads.Option([]uint8{1, 2, 3, 4, 5}),
			IsEmpty: false,
		}, {
			Name: "some slice of struct",
			Data: Monads.Option([]struct {
				name string
				age  uint8
			}{
				{name: "a", age: 1},
				{name: "a", age: 2},
				{name: "a", age: 3},
			}),
			IsEmpty: false,
		}, {
			Name:    "None Empty",
			Data:    Monads.Option("test none empty"),
			IsEmpty: false,
		}, {
			Name:    "empty",
			Data:    Monads.Option(nil),
			IsEmpty: true,
		},
	}
}
