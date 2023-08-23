package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const helpText = `
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
got         - 3,996 words, inspired by Game of Thrones
potter      - 3,998 words, inspired by Harry Potter
trek        - 3,998 words, inspired by Star Trek
wars        - 3,993 words, inspired by Star Wars

NOTE: The "inspired by" lists above are forked from EFF's FANDOM Wikia-based
lists and have additional filtering applied. This filtering removed words which
contain non-ASCII characters, such as the word "café", because entering these
words into password prompts is sometimes difficult. Unfortunately this
filtering reduces the strength of the lists somewhat. Note also that these
particular lists contain hyphenated words.

If you choose from shorter lists, you need to include more words in your
passphrase to achieve the same levels of entropy. See the first EFF link
above for more information.
`

func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

func warn(warning string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, warning, a...)
}

func die(message string, a ...interface{}) {
	warn(message, a...)
	os.Exit(1)
}
