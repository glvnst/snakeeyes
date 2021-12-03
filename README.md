# snakeeyes

This command-line utility generates random [diceware](https://en.wikipedia.org/wiki/Diceware)-style passphrases using the [EFF's passphrase word lists](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases), including their [FANDOM Wikia-based word lists](https://www.eff.org/deeplinks/2018/08/dragon-con-diceware). For reference, the EFF provides detailed [instructions for generating passphrases using these word lists](https://www.eff.org/dice) with regular [dice](https://en.wikipedia.org/wiki/Dice). This program replaces the table dice with a [cryptographically secure pseudorandom number generator](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator): the one from your operating system as proxied by [go's crypto/rand package](https://godoc.org/crypto/rand).

This project utilizes the hard work of the team at [EFF](https://www.eff.org/), particularly the work of [Joseph Bonneau](https://www.eff.org/about/staff/joseph-bonneau). **This project is NOT affiliated with the EFF** but you should [donate to EFF](https://www.eff.org/donate/) because their work is essential.

**This program comes with ABSOLUTELY NO WARRANTY; it is distributed under the terms of the GNU Affero General Public License 3.0. See the `COPYING` file or visit the following URL for the complete license terms:** <https://www.gnu.org/licenses/agpl-3.0.html/>. Be aware that using a software passphrase generator is less secure against advanced attacks than using fair dice and paper word lists, _so if you see your IT department playing with dice you should buy them lunch because they're going the extra mile_.


## Notes / Todo / Status

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
jailbreak stump omega giveaway consoling reclaim
exploring sweep actress unfrozen vertebrae chain
try qualify twig stubbly darwinism elevation
```

You can specify that you want words from a different word list, such as EFF's first short list. In this example we're also specifying that we want 8 words per passphrase:

```
$ snakeeyes -list memorable -words 8 -delimiter -
rigor-stash-twins-elf-shun-sax-ahead-tidal
shape-yelp-barge-juicy-rant-stop-shush-fax
punch-dwarf-mummy-mace-stem-uncle-yoyo-boney
```

## Help Text

Invoking snakeeyes with the `-h` or `--help` arguments will produce the following output:

```
usage: snakeeyes [ [-h|--help] | [-version] | [-words n] [-phrases n] [-list {eff,memorable,touchscreen,got,potter,trek,wars}] ]

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

  -delimiter string
    	the delimiter between words in a passphrase (default " ")
  -list string
    	the word list to choose words from (default "eff")
  -phrases int
    	the number of passphrases to generate (default 3)
  -version
    	report version number and exit
  -words int
    	the number of words to include in each generated passphrase (default 6)
```