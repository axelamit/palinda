package main

import (
	"fmt"
	"time"
	"strings"
	"io/ioutil"
	"sync"
)

const DataFile = "loremipsum.txt"

func count(arr []string, maps chan map[string]int, wg *sync.WaitGroup) {
	freqs := make(map[string]int)
	for _,s := range arr {
		w := strings.ReplaceAll(strings.ReplaceAll(s, ",", ""), ".", "")
		freqs[strings.ToLower(w)]++
	}
	maps <- freqs
	wg.Done()
}

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freqs := make(map[string]int)
	s := strings.Fields(text)

	parts := 8
	p_len := len(s)/parts

	maps := make(chan map[string]int, parts)
	wg := new(sync.WaitGroup)
	wg.Add(parts)

	for i := 0; i < parts; i++ {
		index := (i+1)*p_len
		if i == parts-1 {
			index = len(s)
		}
		go count(s[i*p_len:index], maps, wg)
	}

	for i := 0; i < parts; i++ {
		mp := <-maps 
		for k,v := range mp {
			freqs[k]+=v
		}
	}

	wg.Wait()

	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read in DataFile as a string called data
	data, err := ioutil.ReadFile(DataFile)
	check(err)

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
