package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func RemoveAt[T any](slice []T, i int) []T {
	Tcopy := make([]T, len(slice))
	copy(Tcopy, slice[:])
	if i < 0 || i >= len(Tcopy) {
		// Return the original slice if index is out of bounds
		return Tcopy
	}
	// Remove the element by slicing and concatenating
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
			fmt.Println(inputLine, "not safe", "increasing:", increasing, firstValue, secondValue, "diff:", secondValue-firstValue, "at i:", i)
			if dampened {
				return false
			} else {
				fmt.Println("attempting dampening ...")
				for i := 0; i < len(inputLine); i++ {
					if isSafe(RemoveAt(inputLine, i), true) {
						return true
					}
				}
				defer fmt.Println("dampening attempt complete")
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err opening", fileName)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) < 3 {
			fmt.Println("line has no values!!??")
			return
		}
		if isSafe(line, false) {
			safeCount += 1
			fmt.Println(line, "safe")
		}
	}
	fmt.Println("Safe count:", safeCount)
}
