package scrabble

import "strings"

func letterPoint(s string) (int){
	if strings.ContainsAny("AEIOULNRST",s){
		return 1
	}else if strings.ContainsAny("DG",s){
		return 2
	}else if strings.ContainsAny("BCMP",s){
		return 3
	}else if strings.ContainsAny("FHVWY",s){
		return 4
	}else if strings.ContainsAny("K",s){
		return 5
	}else if strings.ContainsAny("JX",s){
		return 8
	}else if strings.ContainsAny("QZ",s){
		return 10
	}else if s == " "{
		return 0
	} else {
		return 1
	}
}
func Score(s string) (int) {
	point := 0
	for i := 0; i< len(s); i++ {
		point += letterPoint(strings.ToUpper(s[i:i+1]))
	}
	return  point
}