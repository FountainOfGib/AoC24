package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func isSafe(line []string, increasing bool) bool {
	safe := true
	dampened := []string{}
	for i := 1; i < len(line); i++ {
		firstValue, _ := strconv.Atoi(line[i-1])
		secondValue, _ := strconv.Atoi(line[i])
		if (increasing && (secondValue-firstValue > 3 || secondValue-firstValue < 1)) ||
			(!increasing && (firstValue-secondValue > 3 || firstValue-secondValue < 1)) {
			if len(dampened) != 0 {
				fmt.Println(line, "not safe", "increasing:", increasing, firstValue, secondValue, "diff:", secondValue-firstValue, "at i:", i)
				return false
			} else {
				dampened = append(line[:i], line[i+1:]...)
				i--
			}
		}
	}
	return safe
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
		firstValue, _ := strconv.Atoi(line[0])
		secondValue, _ := strconv.Atoi(line[1])
		if isSafe(line, firstValue < secondValue) {
			safeCount += 1
			fmt.Println(line, "safe")
		}
	}
	fmt.Println("Safe count:", safeCount)
}
