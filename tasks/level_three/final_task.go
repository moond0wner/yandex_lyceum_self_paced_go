package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	Word  string
	Count int
}

func deleteSymbolsInText(text string) string {
	symbols := []string{",", ".", "!", "?"}
	result := text
	for _, symbol := range symbols {
		result = strings.ReplaceAll(result, symbol, "")
	}
	return result
}

func getTopWords(wordMap map[string]int, n int) []string {
	pairs := make([]Pair, 0, len(wordMap))
	for word, count := range wordMap {
		pairs = append(pairs, Pair{Word: word, Count: count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	topWords := make([]string, 0, n)
	for i := 0; i < len(pairs) && i < n; i++ {
		topWords = append(topWords, pairs[i].Word)
	}

	return topWords
}

func AnalyzeText(text string) {
	sliceWords := strings.Split(deleteSymbolsInText(strings.ToLower(text)), " ")
	countWords := len(sliceWords)
	wordMap := map[string]int{}
	for _, word := range sliceWords {
		wordMap[word]++
	}
	countUniqueWords := len(wordMap)

	mostFrequentWord := ""
	maxCount := 0
	for word, count := range wordMap {
		if count > maxCount {
			maxCount = count
			mostFrequentWord = word
		}
	}

	topWords := getTopWords(wordMap, 5)
	topWordsText := ""
	for _, word := range topWords {
		topWordsText += fmt.Sprintf("\n\"%s\": %d раз", word, wordMap[word])
	}

	fmt.Println(fmt.Sprintf("Количество слов: %d\n"+
		"Количество уникальных слов: %d\n"+
		"Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n"+
		"Топ-5 самых часто встречающихся слов:%s",
		countWords, countUniqueWords, mostFrequentWord, maxCount, topWordsText))
}
