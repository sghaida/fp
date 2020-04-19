package utils_test

import (
	"testing"

	td "github.com/sghaida/fp/src/testdata"
	"github.com/sghaida/fp/src/utils"
)

func Test_IsEmptyOrNil(t *testing.T) {
	tt := td.CreateVldData()
	for _, testCase := range tt {
		t.Run(testCase.Name, func(t *testing.T) {
			status := utils.IsNilOrEmpty(testCase.Data)
			if status && testCase.HasError == false {
				t.Errorf("expected not empty or nil, got empty or nil")
			}
			if !status && testCase.HasError == true {
				t.Errorf("expected empty or nil to be true, got %v", status)
			}
		})
	}
}
