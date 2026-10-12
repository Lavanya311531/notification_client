package validation

import "regexp"

func ValidateMobile(number string) bool {

	pattern := `^\+?[1-9]\d{9,14}$`

	result, _ := regexp.MatchString(
		pattern,
		number,
	)

	return result
}
