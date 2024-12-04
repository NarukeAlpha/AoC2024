package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadArrayD1() ([]int, []int) {
	var leftArray []int
	var rightArray []int
	file, err := os.Open("d1p1.txt")
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
