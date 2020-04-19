package options_test

import (
	"os"
	"testing"

	"github.com/sghaida/fp/src/Monads/options"
	td "github.com/sghaida/fp/src/testdata"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Test_OptionsCreation(t *testing.T) {
	t.Parallel()
	tt := td.CreateOptionsDataTT()
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Data.HasValue() && tc.IsEmpty {
				t.Errorf("expected Nonetype, got %v", tc.Data.Get())
			}
			if !tc.Data.HasValue() && !tc.IsEmpty {
				t.Errorf("expected SomeType, got NoneType")
			}
		})
	}
}

func TestOptional_GetValue(t *testing.T) {
	t.Parallel()
	t.Run("check NoneType", func(t *testing.T) {
		none := options.None()
		if none.HasValue() {
			t.Errorf("expected None Value, got %v", none.Get())
		}
		switch ot := none.Get().(type) {
		case options.NoneValue:
		default:
			t.Errorf("expeced NoneValue, got %v", ot)
		}
	})
	t.Run("check SomeType", func(t *testing.T) {
		some := options.Some("some data")
		if !some.HasValue() {
			t.Errorf("expected None Value, got %v", some.Get())
		}
		switch ot := some.Get().(type) {
		case string:
		default:
			t.Errorf("expeced SomeValue, got %v", ot)
		}
	})

}
