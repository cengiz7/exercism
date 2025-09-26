package tournament

import (
	"io"
	"strings"
	"errors"
	"fmt"
)

type result struct {
	team string
	MP int
	W int
	D int
	L int
	P int
}
var results []result

func Tally (reader io.Reader, buffer io.Writer) error{
	resultStr := "Team                           | MP |  W |  D |  L |  P\n"
	results = nil
	var iscontainresult bool
	buf := make([]byte,350)
	i, err := reader.Read(buf)
	if err != nil {
		return err
	}
	lines := strings.Split(string(buf[:i]),"\n")
	for _,line := range lines {
		if len(line) == 0 || line[:1] == "#"{
			continue
		}
		temp := strings.Split(line,";")
		if len(temp) < 3  {
			return  errors.New("Bad input format")
		}
		iscontainresult = strings.Contains(temp[2],"win") || strings.Contains(temp[2],"loss") || strings.Contains(temp[2],"draw")
		if !iscontainresult{
			return  errors.New("Bad input format")
		}
		arrangeTable(temp[0],temp[1],temp[2])
	}
	SortResult()
	for _,b := range results{
		resultStr += fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n",b.team,b.MP,b.W,b.D,b.L,b.P)
	}
	bufferBytes := make([]byte,len(resultStr))
	copy(bufferBytes[:],resultStr)
	buffer.Write(bufferBytes)
	return  nil
}

func swap(m , n int){
	c := results[m]
	results[m] = results[n]
	results[n] = c
}

func SortResult(){
	n := len(results)
	for n!=0{
		newn := 0
		for i:=1;i <= n-1;i++ {
			if results[i-1].P < results[i].P{
				swap(i-1,i)
				newn = i
			}
		}
		n = newn
	}
	n = len(results)
	for n!=0{
		newn := 0
		for i:=1;i <= n-1;i++ {
			if results[i-1].P == results[i].P {
				if results[i-1].team[:1] > results[i].team[:1] {
					swap(i-1,i)
					newn = i
				}else if results[i-1].team[:1] == results[i].team[:1]{
					for k := 0;k<len(results[i-1].team);k++{
						if results[i-1].team[k:k+1] > results[i].team[k:k+1]{
							swap(i-1,i)
							newn = i
						}
					}
				}
			}

		}
		n = newn
	}
}

func arrangeStruct(i int,teamName string, rslt string){
	c := result{teamName,1,0,0,0,0}
	if rslt == "win"{
		c.W++
		c.P+= 3
	}else if rslt == "draw" {
		c.D++
		c.P++
	}else {
		c.L++
	}
	if i < 0 { // if team not exist
		results = append(results,c)
		return
	}
	results[i].MP += c.MP
	results[i].W += c.W
	results[i].D += c.D
	results[i].L += c.L
	results[i].P += c.P
}

func arrangeTable (team1 string, team2 string,rslt string){
	index1 := teamexist(team1)
	index2 := teamexist(team2)
	arrangeStruct(index1,team1,rslt)
	if rslt == "win"{
		arrangeStruct(index2,team2,"loss")
	}else if rslt == "loss"{
		arrangeStruct(index2,team2,"win")
	}else {
		arrangeStruct(index2,team2,"draw")
	}
}

func teamexist(name string) (int){
	for index,res := range results{
		if res.team == name{
			return index
		}
	}
	return -1
}