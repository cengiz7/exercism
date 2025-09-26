package isogram

import (
	"unicode"
	"strings"
)

func IsIsogram(s string) bool {
	temp := ""
	runes := []rune(s)
	for i,rn := range runes {
		if unicode.IsLetter(rn){
			ch := strings.ToUpper(s[i:i+1])
			if strings.ContainsAny(temp,ch){
				return false
			}
			temp += ch
		}
	}
	return true
}