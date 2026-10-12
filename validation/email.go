package validation

import "regexp"

func ValidateEmail(email string) bool {

	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	result, _ := regexp.MatchString(
		pattern,
		email,
	)

	return result
}
