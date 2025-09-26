package twofer

import "fmt"

func ShareWith(s string) string {
	if s == ""{
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", s)
}