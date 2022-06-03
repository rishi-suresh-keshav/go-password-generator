package gopasswordgenerator

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

const (
	DefaultPasswordLength          int32 = 12
	DefaultSpecialCharactersLength int32 = 3
	DefaultNumbersLength           int32 = 2
)

var (
	lcaseletters      = []rune("abcdefghijklmnopqrstuvwxyz")
	ucaseletters      = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers           = []rune("0123456789")
	specialCharacters = []rune("!#&*^") // @, /, \, ', "
)

// PasswordGenerator ...
type PasswordGenerator struct {
	length int32
}

// NewPasswordGenerator ...
func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{
		length: DefaultPasswordLength,
	}
}

// WithLength ... supply desired length of password
func (p *PasswordGenerator) WithLength(n int32) *PasswordGenerator {
	p.length = n
	return p
}

// ValidateRequest ... validates the PasswordGenerator request and outputs error if any
func (p *PasswordGenerator) ValidateRequest() (bool, error) {
	if p.length < DefaultPasswordLength {
		return false, ErrInvalidPasswordLength
	}
	return true, nil
}

// Generate ...
// generates string password given the requirements...
// First letter is always Uppercase
// Contains default number of digits which is defined in the constant DefaultNumbersLength
// Contains default special chars which is defined in the constant DefaultSpecialCharactersLength
func (p *PasswordGenerator) Generate() (string, error) {
	_, err := p.ValidateRequest()
	if err != nil {
		return "", err
	}

	t := float64(p.length-DefaultNumbersLength-DefaultSpecialCharactersLength) / 2
	nUpperCaseLetters := int32(math.Ceil(t))
	nLowerCaseLetters := int32(math.Floor(t))

	if (nLowerCaseLetters + nUpperCaseLetters + DefaultNumbersLength + DefaultSpecialCharactersLength) != p.length {
		return "", errors.New("calculated length for characters doesnt match with requested length")
	}

	s := make([]rune, 0, p.length-1)
	for i := int32(0); i < nUpperCaseLetters; i++ {
		s = append(s, ucaseletters[rand.Intn(len(ucaseletters))])
	}

	for i := int32(0); i < nLowerCaseLetters; i++ {
		s = append(s, lcaseletters[rand.Intn(len(lcaseletters))])
	}

	for i := int32(0); i < DefaultNumbersLength; i++ {
		s = append(s, numbers[rand.Intn(len(numbers))])
	}

	for i := int32(0); i < DefaultSpecialCharactersLength; i++ {
		s = append(s, specialCharacters[rand.Intn(len(specialCharacters))])
	}

	password := make([]rune, p.length)
	// always begin with upper case letter
	password[0] = s[0]
	rand.Seed(time.Now().UnixNano())
	temp := make([]rune, len(s)-1)
	temp = s[1:]
	rand.Shuffle(len(temp), func(i, j int) { temp[i], temp[j] = temp[j], temp[i] })
	for i := range password {
		if i == 0 {
			password[i] = ucaseletters[rand.Intn(len(ucaseletters))]
			continue
		}
		password[i] = temp[i-1]
	}

	return string(password), nil
}
