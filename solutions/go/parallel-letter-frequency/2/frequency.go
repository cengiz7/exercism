package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m1 := FreqMap{}
	for _, r := range s {
		m1[r]++
	}
	return m1
}

func ConcurrentFrequency(stringS []string) FreqMap {
	ch := make(chan FreqMap)
	for _, s := range stringS {
		go func(s string) { ch <- Frequency(s) }(s)
	}
	mergeFreq := FreqMap{}
	for range stringS {
		if mergeFreq == nil {
			mergeFreq = <-ch
		} else {
			d := <-ch
			for k, v := range d {
				mergeFreq[k] += v
			}
		}
	}
	return mergeFreq
}