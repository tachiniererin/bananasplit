package bananasplit

import (
	"github.com/rivo/uniseg"
)

// RuneRange represents a Unicode code-point range
type RuneRange struct {
	Start rune
	End   rune
}

// Word represents a matched part which only contains characters of
// the RuneRanges represented by the type field
type Word struct {
	Text string
	Type string
}

var (
	// See https://unicode.org/charts/ for codepoint ranges

	Dingbats                           = RuneRange{0x2700, 0x27BF}
	OrnamentalDingbats                 = RuneRange{0x1F650, 0x1F67F}
	Emoticons                          = RuneRange{0x1F600, 0x1F64F}
	MiscellaneousSymbols               = RuneRange{0x2600, 0x26FF}
	MiscellaneousSymbolsAndPictographs = RuneRange{0x1F300, 0x1F5FF}
	SupplementalSymbolsAndPictographs  = RuneRange{0x1F900, 0x1F9FF}
	SymbolsAndPictographsExtendedA     = RuneRange{0x1FA70, 0x1FAFF}
	TransportAndMapSymbols             = RuneRange{0x1F680, 0x1F6FF}

	// Emoji & Pictographs ranges
	EmojiRange = []RuneRange{
		Dingbats,
		OrnamentalDingbats, Emoticons,
		MiscellaneousSymbols,
		MiscellaneousSymbolsAndPictographs,
		SupplementalSymbolsAndPictographs,
		SymbolsAndPictographsExtendedA,
		TransportAndMapSymbols,
	}
)

// IsPartOfRange checks if the given rune matches one of the RuneRanges
func IsPartOfRange(r rune, rng []RuneRange) bool {
	for _, v := range rng {
		if r >= v.Start && r <= v.End {
			return true
		}
	}
	return false
}

// SplitByRanges splits the given string by the supplied ranges into Words
// If a given part of a string does not match any ranges, it is tagged as
// unmatched instead.
func SplitByRanges(s string, ranges map[string][]RuneRange) []Word {
	var sentence []Word
	var currentWord = new(Word)

	gr := uniseg.NewGraphemes(s)

	for gr.Next() {
		matched := false
		runes := gr.Runes()

		for name, rng := range ranges {
			if IsPartOfRange(runes[0], rng) {
				if currentWord.Type != name && currentWord.Type != "" && len(currentWord.Text) > 0 {
					sentence = append(sentence, *currentWord)
					currentWord = new(Word)
				}
				if currentWord.Type == "" {
					currentWord.Type = name
				}
				matched = true
				break
			}
		}

		if !matched {
			if (currentWord.Type != "unmatched" || currentWord.Type == "") && len(currentWord.Text) > 0 {
				sentence = append(sentence, *currentWord)
				currentWord = new(Word)
			}
			if currentWord.Type == "" {
				currentWord.Type = "unmatched"
			}
		}

		currentWord.Text += string(runes)
	}

	return append(sentence, *currentWord)
}
