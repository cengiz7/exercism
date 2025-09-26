package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix  struct {
	ok   bool
	rows [][]int
	cols [][]int
}

func New (s string) (Matrix, error){
	m := Matrix{false,nil,nil}
	var err = errors.New("")
	var val int
	splitrow := strings.Split(s,"\n")
	if len(splitrow) > 1 {
		lenHist  := 0
		for i, v := range splitrow {
			splitcol := strings.Split(v," ")
			if i == 0 { lenHist = len(splitcol) }
			if len(splitcol) > 1 && len(splitcol) == lenHist {
				for k, v := range splitcol {
					val, err = strconv.Atoi(v)
					if err != nil{ return m , err}
					if !m.Set(i,k,val) {
						return m, err
					}
				}
			} else { return m , err }
		}
		m.ok = true
		return m , nil
	} else {
		return m , err
	}
}

func (m *Matrix) Set (r int, c int , val int) bool {
	if r < len(m.rows) {
		if c < len(m.cols) {
			m.rows[r][c] = val
			m.cols[c][r] = val
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (m *Matrix) Cols () [][]int {
	return m.cols
}

func (m *Matrix) Rows () [][]int {
	return m.rows
}
