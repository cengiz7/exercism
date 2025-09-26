package grains

import (
	"errors"
	"math"
)
func Total() (uint64){
	return 18446744073709551615
}
func Square(i int) (uint64,error) {
	if i < 1 || i > 64 {
		err := errors.New("Out of table")
		return 0,err
	}
	return uint64(math.Pow(2,float64(i-1))), nil

}