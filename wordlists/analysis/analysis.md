# Analysis of Word Lists

## Word Length

List Name       | Unique Words | Min WL | Max WL | Median WL | Mean WL | WL Std. Dev. | WL Q1 | WL Q3
--------------- | -----------: | -----: | -----: | --------: | ------- | ------------ | ----: | ----:
`eff.txt`       |        7,776 |      3 |      9 |         7 | 6.992   | 1.547        |     6 |     8
`effshort1.txt` |        1,296 |      3 |      5 |         5 | 4.540   | 0.612        |     4 |     5
`effshort2.txt` |        1,296 |      3 |     10 |         7 | 7.316   | 1.638        |     6 |     9
`got.txt`       |        3,996 |      3 |     20 |         7 | 6.901   | 2.311        |     5 |     8
`potter.txt`    |        3,998 |      3 |     20 |         7 | 7.017   | 2.336        |     5 |     8
`startrek.txt`  |        3,998 |      3 |     20 |         7 | 7.152   | 2.420        |     5 |     9
`starwars.txt`  |        3,993 |      3 |     16 |         7 | 6.924   | 2.281        |     5 |     8

Notes:

* All words in a wordlist are unique
* WL refers to "Word Length" in the following table

### Distribution Plots

List Name       | Plot
--------------- | -----------------
`eff.txt`       | [![][fig1]][fig1]
`effshort1.txt` | [![][fig2]][fig2]
`effshort2.txt` | [![][fig3]][fig3]
`got.txt`       | [![][fig4]][fig4]
`potter.txt`    | [![][fig5]][fig5]
`startrek.txt`  | [![][fig6]][fig6]
`starwars.txt`  | [![][fig7]][fig7]



[fig1]: histogram_word_lengths_eff.png
[fig2]: histogram_word_lengths_effshort1.png
[fig3]: histogram_word_lengths_effshort2.png
[fig4]: histogram_word_lengths_got.png
[fig5]: histogram_word_lengths_potter.png
[fig6]: histogram_word_lengths_startrek.png
[fig7]: histogram_word_lengths_starwars.png
