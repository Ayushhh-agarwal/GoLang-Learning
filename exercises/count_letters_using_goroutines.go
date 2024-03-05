package exercises

import (
	"fmt"
	"sync"
)

//  Question :
//	For a given set of strings, find out the frequency of each letter in all the strings parallel.
//	For example, if given the following input.
//															Expectations: Slices, Goroutines, Channels
//	Input : [“quick”, “brown”, “fox”, “lazy”, “dog”]
//
//	The output should be.
//	{
//		“a”: 1,
//		“b”: 1,
//		“c”: 1,
//		…
//		“z”: 1
//	}

func CountLettersUsingGoRoutines() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}

	frequencyChan := make(chan map[int32]int, len(words))

	var wg sync.WaitGroup
	for _, word := range words {
		wg.Add(1)
		go countLetters(word, frequencyChan, &wg)
	}

	wg.Wait()
	close(frequencyChan)

	var listOfFrequencyMaps []map[int32]int
	for freqMap := range frequencyChan {
		listOfFrequencyMaps = append(listOfFrequencyMaps, freqMap)
	}

	result := make(map[int32]int)
	for _, freqMap := range listOfFrequencyMaps {
		for char, count := range freqMap {
			result[char] = result[char] + count
		}
	}

	fmt.Println("Frequency of each letter:")
	printResult(result)
}

func countLetters(word string, frequencyMap chan<- map[int32]int, wg *sync.WaitGroup) {
	defer wg.Done()

	frequency := make(map[int32]int)
	for _, char := range word {
		frequency[char] = frequency[char] + 1
	}

	frequencyMap <- frequency
}

func printResult(result map[int32]int) {
	for char, count := range result {
		fmt.Printf("%c : %d\n", char, count)
	}
}
