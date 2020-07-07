package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

//go:generate go run helpers/mkwordlists.go

const helpText = `usage: %s [-h|--help] [-words n] [-phrases n] [-list {eff,memorable,touchscreen,got,potter,trek,wars}]

This command-line utility generates random passphrases using the Electronic
Frontier Foundation's passphrase wordlists. For more info visit these articles:
https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases
https://www.eff.org/deeplinks/2018/08/dragon-con-diceware

This software project is NOT affiliated with EFF, but if you find this utility
useful, please donate to EFF at:
https://www.eff.org/donate/

If you find bugs or would like to contribute code and/or give feedback to this
software project, please visit:
https://github.com/glvnst/snakeeyes

Available Word Lists:

eff         - 7,776 words, like Arnold Reinhold's Diceware, but tweaked by EFF
memorable   - 1,296 words, the most memorable and distinct words per EFF
touchscreen - 1,296 words, EFF experiment optimized for typing on software keyboards
got         - 3,996 words, forked from EFF, contains hyphenated words, inspired by Game of Thrones
potter      - 3,998 words, forked from EFF, contains hyphenated words, inspired by Harry Potter
trek        - 3,998 words, forked from EFF, contains hyphenated words, inspired by Star Trek
wars        - 3,993 words, forked from EFF, contains hyphenated words, inspired by Star Wars

The designation "forked from EFF" indicates that a list started with one of EFF's
FANDOM Wikia-based lists and had additional filtering applied to remove words
which contain non-ASCII characters, such as the word "café". Entering these words
into password prompts has proven difficult to do reliably in several cases. This
filtering reduces the quality of the list somewhat.

If you choose from shorter lists, you need to include more words in your
passphrase to achieve the same levels of entropy. From the first EFF link
above:

"Passphrases generated using the shorter lists will be weaker than the long
list on a per-word basis (10.3 bits/word).  Put another way, this means you
would need to choose more words from the short list, to get comparable security
to the long list—for example, using eight words from the short will provide a
strength of about 82 bits, slightly stronger than six words from the long
list."

Command line options:

`

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

func warn(warning string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, warning, a...)
}

func die(message string, a ...interface{}) {
	warn(message, a...)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, helpText, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var (
		wordCount   = flag.Int("words", 6, "the number of words to include in each generated passphrase")
		phraseCount = flag.Int("phrases", 3, "the number of passphrases to generate")
		listName    = flag.String("list", "eff", "the wordlist to choose words from")
	)
	flag.Usage = usage
	flag.Parse()

	// parsing
	wordList, ok := WordLists[*listName]
	if !ok {
		die("No such list \"%s\".\n", *listName)
	}

	for p := 0; p < *phraseCount; p++ {
		fmt.Println(GenPassphrase(wordList, *wordCount, " "))
	}
}
