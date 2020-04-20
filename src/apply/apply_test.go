package apply_test

import (
	"os"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"

	"github.com/sghaida/fp/src/apply"
)

func TestMain(m *testing.M) {
	code := m.Run()
	_ = faker.SetRandomMapAndSliceSize(100)
	os.Exit(code)
}

func Test_LiftingType(t *testing.T) {
	t.Parallel()
	t.Run("string equality", func(t *testing.T) {
		id := faker.UUIDHyphenated()
		lft := apply.Lift(id)
		orig := lft.Get().(string)
		if id != orig {
			t.Errorf("expected equality to pass, got inequality")
		}
	})
	t.Run("list of strings equality", func(t *testing.T) {
		var lst []string
		_ = faker.FakeData(&lst)
		lft := apply.Lift(lst)
		orig := lft.Get().([]string)
		if lst[1] != orig[1] {
			t.Errorf("expected equality to pass, got inequality")
		}
	})
}

func Test_Apply(t *testing.T) {
	t.Parallel()
	t.Run("apply on int slice should succeed", func(t *testing.T) {
		var lst []int
		_ = faker.FakeData(&lst)
		firstNum := lst[0]
		addOne := func(item int) int {
			return item + 1
		}
		lifted := apply.Lift(lst)
		result := lifted.Apply(addOne)
		number := result.Get().([]int)[0]
		if number != firstNum+1 {
			t.Errorf("expected value=%v, got %v", firstNum+1, number)
		}
	})

	t.Run("apply on string slice should succeed", func(t *testing.T) {
		var lst []string
		_ = faker.FakeData(&lst)
		firstStr := lst[0]
		capitilize := func(item string) string {
			return strings.ToUpper(item)
		}
		lifted := apply.Lift(lst)
		result := lifted.Apply(capitilize)
		resStr := result.Get().([]string)[0]
		if strings.ToUpper(firstStr) != resStr {
			t.Errorf("expected value=%v, got %v", strings.ToUpper(firstStr), resStr)
		}
	})

	t.Run("apply on string should succeed", func(t *testing.T) {
		id := faker.Sentence()
		lft := apply.Lift(id)
		title := func(item string) string {
			return strings.Title(item)
		}
		size := func(item string) int {
			return len(item)
		}
		result := lft.Apply(title).Apply(size)
		orig := result.Get().(int)
		if len(id) != orig {
			t.Errorf("expected value=%v, got %v", len(id), result)
		}
	})
}

func Test_Apply_Fail(t *testing.T) {
	t.Parallel()
	t.Run("apply on string slice should panic", func(t *testing.T) {
		var lst []string
		_ = faker.FakeData(&lst)
		size := func(item string) int {
			return len(item)
		}
		lifted := apply.Lift(lst)

		resStr := func() {
			result := lifted.Apply(size)
			_ = result.Get().([]string)[0]
		}
		assert.Panics(t, resStr, "this should panic")
	})
}
