package letter

import "sync"

type FreqMap map[rune]int
type Safe struct {
	m FreqMap
	sync.RWMutex
}

func Frequency(s string) FreqMap {
	m1 := FreqMap{}
	for _, r := range s {
		m1[r]++
	}
	return m1
}

func makeFreqMap(s string, wg *sync.WaitGroup){
	for _, c := range s {
		safe.Lock  ()
		safe.m[c]  ++
		safe.Unlock()
	}
	wg.Done()
}

var safe Safe
func ConcurrentFrequency(stringS []string) FreqMap {
	var wg sync.WaitGroup
	wg.Add(len(stringS))
	safe.m = FreqMap{}
	for _ , s := range stringS { go makeFreqMap(s, &wg) }
	wg.Wait()
	return safe.m
}