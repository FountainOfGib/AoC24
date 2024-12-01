package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func failHardParseInt(number string) int {
	val, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("err file contains none ints")
		os.Exit(1) // I don't care that this is bad. Go cry.
	}
	return val
}

func parseInput(fileName string) ([]int, map[int]int) {
	A := []int{}
	B := make(map[int]int)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err opening", fileName)
		os.Exit(1) // I don't care that this is bad. Go cry.
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		if len(line) != 2 {
			fmt.Println("err in", fileName, ". must be 2 values separated by 3 spaces")
			os.Exit(1) // I don't care that this is bad. Go cry.
		}
		A = append(A, failHardParseInt(line[0]))
		bVal := failHardParseInt(line[1])
		_, exists := B[bVal]
		if !exists {
			B[bVal] = 1
		} else {
			B[bVal] += 1
		}
	}
	return A, B
}

func main() {
	A, B := parseInput("input.txt")
	sum := 0
	sort.Ints(A)
	for _, v := range A {
		sum += v * B[v]
	}

	fmt.Println("_./*\\._./*\\._./*\\._./*\\._./*\\._")
	fmt.Println(sum)
	fmt.Println("*\\._./*\\._./*\\._./*\\._./*\\._./")
}
