// Package letter contains tools to calculate letter frequency in a given text.
package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently and returns
// the data in form of FreqMap
func ConcurrentFrequency(texts []string) FreqMap {
	// self imposed constraint, don't use len() func.

	var wg sync.WaitGroup
	results := make(chan FreqMap)
	for _, text := range texts {
		wg.Add(1)
		go count(text, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	output := make(FreqMap)
	for item := range results {
		for key, count := range item {
			output[key] += count // combine all maps from result channel.
		}
	}
	return output
}

// count takes a string and counts frequency of all characters in that passed string and sends the output,
// which is in form of a FreqMap through the res channel.
func count(text string, res chan<- FreqMap, wg *sync.WaitGroup) {
	defer wg.Done()
	tempMap := make(FreqMap)

	for _, char := range text {
		tempMap[char]++
	}
	res <- tempMap
}

// Another solution to this problem, this solution uses worker pools, I though this was just
// interesting and wanted to play around with it a bit.
/*
func ConcurrentFrequency(texts []string) FreqMap {
	textLen := len(texts)
	jobs := make(chan string, textLen)
	result := make(chan FreqMap, textLen)
	for w := 0; w < textLen; w++ {
		go worker(result, jobs)
	}

	for _, text := range texts {
		jobs <- text
	}

	output := make(FreqMap)
	for i := 0; i < textLen; i++ {
		for char, count := range <-result {
			output[char] += count
		}
	}
	return output
}

func worker(res chan<- FreqMap, jobs <-chan string) {
	for job := range jobs {
		tempMap := make(FreqMap)
		for _, char := range job {
			tempMap[char]++
		}
		res <- tempMap
	}
}
*/
