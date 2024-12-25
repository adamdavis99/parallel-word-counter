package main

import (
	"bytes"
	"strings"
	"github.com/ledongthuc/pdf"
	"sync"
)

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)

	if err != nil {
		return "", err
	}
	defer f.Close()
	var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)
	return buf.String(), nil
}

func countWords(text string) int {
    words := strings.Fields(text)
    return len(words)
}

func processChunk(chunk string, wg *sync.WaitGroup, results chan<- int) {
    defer wg.Done()
    wordCount := countWords(chunk)
    results <- wordCount
}
