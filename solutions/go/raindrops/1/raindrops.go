package raindrops

import "strconv"

func Convert(value int) string {
	output := ""
	if value % 3 == 0{
		output += "Pling"
	}
	if value % 5 == 0{
		output += "Plang"
	}
	if value % 7 == 0{
		output += "Plong"
	}
	if output == "" {
		return strconv.Itoa(value)
	}
	return output
}