package reverse

func String(s string) string{
	newone := ""
	runes := []rune(s)
	for i:= len(runes); i>0; i-- {
		newone += string(runes[i-1])
	}
	return newone
}