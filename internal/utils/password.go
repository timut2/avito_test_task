package utils

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

var (
	ErrPasswordTooShort      = errors.New("password must be at least 8 characters long")
	ErrPasswordNoDigit       = errors.New("password must contain at least one digit")
	ErrPasswordNoSpecialChar = errors.New("password must contain at least one special character (!@#$%^&*)")
)

func ValidatePassword(password string) error {

	if utf8.RuneCountInString(password) < 8 {
		return ErrPasswordTooShort
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return ErrPasswordNoDigit
	}

	if !regexp.MustCompile(`[!@#$%^&*]`).MatchString(password) {
		return ErrPasswordNoSpecialChar
	}

	return nil
}
