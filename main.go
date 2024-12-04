package main

import (
	"log"
	"sort"
	"strconv"
)

func Day1(leftArray []int, rightArray []int) int {
	log.Print("Starting Day 1 pt1")

	//sort.Ints(rightArray)
	//sort.Ints(leftArray)
	var distance = 0
	for index := range leftArray {
		leftval := intsplit(leftArray[index])
		rightval := intsplit(rightArray[index])
		sort.Ints(leftval)
		sort.Ints(rightval)
		for i := 0; i < len(leftval); i++ {
			if leftval[i] > rightval[i] {
				distance += leftval[i] - rightval[i]
			} else if leftval[i] < rightval[i] {
				distance += rightval[i] - leftval[i]
			}
		}

	}
	log.Printf("Total for day 1 is: %v", distance)
	return distance
}

func intsplit(in int) []int {
	nt := strconv.Itoa(in)
	result := make([]int, len(nt))
	for i, char := range nt {
		number, _ := strconv.Atoi(string(char))
		result[i] = number
	}
	return result
}
