# WIP: snakeeyes

This command-line utility generates random passphrases using the [EFF's passphrase wordlists](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).

* <https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases>
* <https://www.eff.org/deeplinks/2018/08/dragon-con-diceware>
* <https://www.eff.org/dice>
* <https://ssd.eff.org/en/module/animated-overview-how-make-super-secure-password-using-dice>

Be aware that using software to generate passphrases may be less secure against some forms of attack than using fair dice and paper wordlists so if you see your IT department playing with dice you should buy them lunch.

## Notes / Todo / Status / Turtles

* Currently compiles / runs using the <https://godoc.org/crypto/rand> CS-PRNG.
* I need to write some tests.
* It would be nice to be able to print entropy information about the lists as loaded but I haven't really looked into those calculations.
* It would also be nice to be able to specify a target total character length for passphrases but I think I need to research the safest way to accomplish that.
* Experimenting with go generate <https://blog.golang.org/generate> for loading the word lists. This seems like a fine guide on that subject <https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/>

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
$ snakeeyes -list memorable -words 8
bribe fiber doll slum keg stump dense fruit
donor emu date cried cheer thong gem cash
trick graph cane stage fable array mocha blush
```


## Help Text

Invoking snakeeyes with the `-h` or `--help` arguments will produce the following output:

```
usage: snakeeyes [ [-h|--help] | [-version] | [-words n] [-phrases n] [-list {eff,memorable,touchscreen,got,potter,trek,wars}] ]

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

  -list string
    	the wordlist to choose words from (default "eff")
  -phrases int
    	the number of passphrases to generate (default 3)
  -version
    	report version number and exit
  -words int
    	the number of words to include in each generated passphrase (default 6)
```