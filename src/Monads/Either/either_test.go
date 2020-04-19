package Either_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sghaida/fp/src/Monads/Either"
	td "github.com/sghaida/fp/src/testdata"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
func Test_EitherCreation(t *testing.T) {
	t.Parallel()
	either := Either.Either()
	t.Run("empty Either should fail on get", func(t *testing.T) {
		_, err := either.Get()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("nil left should panic", func(t *testing.T) {
		nilLeft1 := func() {
			_, _ = either.Left(nil)
		}
		assert.Panics(t, nilLeft1, "this should panic")
		nilLeft2 := func() {
			Either.Left(nil)
		}
		assert.Panics(t, nilLeft2, "this should panic")
	})

	t.Run("set left after setting right should fail", func(t *testing.T) {
		right, err := either.Right("some data")
		if err == nil {
			// this should fail
			_, err := right.Left(errors.New("some error"))
			if err == nil {
				t.Errorf("expect setting left, got %v", err)
			}
		}
	})

	t.Run("set right after setting left should fail", func(t *testing.T) {
		right, err := either.Left(errors.New("some error"))
		if err == nil {
			// this should fail
			_, err := right.Right("some data")
			if err == nil {
				t.Errorf("expect setting left, got %v", err)
			}
		}
	})

	t.Run("set right using Either Factory should succeed", func(t *testing.T) {
		left := Either.Left(errors.New("some error"))
		_, err := left.Get()
		if err == nil || err.Error() != "some error" {
			t.Errorf("expect setting left, got %v", err)
		}
	})

	tt := td.CreateEitherDataTT()
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Data.IsLeft() {
				data, err := tc.Data.Get()
				if err == nil {
					t.Errorf("expect error, got %v", data)
				}
			}
			if tc.Data.IsRight() {
				_, err := tc.Data.Get()
				if err != nil && tc.HasError {
					t.Errorf("expected some data, got error: %v", err)
				}
			}

		})
	}
}
