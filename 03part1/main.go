package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var fileName = "input.txt"

type textMatch struct {
	instruction int
	text        string
}

func main() {
	textBlob := "LFG"
	sum := 0
	do := true
	instructionMap := make(map[int]textMatch)
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
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don\'t\(\)`)

	mulExpMatches := mulExpRegex.FindAllStringIndex(textBlob, -1)
	for _, match := range mulExpMatches {
		instructionMap[match[0]] = textMatch{0, textBlob[match[0]:match[1]]}
	}

	doMatches := doRegex.FindAllStringIndex(textBlob, -1)
	for _, match := range doMatches {
		instructionMap[match[0]] = textMatch{1, textBlob[match[0]:match[1]]}
	}

	dontMatches := dontRegex.FindAllStringIndex(textBlob, -1)
	for _, match := range dontMatches {
		instructionMap[match[0]] = textMatch{2, textBlob[match[0]:match[1]]}
	}

	for i := 0; i < len(textBlob); i++ {
		instructionBlock, exists := instructionMap[i]
		if exists {
			switch instructionBlock.instruction {
			case 0:
				if !do {
					fmt.Println(instructionBlock.text, "found, but last found a 'dont' - not adding")
					continue
				}
				fmt.Println("x is mul", instructionBlock.text)
				numberRegex := regexp.MustCompile(`\d+`)
				numberMatches := numberRegex.FindAllString(instructionBlock.text, -1)
				if len(numberMatches) != 2 {
					fmt.Println("FUCK YOU SHITTY regex You fucker", numberMatches, "of len fucking", len(numberMatches))
					return
				}
				intA, _ := strconv.Atoi(numberMatches[0])
				intB, _ := strconv.Atoi(numberMatches[1])
				fmt.Println("adding ", intA*intB, instructionBlock.text)
				sum += intA * intB
			case 1:
				fmt.Println("x is do", instructionBlock.text)
				do = true
			case 2:
				fmt.Println("x is don't", instructionBlock.text)
				do = false
			default:
				fmt.Println("you fucked up")
				os.Exit(1) // punish user for default
			}
		}
	}

	fmt.Printf("instructionMap: %v\n", instructionMap)
	// fmt.Println(match)
	// numberRegex := regexp.MustCompile(`\d+`)
	// numberMatches := numberRegex.FindAllString(match, -1)
	// if len(numberMatches) != 2 {
	// 	fmt.Println("FUCK YOU SHITTY regex You fucker", numberMatches, "of len fucking", len(numberMatches))
	// 	return
	// }
	// intA, _ := strconv.Atoi(numberMatches[0])
	// intB, _ := strconv.Atoi(numberMatches[1])
	// sum += intA * intB
	// mulExpRegexGrp := regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)
	// mulExpMatchesSubmatch := mulExpRegexGrp.FindAllStringSubmatch(textBlob, -1)
	// fmt.Println(mulExpMatchesSubmatch)
	fmt.Println("ur sum is", sum, "ya fucker")
}
