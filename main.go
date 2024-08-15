package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	conllu "github.com/brewingweasel/go-conllu"
)

type FrequencyDetails struct {
	Lemma     string
	Frequency int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <path-to-conllu-file>")
	}
	path := os.Args[1]

	sentences, errs := conllu.ParseFile(path)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("error: %v\n", err)
		}
	}

	frequencies := make(map[string]int)

	for _, sentence := range sentences {
		for _, token := range sentence.Tokens {
			if token.UPOS != "PUNCT" && token.UPOS != "SYM" && token.UPOS != "NUM" && token.UPOS != "X" && token.UPOS != "PROPN" {
				frequencies[token.Lemma] = frequencies[token.Lemma] + 1
			}
		}
	}

	frequencySlice := []FrequencyDetails{}

	for lemma, frequency := range frequencies {
		frequencySlice = append(frequencySlice, FrequencyDetails{Lemma: lemma, Frequency: frequency})
	}

	slices.SortFunc(frequencySlice,
		func(a, b FrequencyDetails) int {
			return cmp.Compare(b.Frequency, a.Frequency)
		})

	outputFile := filepath.Join("output", strings.TrimSuffix(path, ".conllu")+"_frequency")
	_ = os.Mkdir("output", os.ModePerm)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, details := range frequencySlice {
		f.WriteString(details.Lemma + "\n")
	}
	f.Close()
}
