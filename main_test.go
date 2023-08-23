package main

import (
	"math"
	"strings"
	"testing"
)

var expectedListSizes = map[string]int{
	"eff":         7776,
	"memorable":   1296,
	"touchscreen": 1296,
	"got":         3996,
	"potter":      3998,
	"trek":        3998,
	"wars":        3993,
}

func TestMain(t *testing.T) {
	// ensure that all expected word lists are loaded and have the proper length
	for name, expectedSize := range expectedListSizes {
		list, ok := WordLists[name]
		if !ok {
			t.Errorf("No word list named \"%s\" was found", name)
		}
		actualSize := len(list)
		if actualSize != expectedSize {
			t.Errorf("Word list \"%s\" was not expected size. want: %d, got: %d", name, expectedSize, actualSize)
		}
	}
}

func TestGenPassphrase(t *testing.T) {
	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}
	sampleSize := 100000
	counts := make(map[string]int)
	// the following number is completely pulled out of thin air.
	// I'm not trying to "diehard" the crypto/rand CS-PRNG, but
	// I am looking for structural issues in this code that I could
	// sniff out here. if you are reading this and know the math help
	// me out.
	allowedDeviance := float64(sampleSize) * 0.015

	for i := 0; i < sampleSize; i++ {
		phrase := GenPassphrase(words, 10, " ")
		phraseLen := len(phrase)
		// these phrase lengths should all be 39-59 characters
		// (wordlen * 10 words) + 9 inter-word spaces
		// wordlen is between 3 and 5, as in len("one") and len("three")
		if !(phraseLen > 38 && phraseLen < 60) {
			t.Errorf("expected test passphrase \"%s\" length 31-47 chars, got: %d chars", phrase, phraseLen)
		}

		returnedWords := strings.Split(phrase, " ")
		for _, word := range returnedWords {
			counts[word]++
		}
	}

	for word, count := range counts {
		deviance := math.Abs(float64(sampleSize - count))
		if deviance > allowedDeviance {
			t.Errorf("deviance of sample word \"%s\" is: %f, expecting less than: %f. counts: %+v", word, deviance, allowedDeviance, counts)
		}
	}
}

func TestFilterWordList(t *testing.T) {
	tests := []struct {
		wordList         []string
		maxPassphraseLen int
		expected         int
	}{
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 60,
			expected:         6001,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 45,
			expected:         4501,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 30,
			expected:         3001,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 20,
			expected:         2001,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 10,
			expected:         1001,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 9,
			expected:         901,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 8,
			expected:         801,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 7,
			expected:         701,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 6,
			expected:         601,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 5,
			expected:         501,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 4,
			expected:         401,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 3,
			expected:         301,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 2,
			expected:         101,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 1,
			expected:         101,
		},
		{
			wordList:         WordLists["eff"],
			maxPassphraseLen: 0,
			expected:         0,
		},
	}

	for i, test := range tests {
		actual := len(FilterWordList(test.wordList, test.maxPassphraseLen))
		if actual != test.expected {
			t.Errorf("Test case %d: expected %d, but got %d", i, test.expected, actual)

		}
		// if !reflect.DeepEqual(actual, test.expected) {
		// }
	}
}
