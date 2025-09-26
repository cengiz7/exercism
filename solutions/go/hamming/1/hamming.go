package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		err := errors.New("DNA size error.")
		return -1, err
	}

	for i:= 0 ; i < len(a); i++ {
		if a[i] != b[i]{
			distance ++
		}
	}
	return distance, nil
}