package robotname

import (
	"math/rand"
	"time"

)

type Robot struct {
	name string
}
var usedNames []string
var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const integerBytes = "0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = RandStringBytesMaskImprSrc()
	}
	return r.name
}

func (r *Robot) Reset()  {
	r.name = RandStringBytesMaskImprSrc()
}

  //fast rand string generating (founded from internet and arranged)
func RandStringBytesMaskImprSrc() string {
	size:= 5
	b := make([]byte, size)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := size-1, src.Int63(), letterIdxMax; i >= 0; {
		if i < size && i> size-4 {
			if idx := int(cache & letterIdxMask); idx < len(integerBytes) {
				b[i] = integerBytes[idx]
				i--
			}
		}else {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				b[i] = letterBytes[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}

		cache >>= letterIdxBits
		remain--
	}
	if contains(usedNames,string(b)){
		return RandStringBytesMaskImprSrc()
	}
	usedNames = append(usedNames,string(b))
	return string(b)
}

func contains(s []string,e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}