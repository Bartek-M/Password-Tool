package validate

import "unicode"

func Validate(passw string) int {
	score := 0

	if len(passw) >= 16 {
		score += 2
	} else if len(passw) >= 8 {
		score += 1
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range passw {
		if unicode.IsNumber(char) {
			hasDigit = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else {
			hasSpecial = true
		}
	}

	if hasUpper {
		score += 2
	}
	if hasLower {
		score += 2
	}
	if hasDigit {
		score += 2
	}
	if hasSpecial {
		score += 2
	}

	return score
}
