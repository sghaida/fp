package Monads_test

import (
	"os"
	"testing"

	"github.com/sghaida/fp/src/Monads"
	td "github.com/sghaida/fp/src/testdata"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_OptionsCreation(t *testing.T) {
	t.Parallel()
	tt := td.CreateTTOptionsData()
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Data.HasValue() && tc.IsEmpty {
				t.Errorf("expected Nonetype, got %v", tc.Data.GetValue())
			}
			if !tc.Data.HasValue() && !tc.IsEmpty {
				t.Errorf("expected SomeType, got NoneType")
			}
		})
	}
}

func TestOptional_GetValue(t *testing.T) {
	t.Parallel()
	option := Monads.Optional{}
	t.Run("check NoneType", func(t *testing.T) {
		none := option.None()
		if none.HasValue() {
			t.Errorf("expected None Value, got %v", none.GetValue())
		}
		switch ot := none.GetValue().(type) {
		case Monads.NoneValue:
		default:
			t.Errorf("expeced NoneValue, got %v", ot)
		}
	})
	t.Run("check SomeType", func(t *testing.T) {
		some := option.Some("some data")
		if !some.HasValue() {
			t.Errorf("expected None Value, got %v", some.GetValue())
		}
		switch ot := some.GetValue().(type) {
		case string:
		default:
			t.Errorf("expeced SomeValue, got %v", ot)
		}
	})

}
