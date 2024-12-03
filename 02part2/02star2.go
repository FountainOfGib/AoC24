package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var fileName = "input.txt"

func RemoveAt[T any](slice []T, i int) []T {
	Tcopy := make([]T, len(slice))
	copy(Tcopy, slice[:])
	if i < 0 || i >= len(Tcopy) {
		return Tcopy
	}
	return append(Tcopy[:i], Tcopy[i+1:]...)
}

func isSafe(inputLine []string, dampened bool) bool {
	a, _ := strconv.Atoi(inputLine[0])
	b, _ := strconv.Atoi(inputLine[1])
	increasing := a < b
	for i := 1; i < len(inputLine); i++ {
		firstValue, _ := strconv.Atoi(inputLine[i-1])
		secondValue, _ := strconv.Atoi(inputLine[i])
		if (increasing && (secondValue-firstValue > 3 || secondValue-firstValue < 1)) ||
			(!increasing && (firstValue-secondValue > 3 || firstValue-secondValue < 1)) {
			// fmt.Println(inputLine, "not safe", "increasing:", increasing, firstValue, secondValue, "diff:", secondValue-firstValue, "at i:", i)
			if dampened {
				return false
			} else {
				// fmt.Println("attempting dampening ...")
				// defer fmt.Println("dampening complete ...")
				return isSafe(RemoveAt(inputLine, i-1), true) ||
					isSafe(RemoveAt(inputLine, i), true) ||
					isSafe(RemoveAt(inputLine, 0), true) // stupid hack to get around complex checking of increasing / decreasing
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening", fileName)
		return
	}
	defer file.Close()

	// Channel to communicate lines to workers
	linesChan := make(chan []string)

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 100

	// Safe count
	var safeCount int
	var mu sync.Mutex // Protects access to safeCount

	// Worker function
	worker := func() {
		defer wg.Done()
		for line := range linesChan {
			if len(line) < 3 {
				fmt.Println("Line has no values!!??")
				continue
			}
			if isSafe(line, false) {
				mu.Lock()
				safeCount++
				mu.Unlock()
				// fmt.Println(line, "safe")
			} else {
				// fmt.Println(line, "not safe")
			}
		}
	}

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	// Read lines from the file and send them to workers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		linesChan <- line
	}

	// Close the channel and wait for workers to finish
	close(linesChan)
	wg.Wait()

	fmt.Println("Safe count:", safeCount)
}
