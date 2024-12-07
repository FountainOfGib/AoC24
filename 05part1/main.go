package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func occoursBefore(previousValues []int, invalidValues []int) bool {
	fmt.Printf("previousValues: %v ivalidValues: %v\n", previousValues, invalidValues)
	for _, previousVal := range previousValues {
		for _, invalidVal := range invalidValues {
			if previousVal == invalidVal {
				return true
			}
		}
	}
	return false
}

func isLineValid(orderMap map[int][]int, inputArray []int) bool {
	for i, v := range inputArray {
		invalidValues, exists := orderMap[v]
		if !exists {
			continue
		}
		if occoursBefore(inputArray[:i], invalidValues) {
			return false
		}
	}
	return true
}

func main() {
	orderMap := make(map[int][]int)
	total := 0

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("cant open file")
		return
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		splitLine := strings.Split(line, "|")
		A, _ := strconv.Atoi(splitLine[0])
		B, _ := strconv.Atoi(splitLine[1])
		val, exists := orderMap[A]
		if exists {
			orderMap[A] = append(val, B)
		} else {
			orderMap[A] = []int{B}
		}
	}
	fmt.Println(orderMap)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		var intArray []int
		for _, v := range line {
			intVal, _ := strconv.Atoi(v)
			intArray = append(intArray, intVal)
		}
		if isLineValid(orderMap, intArray) {
			fmt.Printf("intArray: %v is valid!\n", intArray)
			total += intArray[len(intArray)/2]
		} else {
			fmt.Printf("intArray: %v is invalid!\n", intArray)
		}
	}

	fmt.Println("total:", total)
}
