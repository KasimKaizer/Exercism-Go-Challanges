// Package letter contains solution for the Parallel letter Frequency exercise on Exercism.
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
		go count(text, results, &wg) // create a go routine to count all runes in a text, and add 1 to wg.
	}

	go func() {
		wg.Wait()      // wait for all goroutines to have completed.
		close(results) // close results channel, so we can range over results channel without any issue.
	}()

	// what we are doing here is that we have a FreqMap for each text in texts array, now we combine the data
	// from all those FreqMap into a single map.
	// maybe we could run this inside a go routines as well, like divide the counting to different go routines
	// which each working on part of the result channel, and then we just range over the resulting maps and combine
	// them. but that would be unnecessary for this exercise.
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
	// signal done to the passed Wait Group once the execution of this func has completed.
	defer wg.Done()
	tempMap := make(FreqMap)

	for _, char := range text {
		tempMap[char]++ // count frequency of each rune in the text
	}
	res <- tempMap // send the map through res.
}

// Another solution to this problem, this solution uses worker pools, I though this was just
// interesting and wanted to play around a bit.
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
