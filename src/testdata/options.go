package helpers_testData

import (
	"github.com/sghaida/fp/src/Monads/options"
)

type OptionsDataTT struct {
	Name    string
	Data    options.Optional
	IsEmpty bool
}

func CreateOptionsDataTT() []OptionsDataTT {
	return []OptionsDataTT{
		{
			Name:    "some string",
			Data:    options.Option("this is not empty"),
			IsEmpty: false,
		}, {
			Name:    "some number",
			Data:    options.Option(1.32345),
			IsEmpty: false,
		}, {
			Name: "some struct",
			Data: options.Option(
				struct {
					Name string
					Age  uint16
				}{"Saddam Abu Ghaida", 40}),
			IsEmpty: false,
		}, {
			Name:    "some primitive slice",
			Data:    options.Option([]uint8{1, 2, 3, 4, 5}),
			IsEmpty: false,
		}, {
			Name: "some slice of struct",
			Data: options.Option([]struct {
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
			Data:    options.Option("test none empty"),
			IsEmpty: false,
		}, {
			Name:    "empty",
			Data:    options.Option(nil),
			IsEmpty: true,
		},
	}
}
