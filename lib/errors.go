package gopasswordgenerator

import (
	"errors"
	"fmt"
)

// ErrInvalidPasswordLength ... thrown when password length specified is less than the default length
var ErrInvalidPasswordLength = errors.New(fmt.Sprintf("invalid password length. should be a minimum of %d characters", DefaultPasswordLength))
