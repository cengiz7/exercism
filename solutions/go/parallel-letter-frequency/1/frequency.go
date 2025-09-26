package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m1 := FreqMap{}
	for _, r := range s {
		m1[r]++
	}
	return m1
}
func worker(jobs <-chan string, results chan<- FreqMap ) {
	for j := range jobs {
		m1 := FreqMap{}
		for _, r := range j {
			m1[r]++
		}
		results <- m1

	}
}

func ConcurrentFrequency(runeSlice []string) FreqMap{
	size := len(runeSlice)
	jobs := make(chan string, size)
	results := make(chan FreqMap, size)
	m1 := FreqMap{}

	for i :=0 ; i< size ; i ++ {
		go worker(jobs, results)
	}

	for i:= 0 ; i< size ; i++ {
		jobs <- runeSlice[i]
	}
	close(jobs)

	m1Temp := FreqMap{}
	for j := 0; j< size ; j ++ {
		m1Temp = <-results
		for aa,bb := range m1Temp{
			m1[aa] += bb
		}
	}
	return m1
}
























/*package letter

import (
	"sync"
	"fmt"
	"strconv"
)

type FreqMap map[rune]int
var mux sync.Mutex
var wg sync.WaitGroup

func Frequency(s string) FreqMap {
	m1 := FreqMap{}
	for _, r := range s {
		m1[r]++
	}
	return m1
}

func ConcurrentResponse(s string,c chan FreqMap){
	m2 := FreqMap{}
	for _, r := range s {
		//mux.Lock()
		m2[r]++
		//mux.Unlock()
	}
	c <-m2
	defer wg.Done()


}

func ConcurrentFrequency(runeSlice []string) FreqMap{
	m3 := FreqMap{}
	wg.Add(len(runeSlice))
	var chans [3]chan FreqMap
	for i := range chans {
		chans[i] = make(chan FreqMap)
		go ConcurrentResponse(runeSlice[i],chans[i])
	}
	wg.Wait()
	for i := range chans {
		for aa,bb := range <-chans[i]{
			fmt.Println(string(aa) + " : "+ strconv.Itoa(bb))
		}
		close(chans[i])
	}
	return  m3
}*/