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
		if _, hasWord := repeatabilityWords[word]; hasWord {
			repeatabilityWords[word]++
		} else {
			repeatabilityWords[word] = 1
		}
	}

	repeatabilityWords = chooseRecurringMoreThan(repeatabilityWords, 10)

	return repeatabilityWords
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

func chooseRecurringMoreThan(words map[string]int, repeatability int) map[string]int {
	for word, count := range words {
		if count < repeatability {
			delete(words, word)
		}
	}

	return words
}

func clearNumbers(words []string) []string {
	newWords := make([]string, 0, len(words))

	for _, word := range words {
		if isInt(word) {
			continue
		}

		newWords = append(newWords, word)
	}

	return newWords
}

func isInt(strVal string) bool {
	if _, err := strconv.Atoi(strVal); err == nil {
		return true
	}

	return false
}
