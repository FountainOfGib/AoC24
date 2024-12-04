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

func getMatchesAt(fileText []string, x int, y int) int { // don' worry about it chief
	width := len(fileText[0])
	height := len(fileText)
	total := 0
	canAppearLeft := x >= 3
	canAppearUp := y >= 3
	canAppearRight := width-(x+1) >= 3
	canAppearDown := height-(y+1) >= 3
	// left case
	if canAppearLeft &&
		fileText[y][x-0] == 'X' &&
		fileText[y][x-1] == 'M' &&
		fileText[y][x-2] == 'A' &&
		fileText[y][x-3] == 'S' {
		total++
	}
	// up case
	if canAppearUp &&
		fileText[y-0][x] == 'X' &&
		fileText[y-1][x] == 'M' &&
		fileText[y-2][x] == 'A' &&
		fileText[y-3][x] == 'S' {
		total++
	}
	// down case
	if canAppearDown &&
		fileText[y+0][x] == 'X' &&
		fileText[y+1][x] == 'M' &&
		fileText[y+2][x] == 'A' &&
		fileText[y+3][x] == 'S' {
		total++
	}
	// right case
	if canAppearRight &&
		fileText[y][x+0] == 'X' &&
		fileText[y][x+1] == 'M' &&
		fileText[y][x+2] == 'A' &&
		fileText[y][x+3] == 'S' {
		total++
	}
	// upleft
	if canAppearUp && canAppearLeft &&
		fileText[y-0][x-0] == 'X' &&
		fileText[y-1][x-1] == 'M' &&
		fileText[y-2][x-2] == 'A' &&
		fileText[y-3][x-3] == 'S' {
		total++
	}
	// upright
	if canAppearUp && canAppearRight &&
		fileText[y-0][x+0] == 'X' &&
		fileText[y-1][x+1] == 'M' &&
		fileText[y-2][x+2] == 'A' &&
		fileText[y-3][x+3] == 'S' {
		total++
	}
	// downright
	if canAppearDown && canAppearRight &&
		fileText[y+0][x+0] == 'X' &&
		fileText[y+1][x+1] == 'M' &&
		fileText[y+2][x+2] == 'A' &&
		fileText[y+3][x+3] == 'S' {
		total++
	}
	// downleft
	if canAppearDown && canAppearLeft &&
		fileText[y+0][x-0] == 'X' &&
		fileText[y+1][x-1] == 'M' &&
		fileText[y+2][x-2] == 'A' &&
		fileText[y+3][x-3] == 'S' {
		total++
	}
	return total
}

func main() {
	xmasFinds := 0
	fileText, _ := fileToSliceOfStrings(fileName)
	fmt.Println(fileText)
	for y := 0; y < len(fileText); y++ {
		for x := 0; x < len(fileText[0]); x++ {
			if fileText[y][x] != 'X' {
				continue
			}
			xmasFinds += getMatchesAt(fileText, x, y)
		}
	}
	fmt.Println(fileText, "\n", xmasFinds)
}
