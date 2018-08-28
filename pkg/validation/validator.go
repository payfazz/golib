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
	return IsMatch(SanitizePhone(input), "^+62[0-9]{10,13}$")
}

// IsValidPassword ...
func IsValidPassword(input string) bool {
	return IsMatch(input, "^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.{6,})$")
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
	sanitized := fmt.Sprintf("%s%s", "+62", matches[len(matches)-1])
	return sanitized
}
