package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(size, valueRange int) []int {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, 0, size)
	for i := 0; i < size; i++ {
		numbers = append(numbers, rand.Intn(valueRange))
	}

	return numbers
}

func partition(slice []int) int {
	lowIndex, highIndex := 0, len(slice)-1
	pivotIndex := highIndex / 2
	pivotValue := slice[pivotIndex]

	slice[pivotIndex], slice[highIndex] = slice[highIndex], slice[pivotIndex]

	for i := range slice {
		if slice[i] < pivotValue {
			slice[lowIndex], slice[i] = slice[i], slice[lowIndex]
			lowIndex += 1
		}
	}

	slice[lowIndex], slice[highIndex] = slice[highIndex], slice[lowIndex]

	return lowIndex
}

func quicksort(slice []int, done chan int) {
	if len(slice) <= 1 {
		done <- 1
		return
	}

	pivotIndex := partition(slice)

	c := make(chan int)
	go quicksort(slice[:pivotIndex], c)
	go quicksort(slice[pivotIndex+1:], c)

	for i := 0; i < 2; i++ {
		<-c
	}

	done <- 1
	return
}

func main() {
	numbers := generateRandomNumbers(10, 100)
	fmt.Println("Before:", fmt.Sprint(numbers))

	done := make(chan int)
	beginTime := time.Now()

	go quicksort(numbers, done)
	<-done

	endTime := time.Since(beginTime)
	fmt.Println(endTime)

	fmt.Println("After :", fmt.Sprint(numbers))
}
