#!/usr/bin/env python3
import json
import os

import matplotlib.pyplot as plt
import numpy as np

input_basedir=".."

def np_encoder(object):
    """helper for json-encoding numpy data"""
    # from: https://stackoverflow.com/a/65151218
    if isinstance(object, np.generic):
        return object.item()


def analyze_wordlist(filename):
    # Read words from file
    with open(os.path.join(input_basedir, filename), "r") as f:
        words = f.read().splitlines()

    # Calculate word lengths
    word_lengths = np.array([len(word) for word in words])

    # Calculate statistics
    mean = np.mean(word_lengths)
    median = np.median(word_lengths)
    first_quartile = np.percentile(word_lengths, 25)
    third_quartile = np.percentile(word_lengths, 75)
    stddev = np.std(word_lengths)
    min_length = np.min(word_lengths)
    max_length = np.max(word_lengths)
    word_count = len(words)
    unique_word_count = len(np.unique(words))

    # Print statistics
    stats = {
        "List Name": filename,
        "Word Count": word_count,
        "Unique Word Count": unique_word_count,
        "Min word length": min_length,
        "Max word length": max_length,
        "Mean word length": mean,
        "Median word length": median,
        "Word length standard deviation": stddev,
        "First quartile of word length": first_quartile,
        "Third quartile of word length": third_quartile,
    }

    # Plots
    # Create bins for each word length
    # we add 2 to include the upper edge
    bins = range(min_length, max_length + 2)

    # Plot histogram with bins
    # align='left' aligns the bins to the left edge
    plt.hist(word_lengths, bins=bins, align="left")

    # ensure we only show ticks for existing lengths
    plt.xticks(range(min_length, max_length + 1))
    plt.xlabel("Word Length")
    plt.ylabel("Count")
    plt.title(f"Histogram: Word Lengths in {filename}")
    # plt.show()
    plot_filename = f"histogram_word_lengths_{os.path.splitext(filename)[0]}.png"
    plt.savefig(plot_filename, dpi=300, bbox_inches="tight")
    plt.close()

    return stats, plot_filename


def main():
    filenames = (
        "eff.txt",
        "effshort1.txt",
        "effshort2.txt",
        "got.txt",
        "potter.txt",
        "startrek.txt",
        "starwars.txt",
    )
    results = []
    for filename in filenames:
        stats, plot_file = analyze_wordlist(filename)
        results.append({"stats": stats, "plot_file": plot_file})

    print(json.dumps(results, indent=2, default=np_encoder))


if __name__ == "__main__":
    main()
