package main

import "fmt"

func RemoveAt[T any](slice []T, i int) []T {
	Tcopy := make([]T, len(slice))
	copy(Tcopy, slice[:])
	if i < 0 || i >= len(Tcopy) {
		// Return the original slice if index is out of bounds
		return Tcopy
	}
	// Remove the element by slicing and concatenating
	return append(Tcopy[:i], Tcopy[i+1:]...)
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	fmt.Println(a, "\n ---")
	for i := 1; i < len(a); i++ {
		b := RemoveAt(a, i)
		fmt.Println(b, a)
	}

}
