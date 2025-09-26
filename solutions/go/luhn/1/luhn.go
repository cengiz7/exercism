package luhn

import (
	"strings"
	"strconv"
)

func Valid(s string) bool {
	doublCheck := 1
	sum := 0
	s = strings.Replace(s, " ", "", -1)

	for i := len(s); i > 0; i-- {
		value := int(s[i-1])

		if value > 57 || value < 48 {
			return false
		}

		digit, _ := strconv.Atoi(string(value))
		if doublCheck%2 == 0 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		doublCheck++
	}
	if sum%10 == 0 && len(s) > 1 {
		return true
	}
	return false
}