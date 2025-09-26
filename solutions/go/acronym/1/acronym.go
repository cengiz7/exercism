package acronym

import "strings"

func Abbreviate(s string) string {
	acronymText := ""
	s = strings.Replace(s,"-", " ",-1)
	s = strings.Title(s)
	strArray := strings.Split(s," ")
	for _, k := range strArray{
		acronymText += k[0:1]
	}
	return acronymText
}