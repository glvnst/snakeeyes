TARGETS=eff.txt effshort1.txt effshort2.txt got.txt potter.txt startrek.txt starwars.txt
DOWNLOADS=eff_large_wordlist.txt eff_short_wordlist_1.txt eff_short_wordlist_2_0.txt gameofthrones_8k-2018.txt harrypotter_8k_3column-txt.txt memory-alpha_8k_2018.txt starwars_8k_2018.txt

CURL=curl -sSLO
TR=LC_ALL=UTF-8 tr '\r' '\n'

.PHONY: all download clean deepclean

all: $(TARGETS)

download: $(DOWNLOADS)

clean:
	@rm -fv $(TARGETS) *.bak

deepclean:
	@rm -fv $(DOWNLOADS) *.bak

eff.txt: eff_large_wordlist.txt
	cut -f 2 "$<" >"$@"

effshort1.txt: eff_short_wordlist_1.txt
	cut -f 2 "$<" >"$@"

effshort2.txt: eff_short_wordlist_2_0.txt
	cut -f 2 "$<" >"$@"

got.txt: gameofthrones_8k-2018.txt
	$(TR) <"$<" \
	  | iconv -f macroman -t utf-8 \
	  | perl -ne 'next if m/[^a-zA-Z0-9\s\-]/; next unless m/^\d+-\d+-\d+\s+(.+?)\s*$$/; print "$$1\n";' \
	  | sort \
	  | uniq \
	  >"$@"

potter.txt: harrypotter_8k_3column-txt.txt
	$(TR) <"$<" \
	  | iconv -f macroman -t utf-8 \
	  | perl -ne 'next if m/[^a-zA-Z0-9\s\-]/; next unless m/^\d+-\d+-\d+\s+(.+?)\s*$$/; print "$$1\n";' \
	  | sort \
	  | uniq \
	  >"$@"

startrek.txt: memory-alpha_8k_2018.txt
	$(TR) <"$<" \
	  | iconv -f macroman -t utf-8 \
	  | perl -ne 'next if m/[^a-zA-Z0-9\s\-]/; next unless m/^\d+-\d+-\d+\s+(.+?)\s*$$/; print "$$1\n";' \
	  | sort \
	  | uniq \
	  >"$@"

starwars.txt: starwars_8k_2018.txt
	$(TR) <"$<" \
	  | iconv -f macroman -t utf-8 \
	  | perl -ne 'next if m/[^a-zA-Z0-9\s\-]/; next unless m/^\d+-\d+-\d+\s+(.+?)\s*$$/; print "$$1\n";' \
	  | sort \
	  | uniq \
	  >"$@"

eff_large_wordlist.txt:
	$(CURL) "https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt"

eff_short_wordlist_1.txt:
	$(CURL) "https://www.eff.org/files/2016/09/08/eff_short_wordlist_1.txt"

eff_short_wordlist_2_0.txt:
	$(CURL) "https://www.eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt"

gameofthrones_8k-2018.txt:
	$(CURL) "https://www.eff.org/files/2018/08/29/gameofthrones_8k-2018.txt"

harrypotter_8k_3column-txt.txt:
	$(CURL) "https://www.eff.org/files/2018/08/29/harrypotter_8k_3column-txt.txt"

memory-alpha_8k_2018.txt:
	$(CURL) "https://www.eff.org/files/2018/08/29/memory-alpha_8k_2018.txt"

starwars_8k_2018.txt:
	$(CURL) "https://www.eff.org/files/2018/08/29/starwars_8k_2018.txt"
