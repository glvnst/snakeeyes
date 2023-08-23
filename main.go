// Copyright (C) 2020 Ben Burke
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

//go:generate go run helpers/mkwordlists.go

// filled at build time with ldflags by GoReleaser (part of build action)
var (
	product = "snakeeyes"
	version = "DEV"
	commit  = "No commit ID recorded."
	date    = "No build date recorded."
)

// GenPassphrase randomly chooses nWords from the given dictionary and returns them joined by the given delimiter
func GenPassphrase(dictionary []string, nWords int, delimiter string) string {
	phrase := make([]string, nWords)
	dictLen := big.NewInt(int64(len(dictionary)))

	for w := 0; w < nWords; w++ {
		wordIndexBig, err := rand.Int(rand.Reader, dictLen)
		if err != nil {
			panic(err)
		}
		wordIndex := wordIndexBig.Int64()
		phrase[w] = dictionary[wordIndex]
	}

	return strings.Join(phrase, delimiter)
}

func GenBoundedPassphrase(wordList []string, maxPassphraseLength int, delimiter string, expectedWordCount int, expectedEntropy float64, tweakWordCase bool, injectInts bool, injectSymbols bool) (string, error) {
	if len(wordList) == 0 {
		return "", errors.New("word list cannot be empty")
	}

	// Calculate the entropy of the word list
	wordEntropy := math.Log2(float64(len(wordList)))

	// Create a slice to hold the words of the passphrase
	passphrase := make([]string, 0, expectedWordCount)

	// Track the length of the passphrase
	length := 0

	for i := 0; i < expectedWordCount; i++ {
		// Select a word at random
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(wordList))))
		word := wordList[index.Int64()]

		// Tweak the case of the word if specified
		if tweakWordCase {
			// Note: this is a basic implementation that just toggles case.
			// A more complete implementation might include more case variations.
			word = strings.ToUpper(word)
		}

		// Check if adding the word would exceed the max length
		if length+len(word)+len(delimiter) > maxPassphraseLength {
			break
		}

		// Add the word to the passphrase
		passphrase = append(passphrase, word)

		// Update the length
		length += len(word) + len(delimiter)
	}

	// Inject a random 2-digit number if specified
	if injectInts {
		position, _ := rand.Int(rand.Reader, big.NewInt(int64(len(passphrase)+1)))
		number, _ := rand.Int(rand.Reader, big.NewInt(100))
		passphrase = append(passphrase[:position.Int64()], append([]string{fmt.Sprintf("%02d", number)}, passphrase[position.Int64():]...)...)
	}

	// Inject a random symbol if specified
	if injectSymbols {
		symbols := []string{"!", "@", "#"}
		position, _ := rand.Int(rand.Reader, big.NewInt(int64(len(passphrase)+1)))
		symbolIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		passphrase = append(passphrase[:position.Int64()], append([]string{symbols[symbolIndex.Int64()]}, passphrase[position.Int64():]...)...)
	}

	// Calculate the entropy of the generated passphrase
	actualEntropy := float64(len(passphrase)) * wordEntropy

	// Check if the generated passphrase meets the expected entropy
	if actualEntropy < expectedEntropy {
		return "", errors.New("generated passphrase does not meet expected entropy")
	}

	// Join the words together with the delimiter to form the final passphrase
	return strings.Join(passphrase, delimiter), nil
}

// AverageWordLength returns the average length of word in a wordlist that is
// characterized by the given wordLenCounts map.
func AverageWordLength(wordLenCounts map[int]int, maxLen int) float64 {
	// Calculate total count of words
	totalCount := 0
	for wordLen, count := range wordLenCounts {
		if wordLen > maxLen {
			continue
		}
		totalCount += count
	}

	// Calculate total sum of word counts
	totalSum := 0
	for wordLen, count := range wordLenCounts {
		if wordLen > maxLen {
			continue
		}
		totalSum += wordLen * count
	}

	// Calculate weighted average length
	return float64(totalSum) / float64(totalCount)
}

// FilterWordList filters the provided wordlist based on the maximum passphrase
// length, to optimize the entropy of a generated passphrase; with entropy in
// a limited-length passphrase there is a tension between factors:
//  1. number of words in the pool you're randomly choosing words from (more is better)
//  2. number of words used in the passphrase (more is a lot better)
//  3. if you have a max passphrase string length (like a lot of legacy websites)
//     you are better off choosing from short words because you can fit more of
//     them in the passphrase (see #2 above) but selecting only from short words
//     is using a smaller word pool (#1 above)
//
// Here we attempt to find the optimal word pool word length cutoff for the
// given wordlist and maxPassphraseLen; we return a filtered version of the
// wordList
func FilterWordList(wordList []string, maxPassphraseLen int) []string {
	var wordLenCounts = make(map[int]int)
	var cumulativeWordLenCounts = make(map[int]int)

	// Create map of word lengths to counts (how many words have a given length)
	for _, word := range wordList {
		wordLenCounts[len(word)]++
	}

	// Create map of cumulative word counts at each given length; e.g. at a
	// cut-off of 3 we have 10 1-letter words, 30 2-letter words, 50 3-letter
	// words in the pool, so cumulativeWordLenCounts[3] == 90 in this example
	accumulatedCount := 0
	for _, wordLen := range SortedKeys(wordLenCounts) {
		accumulatedCount += wordLenCounts[wordLen]
		cumulativeWordLenCounts[wordLen] = accumulatedCount
	}

	// Find optimal word length cut-off to maximize generated passphrase entropy
	var optimalWordLenLimit int
	var optimalWordLenLimitEntropy float64
	for _, wordLen := range SortedKeys(cumulativeWordLenCounts) {
		wordPoolSize := cumulativeWordLenCounts[wordLen]
		avgWordLen := AverageWordLength(wordLenCounts, wordLen)

		// Calculate number of words that can fit into the passphrase
		// -1 because there is no trailing delimiter on the final passphrase
		// +1 for inter-word delimiters
		numWords := int(float64(maxPassphraseLen-1) / (avgWordLen + 1.0))

		// Calculate entropy
		entropy := float64(numWords) * math.Log2(float64(wordPoolSize))
		if entropy >= optimalWordLenLimitEntropy {
			optimalWordLenLimitEntropy = entropy
			optimalWordLenLimit = wordLen
		}
		fmt.Printf("wordLen: %d; wordPoolSize: %d; avgWordLen: %f; numWords: %d; entropy: %f\n",
			wordLen, wordPoolSize, avgWordLen, numWords, entropy)
	}
	fmt.Printf("maxPassphraseLen: %d; optimalWordLenLimit: %d; optimalWordLenLimitEntropy: %f\n", maxPassphraseLen, optimalWordLenLimit, optimalWordLenLimitEntropy)

	// Filter words
	var filteredWords []string
	for _, word := range wordList {
		if len(word) <= optimalWordLenLimit {
			filteredWords = append(filteredWords, word)
		}
	}

	return filteredWords
}

func handleGenerate(args *cli.Context) error {
	// parsing
	listName := args.String("list")
	phraseCount := args.Int("phrases")
	wordCount := args.Int("words")
	delimiter := args.String("delimiter")

	wordList, ok := WordLists[listName]
	if !ok {
		die("No such list \"%s\".\n", listName)
	}

	for p := 0; p < phraseCount; p++ {
		fmt.Println(GenPassphrase(wordList, wordCount, delimiter))
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:      product,
		Usage:     "command-line diceware-style passphrase generator",
		UsageText: helpText,
		Version:   fmt.Sprintf("%s (%s / %s)", version, date, commit),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "words",
				Value: 6,
				Usage: "number of words to include in each generated passphrase",
			},
			&cli.IntFlag{
				Name:  "phrases",
				Value: 3,
				Usage: "number of passphrases to generate",
			},
			&cli.StringFlag{
				Name:  "delimiter",
				Value: "-",
				Usage: "delimiter between words in a passphrase",
			},
			&cli.StringFlag{
				Name:  "list",
				Value: "eff",
				Usage: " word list to use; one of: eff, memorable, touchscreen, got, potter, trek, wars",
			},
		},
		Action: handleGenerate,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
