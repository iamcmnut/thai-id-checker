package thaichecker_test

import (
	"testing"

	"github.com/iamcmnut/thaichecker"

	"github.com/stretchr/testify/assert"
)

type TestValidateCase struct {
	ThaiID        string
	ExpectedValue bool
}

func TestValidate(t *testing.T) {

	testCases := []TestValidateCase{
		TestValidateCase{
			ThaiID:        "1234567890121",
			ExpectedValue: true,
		},
		TestValidateCase{
			ThaiID:        "123456789012",
			ExpectedValue: false,
		},
		TestValidateCase{
			ThaiID:        "12345678901211",
			ExpectedValue: false,
		},
		TestValidateCase{
			ThaiID:        "9999999999994",
			ExpectedValue: true,
		},
	}

	for _, v := range testCases {
		res := thaichecker.Validate(v.ThaiID)
		assert.Equal(t, v.ExpectedValue, res, "Value must be equal")
	}
}
