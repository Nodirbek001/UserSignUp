package utils

import (
	"NewProUser/constants"
	"regexp"
	"unicode"
)

var usbPhoneRegex = regexp.MustCompile(`[+]{1}99{1}[0-9]{10}$`)

func IsPhoenValid(p string) bool {
	return usbPhoneRegex.MatchString(p)
}

func ValidatePassword(p string) error {
	if len(p) < 8 {
		return constants.ErrPasswordTooShort
	}
	if len(p) > 256 {
		return constants.ErrPasswordTooLong
	}
	hasAnyDigit := false
	hasAnyAlphabetic := false
	for _, c := range p {
		if unicode.IsDigit(c) {
			hasAnyDigit = true
		}

		if unicode.IsLetter(c) {
			hasAnyAlphabetic = true
		}
	}

	if !hasAnyDigit {
		return constants.ErrMustContainDigit
	}

	if !hasAnyAlphabetic {
		return constants.ErrMustContainAlphabetic
	}

	return nil
}
