package frequency_analysis

import (
	"testing"
)

var testCases = map[string]map[string]int{
	`lorem lorem lorem lorem lorem lorem lorem lorem lorem lorem lorem lorem lorem lorem`: map[string]int{
		"lorem": 14,
	},
}

func TestAnalyze(t *testing.T) {
	for inputText, expected := range testCases {
		if result := Analyze(inputText); result["lorem"] != expected["lorem"] {
			t.Errorf("Strings must by match: %v - %v", expected, result)
		}
	}
}
