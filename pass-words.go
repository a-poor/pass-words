package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const VOCAB_FILE = "vocab.txt"

func readVocab(path string) ([]string, error) {
	data, err := ioutil.ReadFile(VOCAB_FILE)
	if err != nil {
		return nil, err
	}
	vocab := strings.Split(string(data), "\n")
	return vocab, nil
}

func wordPicker(vocab *[]string) string {
	i := rand.Intn(len(*vocab))
	return (*vocab)[i]
}

func main() {
	// Declare and parse command-line flags.
	nWords := flag.Int("nwords", 3, "Number of words to randomly generate. (Must be > 0)")
	sep := flag.String("sep", ".", "Character used toseparate random words.")
	flag.Parse()

	// Check that the number of words is positive.
	if *nWords <= 0 {
		fmt.Println("`-nwords` must be > 0")
		os.Exit(1)
	}

	// Import the vocab and check for errors.
	vocab, err := readVocab(VOCAB_FILE)
	if err != nil {
		fmt.Println("Error reading from vocab file:", err)
		os.Exit(1)
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
	fmt.Printf("\nYOUR PASSWORD IS: %s\n", password)
	fmt.Println()

}
