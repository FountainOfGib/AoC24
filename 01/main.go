package main

// asda

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

func parseInput(fileName string) ([]int, []int) {
	A, B := []int{}, []int{}

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
		B = append(B, failHardParseInt(line[1]))
		fmt.Print("\n")
	}
	return A, B
}

func main() {
	A, B := parseInput("input.txt")
	diff := []int{}
	sort.Ints(A)
	sort.Ints(B)
	for i, _ := range A {
		delta := A[i] - B[i]
		diff = append(diff, max(delta, delta*-1))
	}
	fmt.Println("_./*\\._./*\\._./*\\._./*\\._./*\\._")
	fmt.Println(diff)
	fmt.Println("*\\._./*\\._./*\\._./*\\._./*\\._./")
	sum := 0
	for _, num := range diff {
		sum += num
	}
	fmt.Println(sum)
}
