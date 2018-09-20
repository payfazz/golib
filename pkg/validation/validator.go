package validation

import (
	"fmt"
	"regexp"
)

// IsAlphabet ...
func IsAlphabet(input string) bool {
	return IsMatch(input, "^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")
}

// IsEmail ...
func IsEmail(input string) bool {
	return IsMatch(input, "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

// IsPhone ...
func IsPhone(input string) bool {
	sanitized := SanitizePhone(input)
	return IsMatch(sanitized, `^\+62[0-9]{10,13}$`)
}

// IsValidPassword ...
func IsValidPassword(input string) bool {
	length, uppercase, lowercase, number := false, false, false, false
	length = IsMatch(input, "^.{8,}$")
	uppercase = IsMatch(input, "[A-Z]+")
	lowercase = IsMatch(input, "[a-z]+")
	number = IsMatch(input, "[0-9]+")
	return length && uppercase && lowercase && number
}

// IsMatch ...
func IsMatch(input, regex string) bool {
	r, errExp := regexp.Compile(regex)
	if errExp != nil {
		panic(errExp)
	}
	if !r.MatchString(input) {
		return false
	}
	return true
}

// SanitizePhone ...
func SanitizePhone(input string) string {
	stripExp, _ := regexp.Compile(`[^0-9]+`)
	stripped := stripExp.ReplaceAllString(input, "")

	preExp, _ := regexp.Compile(`^(\+?62|0)([0-9]*)`)
	matches := preExp.FindStringSubmatch(stripped)
	if len(matches) == 0 {
		return ""
	}
	sanitized := fmt.Sprintf("%s%s", "+62", matches[len(matches)-1])
	return sanitized
}
