package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func LoadArrayD1() ([]int, []int) {
	var leftArray []int
	var rightArray []int
	file, err := os.Open("./Input/d1.txt")
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		fields := strings.Fields(line)
		lv, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Panicf("Error converting left side of array to int: %v", err)
		}
		leftArray = append(leftArray, lv)
		rv, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Panicf("Error converting right side of array to int: %v", err)
		}
		rightArray = append(rightArray, rv)

	}
	return leftArray, rightArray
}
func Day1p1(leftArray []int, rightArray []int) int {

	sort.Ints(rightArray)
	sort.Ints(leftArray)
	var distance = 0
	for index := range leftArray {
		if leftArray[index] > rightArray[index] {
			distance += leftArray[index] - rightArray[index]
		} else if leftArray[index] < rightArray[index] {
			distance += rightArray[index] - leftArray[index]
		}
	}
	return distance
}
func Day1p2(leftArray []int, rightArray []int) int {
	sort.Ints(leftArray)
	sort.Ints(rightArray)
	var similarity = 0
	for index := range leftArray {
		count := 0
		for _, value := range rightArray {
			if leftArray[index] == value {
				count++
			}
		}
		sim := leftArray[index] * count
		similarity += sim
	}
	return similarity
}

func LoadArrayD2() [][]int {
	var nestedarray [][]int
	file, err := os.Open("./Input/d2.txt")
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var arrofval []int
		line := sc.Text()
		fields := strings.Fields(line)
		for _, val := range fields {
			valint, err := strconv.Atoi(val)
			if err != nil {
				log.Panicf("Error converting value to int: %v", err)
			}
			arrofval = append(arrofval, valint)
		}
		nestedarray = append(nestedarray, arrofval)
	}
	return nestedarray
}

func withinrangeminus(i int, x int) bool {
	switch {
	case i-1 == x:
		return true
	case i-2 == x:
		return true
	case i-3 == x:
		return true
	default:
		return false

	}
}
func withinrangemas(i int, x int) bool {
	switch {
	case i+1 == x:
		return true
	case i+2 == x:
		return true
	case i+3 == x:
		return true
	default:
		return false
	}
}

func Day2p1(list [][]int) int {
	var isSafe = 0
	for _, value := range list {
		if withinrangemas(value[0], value[1]) {
			for i, v := range value {
				if i >= len(value)-1 {
					isSafe++
					break
				}
				if !withinrangemas(v, value[i+1]) {
					break
				}
			}
		}
		if withinrangeminus(value[0], value[1]) {
			for i, v := range value {
				if i >= len(value)-1 {
					isSafe++
					break
				}
				if !withinrangeminus(v, value[i+1]) {
					break
				}
			}
		}

	}
	return isSafe

}

func Day2p2(list [][]int) int {
	var isSafe = 0
	for _, value := range list {
		tolerance := 0
		if withinrangemas(value[0], value[1]) {
			for i, v := range value {
				if i >= len(value)-1 {
					isSafe++
					break
				}
				if !withinrangemas(v, value[i+1]) {
					tolerance++
				}
				if tolerance >= 1 {
					break
				}
			}
		}
		if withinrangeminus(value[0], value[1]) {
			for i, v := range value {
				if i >= len(value)-1 {
					isSafe++
					break
				}
				if !withinrangeminus(v, value[i+1]) {
					tolerance++
				}
				if tolerance >= 1 {
					break
				}
			}
		}

	}
	return isSafe

}
