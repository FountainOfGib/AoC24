package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileName = "sample.txt"

func fileToSliceOfStrings(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a slice of strings
	var lines []string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getMatchesAt(fileText []string, x int, y int) bool { // don' worry about it chief
	return (fileText[y-1][x-1] == 'M' && fileText[y+1][x+1] == 'S' ||
		fileText[y-1][x-1] == 'S' && fileText[y+1][x+1] == 'M') &&
		(fileText[y-1][x+1] == 'M' && fileText[y+1][x-1] == 'S' ||
			fileText[y-1][x+1] == 'S' && fileText[y+1][x-1] == 'M')
}

func main() {
	xmasFinds := 0
	fileText, _ := fileToSliceOfStrings(fileName)
	fmt.Println(fileText)
	for y := 1; y < len(fileText)-1; y++ {
		for x := 1; x < len(fileText[0])-1; x++ {
			if fileText[y][x] != 'A' {
				continue
			}
			if getMatchesAt(fileText, x, y) {
				xmasFinds++
			}
		}
	}
	fmt.Println(fileText, "\n", xmasFinds)
}
