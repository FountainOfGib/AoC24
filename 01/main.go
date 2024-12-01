package main

// asda

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func parseInput(fileName string) ([]int, []int) {
	A, B := []int{}, []int{}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err opening", fileName)
		os.Exit(1)
		// return {}, {}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return A, B
}

func main() {
	A, B := parseInput("input.txt")
	diff := []int{}
	sort.Ints(A)
	sort.Ints(B)
	for i := 0; i < len(A); i++ {
		delta := A[i] - B[i]
		diff = append(diff, max(delta, delta*-1))
	}
	fmt.Println("_./*\\._./*\\._./*\\._./*\\._./*\\._")
	fmt.Println(diff)
	fmt.Println("*\\._./*\\._./*\\._./*\\._./*\\._./")
}
