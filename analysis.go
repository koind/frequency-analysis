package frequency_analysis

import (
	"strconv"
	"strings"
	"unicode"
)

func Analyze(text string) map[string]int {
	if text == "" {
		return nil
	}

	repeatabilityWords := map[string]int{}

	words := leaveLetters(text)
	for _, word := range words {
		repeatabilityWords[word]++
	}

	return choosePopularWords(repeatabilityWords, 10)
}

func leaveLetters(text string) []string {
	if text == "" {
		return nil
	}

	f := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	data := strings.FieldsFunc(text, f)

	return clearNumbers(data)
}

func choosePopularWords(words map[string]int, quantity int) map[string]int {
	popularWord := rankByWordCount(words)
	list := make(map[string]int, quantity)

	i := 0
	for _, pl := range popularWord {
		if i <= quantity {
			list[pl.Key] = pl.Value
		}
		i++
	}

	return list
}

func clearNumbers(words []string) []string {
	newWords := make([]string, 0, len(words))

	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		}

		newWords = append(newWords, word)
	}

	return newWords
}
