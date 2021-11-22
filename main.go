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
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

//go:generate go run helpers/mkwordlists.go

// filled at build time with ldflags by GoReleaser (part of build action)
var (
	product = "snakeeyes"
	version = "DEV"
	commit  = "No commit ID recorded."
	date    = "No build date recorded."
)

const helpText = `usage: %s [ [-h|--help] | [-version] | [-words n] [-phrases n] [-list {eff,memorable,touchscreen,got,potter,trek,wars}] ]

This command-line utility generates random passphrases using the Electronic
Frontier Foundation's passphrase word lists. For more info visit these articles:
https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases
https://www.eff.org/deeplinks/2018/08/dragon-con-diceware

This software project is NOT affiliated with EFF, but if you find this utility
useful, please donate to EFF at:
https://www.eff.org/donate/

If you find bugs or would like to get or contribute to the source code of this
software project, please visit:
https://github.com/glvnst/snakeeyes

This program comes with ABSOLUTELY NO WARRANTY; it is distributed under the
terms of the GNU Affero General Public License 3.0. Visit the following URL for
more information:
https://www.gnu.org/licenses/agpl-3.0.html/

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
		wordCount     = flag.Int("words", 6, "the number of words to include in each generated passphrase")
		phraseCount   = flag.Int("phrases", 3, "the number of passphrases to generate")
		delimiter     = flag.String("delimiter", " ", "the delimiter between words in a passphrase")
		listName      = flag.String("list", "eff", "the word list to choose words from")
		reportVersion = flag.Bool("version", false, "report version number and exit")
	)
	flag.Usage = usage
	flag.Parse()

	if *reportVersion {
		die("%s %s\n %s\n %s\n", product, version, commit, date)
	}

	// parsing
	wordList, ok := WordLists[*listName]
	if !ok {
		die("No such list \"%s\".\n", *listName)
	}

	for p := 0; p < *phraseCount; p++ {
		fmt.Println(GenPassphrase(wordList, *wordCount, *delimiter))
	}
}
