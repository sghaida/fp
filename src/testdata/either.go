package helpers_testData

import (
	"errors"

	"github.com/sghaida/fp/src/Either"
)

type EitherDataTT struct {
	Name     string
	Data     Either.Type
	HasError bool
}

func CreateEitherDataTT() []EitherDataTT {
	either := Either.Either()
	emptyRight, _ := either.Right(nil)
	data := []EitherDataTT{
		{
			Name:     "nil right should pass",
			Data:     emptyRight,
			HasError: false,
		}, {
			Name:     "right string should pass",
			Data:     Either.Right("this is not empty"),
			HasError: false,
		}, {
			Name:     "right number should pass",
			Data:     Either.Right(1.32345),
			HasError: false,
		}, {
			Name: "right struct should pass",
			Data: Either.Right(
				struct {
					Name string
					Age  uint16
				}{"Saddam Abu Ghaida", 40}),
			HasError: false,
		}, {
			Name:     "some primitive slice should pass",
			Data:     Either.Right([]uint8{1, 2, 3, 4, 5}),
			HasError: false,
		}, {
			Name: "right slice of struct should pass",
			Data: Either.Right([]struct {
				name string
				age  uint8
			}{
				{name: "a", age: 1},
				{name: "a", age: 2},
				{name: "a", age: 3},
			}),
			HasError: false,
		}, {
			Name:     "right Empty should pass",
			Data:     Either.Right("test none empty"),
			HasError: false,
		}, {
			Name:     "left None empty should pass",
			Data:     Either.Left(errors.New("some error")),
			HasError: false,
		},
	}
	return data
}
