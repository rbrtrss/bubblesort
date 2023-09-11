package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// fmt.Println("Implementing a Bubble Sort algorithm")
	// fmt.Printf("%x \n", makeRandomSlice(5, 10))
	s1 := makeRandomSlice(50000, 10)
	printSlice(s1, 5)
	bubbleSort(s1[:5])
	printSlice(s1, 5)
	// checkSorted([]int{1, 2, 3})
	// checkSorted([]int{3, 2, 1})
	// checkSorted([]int{1, 2, 3, 4, 5})
	// checkSorted([]int{4, 2, 4})
	// checkSorted([]int{2, 2, 3})
}

func makeRandomSlice(numItems, max int) []int {
	s := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		s[i] = rand.Intn(max)
	}
	return s
}

func printSlice(slice []int, numItems int) {
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
