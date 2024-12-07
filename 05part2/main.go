package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileName = "input.txt"

func reorderLine(orderMap map[int][]int, inputArray []int) []int {
	arrayCopy := make([]int, len(inputArray))
	copy(arrayCopy, inputArray)
	for !isLineValid(orderMap, arrayCopy) {
		for i, v := range arrayCopy {
			invalidValues, exists := orderMap[v]
			if !exists {
				continue
			}
			loc := occoursBefore(arrayCopy[:i], invalidValues)
			if loc >= 0 {
				tempSwapBuffer := arrayCopy[loc]
				arrayCopy[loc] = arrayCopy[i]
				arrayCopy[i] = tempSwapBuffer
				break
			}
		}
	}
	return arrayCopy
}

func occoursBefore(previousValues []int, invalidValues []int) int {
	// fmt.Printf("previousValues: %v ivalidValues: %v\n", previousValues, invalidValues)
	for i, previousVal := range previousValues {
		for _, invalidVal := range invalidValues {
			if previousVal == invalidVal {
				return i
			}
		}
	}
	return -1
}

func isLineValid(orderMap map[int][]int, inputArray []int) bool {
	for i, v := range inputArray {
		invalidValues, exists := orderMap[v]
		if !exists {
			continue
		}
		if occoursBefore(inputArray[:i], invalidValues) >= 0 {
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
		} else {
			newline := reorderLine(orderMap, intArray)
			fmt.Printf("intArray: %v was invalid!\n", newline)
			total += newline[len(newline)/2]
		}
	}

	fmt.Println("total:", total)
}
