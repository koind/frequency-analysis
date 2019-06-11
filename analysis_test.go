package frequency_analysis

import (
	"testing"
)

var testText = `Lorem ipsum dolor sit ipsum, mei labores accumsan labores ei.`
var expected = map[string]int{
	"ipsum":    2,
	"Lorem":    1,
	"accumsan": 1,
	"ei":       1,
	"labores":  1,
	"mei":      1,
	"sit":      1,
	"dolor":    1,
}

func TestAnalyze(t *testing.T) {
	result := Analyze(testText)

	for text, _ := range result {
		if _, hasText := expected[text]; !hasText {
			t.Errorf("Strings must by match: %v - %v", text, expected)
		}
	}
}
