package gopasswordgenerator

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestPasswordGenerator_Generate(t *testing.T) {
	testcases := []struct {
		testcasename               string
		passwordLength             int32
		nExpectedUpperCaseLetters  int32
		nExpectedLowerCaseLetters  int32
		nExpectedSpecialCharacters int32
		nExpectedNumbers           int32
		expectedError              error
	}{
		{
			"TC - 1",
			12,
			4,
			3,
			DefaultSpecialCharactersLength,
			DefaultNumbersLength,
			nil,
		},
		{
			"TC - 2",
			11,
			0,
			0,
			DefaultSpecialCharactersLength,
			DefaultNumbersLength,
			ErrInvalidPasswordLength,
		},
		{
			"TC - 3",
			13,
			4,
			4,
			DefaultSpecialCharactersLength,
			DefaultNumbersLength,
			nil,
		},
		{
			"TC - 4",
			20,
			8,
			7,
			DefaultSpecialCharactersLength,
			DefaultNumbersLength,
			nil,
		},
	}

	for _, tc := range testcases {
		pg := NewPasswordGenerator().
			WithLength(tc.passwordLength)

		//_, validateErr := pg.ValidateRequest()
		//assert.Equal(t, tc.expectedError, validateErr, tc.testcasename)
		//
		//if validateErr != nil {
		//	continue
		//}

		password, actualErr := pg.Generate()
		if errors.Is(actualErr, ErrInvalidPasswordLength) {
			assert.Equal(t, tc.expectedError, actualErr, tc.testcasename)
			continue
		}
		assert.Nil(t, nil, actualErr, tc.testcasename)

		assert.Equal(t, int32(len(password)), tc.passwordLength, tc.testcasename)

		letters := []rune(password)
		nActualUpper := int32(0)
		nActualLower := int32(0)
		nActualDigit := int32(0)
		nActualSpecialChar := int32(0)
		for _, l := range password {
			if unicode.IsUpper(l) {
				nActualUpper = nActualUpper + 1
			}

			if unicode.IsLower(l) {
				nActualLower = nActualLower + 1
			}

			if unicode.IsDigit(l) {
				nActualDigit = nActualDigit + 1
			}

			// TODO :: Need to check
			//if unicode.IsSymbol(l) {
			//	nActualSpecialChar = nActualSpecialChar + 1
			//}

			if strings.ContainsAny(string(specialCharacters), string(l)) {
				nActualSpecialChar = nActualSpecialChar + 1
			}
		}

		assert.Equal(t, true, unicode.IsUpper(letters[0]), fmt.Sprintf("First letter should be upper case: %s", tc.testcasename))
		assert.Equal(t, tc.nExpectedUpperCaseLetters, nActualUpper, fmt.Sprintf("Number of uppercase letters: %s", tc.testcasename))
		assert.Equal(t, tc.nExpectedLowerCaseLetters, nActualLower, fmt.Sprintf("Number of lowercase letters: %s", tc.testcasename))
		assert.Equal(t, tc.nExpectedNumbers, nActualDigit, fmt.Sprintf("Number of numbers: %s", tc.testcasename))
		assert.Equal(t, tc.nExpectedSpecialCharacters, nActualSpecialChar, fmt.Sprintf("Number of special letters: %s", tc.testcasename))
	}
}

func Benchmark_TestHandleRequest(b *testing.B) {
	tc := struct {
		testcasename   string
		passwordLength int32
		expectedError  error
	}{
		"TC - 1",
		12,
		nil,
	}

	for n := 0; n < b.N; n++ {
		pg := NewPasswordGenerator().
			WithLength(tc.passwordLength)

		_, validateErr := pg.ValidateRequest()
		assert.Equal(b, tc.expectedError, validateErr, tc.testcasename)

		password, actualErr := pg.Generate()
		assert.Equal(b, nil, actualErr, tc.testcasename)

		firstLetter := []rune(password)[0]
		assert.Equal(b, true, unicode.IsUpper(firstLetter), tc.testcasename)
	}
}
