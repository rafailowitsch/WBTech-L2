package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func findAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	anagrams := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		runes := []rune(word)
		slices.Sort(runes)
		sortedWord := string(runes)

		if !slices.Contains(anagrams[sortedWord], word) {
			anagrams[sortedWord] = append(anagrams[sortedWord], word)
		}
	}

	// Формируем мапу множеств анаграмм
	for _, words := range anagrams {
		if len(words) == 1 {
			continue
		}

		sort.Strings(words)
		key := words[0]
		anagramSets[key] = words
	}

	return anagramSets
}

func main() {
	words := []string{"Пятак", "пятка", "тяпка", "птк", "листок", "слиток", "столиК", "листок"}

	fmt.Println(findAnagramSets(words))
}
