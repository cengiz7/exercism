package twelve

import "fmt"

type DayNsong struct {
	day string
	line string
}

var arr = []DayNsong{
	{"first", ", and a Partridge in a Pear Tree." },{"second",", two Turtle Doves"  },
	{"third", ", three French Hens" },{"fourth", ", four Calling Birds" },
	{"fifth", ", five Gold Rings" },{"sixth", ", six Geese-a-Laying" },
	{"seventh",", seven Swans-a-Swimming"  },{"eighth", ", eight Maids-a-Milking" },
	{"ninth",", nine Ladies Dancing"},{"tenth", ", ten Lords-a-Leaping" },
	{"eleventh", ", eleven Pipers Piping" }, {"twelfth",", twelve Drummers Drumming"}}
func Song() string {
	song := ""
	for i,bb :=range arr{
		song += fmt.Sprintf("On the %s day of Christmas my true love gave to me" ,bb.day)
		k := i
		for ;k>=0;k--{
			if i == 0 { song += ", a Partridge in a Pear Tree."
			}else{
				song += arr[k].line
			}
		}
		song += "\n"
	}
	return song
}

func Verse(i int)string{
	song := fmt.Sprintf("On the %s day of Christmas my true love gave to me" ,arr[i-1].day)
	if i == 1{ song += ", a Partridge in a Pear Tree."
	}else {
		for ;i>=1;i--{
			song+= arr[i-1].line
		}
	}
return song
}