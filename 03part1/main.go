package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var fileName = "input.txt"

func main() {
	textBlob := "LFG"
	sum := 0
	fmt.Println(textBlob)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open this shit")
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textBlob += scanner.Text()
	}
	mulExpRegex := regexp.MustCompile(`mul\(\d+\,\d+\)`)
	mulExpMatches := mulExpRegex.FindAllString(textBlob, -1)
	for _, match := range mulExpMatches {
		fmt.Println(match)
		numberRegex := regexp.MustCompile(`\d+`)
		numberMatches := numberRegex.FindAllString(match, -1)
		if len(numberMatches) != 2 {
			fmt.Println("FUCK YOU SHITTY regex You fucker", numberMatches, "of len fucking", len(numberMatches))
			return
		}
		intA, _ := strconv.Atoi(numberMatches[0])
		intB, _ := strconv.Atoi(numberMatches[1])
		sum += intA * intB
	}
	// mulExpRegexGrp := regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)
	// mulExpMatchesSubmatch := mulExpRegexGrp.FindAllStringSubmatch(textBlob, -1)
	// fmt.Println(mulExpMatchesSubmatch)
	fmt.Println("ur sum is", sum, "ya fucker")
}
