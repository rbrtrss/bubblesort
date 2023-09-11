package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}

func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

func printSlice(slice []int, numItems int) {
	if numItems > len(slice) {
		numItems = len(slice)
	}
	s := slice[:numItems]
	fmt.Printf("%d \n", s)
}

// func isSorted(slice []int) bool {
// 	sorted := true
// 	for i := 0; i < (len(slice) - 1); i++ {
// 		if slice[i] > slice[i+1] {
// 			sorted = false
// 			break
// 		}
// 	}
// 	return sorted
// }

func checkSorted(slice []int) {
	sorted := true
	msg := "The slice is sorted"
	for i := 0; i < (len(slice) - 1); i++ {
		if slice[i] > slice[i+1] {
			sorted = false
			break
		}
	}
	if !sorted {
		msg = "The slice is NOT sorted!"
	}
	fmt.Println(msg)
}

func bubbleSort(slice []int) {
	countSwtiches := 0
	for i := 0; i < (len(slice) - 1); i++ {
		prev := slice[i]
		next := slice[i+1]
		if prev > next {
			slice[i] = next
			slice[i+1] = prev
			countSwtiches++
		}
	}
	if countSwtiches > 0 {
		bubbleSort(slice[:(len(slice) - 1)])
	}
}
