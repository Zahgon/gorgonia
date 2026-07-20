package main

import "strings"

const START rune = 0x02
const END rune = 0x03

var sentences []string
var vocab []rune
var vocabIndex map[rune]int

func initVocab(ss []string, thresh int) { _ = "STUB: not implemented"; return }

func init() {
	sentencesRaw := strings.Split(corpus, "\n")
	for _, s := range sentencesRaw {
		s2 := strings.TrimSpace(s)
		if s2 != "" {
			sentences = append(sentences, s2)
		}
	}

	initVocab(sentences, 1)
}
