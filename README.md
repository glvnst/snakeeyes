# snakeeyes

This command-line utility generates random [diceware](https://en.wikipedia.org/wiki/Diceware)-style passphrases using the [EFF's passphrase word lists](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases), including their [FANDOM Wikia-based word lists](https://www.eff.org/deeplinks/2018/08/dragon-con-diceware). For reference, the EFF provides detailed [instructions for generating passphrases using these word lists](https://www.eff.org/dice) with regular [dice](https://en.wikipedia.org/wiki/Dice). This program replaces the table dice with a [cryptographically secure pseudorandom number generator](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator): the one from your operating system as proxied by [go's crypto/rand package](https://godoc.org/crypto/rand).

This project utilizes the hard work of the team at [EFF](https://www.eff.org/), particularly the work of [Joseph Bonneau](https://www.eff.org/about/staff/joseph-bonneau). **This project is NOT affiliated with the EFF** but you should [donate to EFF](https://www.eff.org/donate/) because their work is essential.

**This program comes with ABSOLUTELY NO WARRANTY; it is distributed under the terms of the GNU Affero General Public License 3.0. See the `COPYING` file or visit the following URL for the complete license terms:** <https://www.gnu.org/licenses/agpl-3.0.html/>. Be aware that using a software passphrase generator is less secure against advanced attacks than using fair dice and paper word lists, _so if you see your IT department playing with dice you should buy them lunch because they're going the extra mile_.


## Notes / Todo / Status

* See the [analysis of the modified word lists](wordlists/analysis/analysis.md)
* I want to improve the automated tests.
* It would be nice to have an option to print entropy information about the lists used and passwords generated but I haven't really looked into those calculations and could use some help.
* Because password prompts sometimes have silly length limits, it would also be nice to be able to specify a target total character length for the generated passphrases. I think I need to research the safest way to accomplish that without firing any crypto footguns.
* I'm using [`go generate`](https://blog.golang.org/generate) (to `go`-ify the word lists) and I used this [nice intro](https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/). Eventually I want to rewrite the word list fetching and preprocessing in go (it is currently based on Make, curl, posix shell utilities, and perl).
* I want to enable some kind of auto update mechanism
	* Using [The Update Framework](https://theupdateframework.com/) seems like a good idea
		* [flynn's go-tuf](https://github.com/flynn/go-tuf) and [kolide's updater](https://github.com/kolide/updater) are golang implementations
		* I have [some](https://github.com/theupdateframework/notary/issues/1566) [concerns](https://github.com/theupdateframework/notary/issues/1564) about the status of the TUF project / notary
		* I'm not sure if I have to run my own notary instance or if I can use a public one hosted somewhere -- need to look into it
	* Restic rolled their own [update system based on github API calls](https://github.com/restic/restic/tree/master/internal/selfupdate)
* I want to document the exact commands for verifying that a snakeeyes binary matches a given commit in the repo. Good starting place: <https://blog.filippo.io/reproducing-go-binaries-byte-by-byte/>

## Example

When invoked without arguments, snakeeyes prints three passphrases consisting of six words each.

```
$ snakeeyes
shower-hypnosis-stability-quit-rabid-stumbling
footwear-cubical-monotype-cognition-garage-pentagon
sleek-iodine-reaction-faction-confirm-handsaw
```

You can specify that you want words from a different word list, such as EFF's first short list. In this example we're also specifying that we want 8 words per passphrase:

```
$ snakeeyes --list memorable --words 8 --delimiter ' '
punk dean snout level donor work dude grab
spoon class penny mardi wink math rant chuck
hurt squad dice tank point recap chew broad
```

## Help Text

Invoking snakeeyes with the `-h` or `--help` arguments will produce the following output:

```
NAME:
   snakeeyes - command-line diceware-style passphrase generator

USAGE:

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


VERSION:
   DEV (No build date recorded. / No commit ID recorded.)

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --words value      number of words to include in each generated passphrase (default: 6)
   --phrases value    number of passphrases to generate (default: 3)
   --delimiter value  delimiter between words in a passphrase (default: "-")
   --list value       word list to use; one of: eff, memorable, touchscreen, got, potter, trek, wars (default: "eff")
   --help, -h         show help
   --version, -v      print the version
```
