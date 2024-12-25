package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)
func main() {
    // TODO: implement parallel word counter
	if len(os.Args) < 2 {
        log.Fatal("Please provide the path to the PDF file.")
    }
	pdfPath := os.Args[1]

	text, err := os.ReadFile(pdfPath)
    if err != nil {
        log.Fatalf("Error reading PDF: %v", err)
    }

	numCores := runtime.NumCPU()
    var wg sync.WaitGroup
    results := make(chan int, numCores)

	start := time.Now()

	lines := strings.Split(string(text), "\n")

	chunkSize := (len(lines) + numCores - 1) / numCores // Ceiling division

	for i := 0; i < len(lines); i += chunkSize {
        end := i + chunkSize
        if end > len(lines) {
            end = len(lines)
        }
        chunk := strings.Join(lines[i:end], "\n")

        wg.Add(1)
        go processChunk(chunk, &wg, results)
    }

	go func() {
        wg.Wait()
        close(results)
    }()

    // Aggregate results
    totalWords := 0
    for count := range results {
        totalWords += count
    }

    fmt.Printf("Total word count: %d\n", totalWords)
	fmt.Printf("Time taken in parallel: %s\n", time.Since(start))


	start2 := time.Now()
	totalWords = 0
	for _, s := range lines {
		totalWords += countWords(s)
	}

	fmt.Printf("Total word count: %d\n", totalWords)
	fmt.Printf("Time taken in sequential: %s\n", time.Since(start2))
}