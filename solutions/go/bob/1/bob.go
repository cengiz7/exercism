package bob

import (
	"strings"
	"unicode"
)

func isAllUpper(txt string) bool{
	for _, k := range txt {
		if unicode.IsLower(k){
			return false
		}
	}
	return true
}

func containsLetter(txt string) bool{
	for _, r := range txt {
		if unicode.IsLetter(r){
			return true
		}
	}
	return false
}


func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var lstChr string= ""
	
	if len(remark) > 1{
		lstChr = string(remark[len(remark)-1])
	}
	
	switch lstChr {

	case "?":
		if isAllUpper(remark) && containsLetter(remark) {
			return "Calm down, I know what I'm doing!"
			break
		}
		return "Sure."

	case "!":
		if isAllUpper(remark){
			return "Whoa, chill out!"
			break
		}
		return "Whatever."

	default:
		if len(remark) == 0{
			return "Fine. Be that way!"
		}
		if isAllUpper(remark) && containsLetter(remark){
			return "Whoa, chill out!"
		} else {
			return "Whatever."
		}

	}
	return ""
}