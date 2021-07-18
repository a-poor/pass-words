package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//go:embed vocab.txt
var vocabFile string

func getSplitVocab() ([]string, error) {
	return strings.Split(vocabFile, "\n"), nil
}

func wordPicker(vocab *[]string) string {
	i := rand.Intn(len(*vocab))
	return (*vocab)[i]
}

func main() {
	// Declare and parse command-line flags.
	nWords := flag.Int("nwords", 3, "Number of words to randomly generate. (Must be > 0)")
	sep := flag.String("sep", ".", "Character used toseparate random words.")
	showVocabWords := flag.Bool("vocab", false, "Print out the vocabulary and then quit.")
	showVocabSize := flag.Bool("vocab-size", false, "Show the size of the vocabulary and quit.")
	flag.Parse()

	// Check that the number of words is positive.
	if *nWords <= 0 {
		fmt.Println("`-nwords` must be > 0")
		os.Exit(1)
	}

	// Import the vocab and check for errors.
	vocab, err := getSplitVocab()
	if err != nil {
		fmt.Println("Error reading from vocab file:", err)
		os.Exit(1)
	}

	// Check if the user wants to see the size of the vocabulary.
	if *showVocabSize {
		fmt.Println("Vocab size:", len(vocab))
		return
	}

	// Check if the user wants to SEE the vocabulary.
	if *showVocabWords {
		for _, word := range vocab {
			fmt.Println(word)
		}
		return
	}

	// Count the vocab and check that it's not empty.
	vocabSize := len(vocab)
	if vocabSize < 1 {
		fmt.Println("Error: vocab file is empty")
		os.Exit(1)
	}

	// Set the random seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Generate random words for the password.
	words := make([]string, 0)
	for i := 0; i < *nWords; i++ {
		words = append(words, wordPicker(&vocab))
	}

	// Join the password words
	password := strings.Join(words, *sep)
	fmt.Println(password)

}
