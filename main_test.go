package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	var testLeftArray = []int{3, 4, 2, 1, 3, 3}
	var testRightArray = []int{4, 3, 5, 3, 9, 3}

	t.Run("Day 1 Part 1 test values", func(t *testing.T) {
		dis := Day1p1(testLeftArray, testRightArray)
		t.Log("Test run value", dis)
	})
	t.Run("Day 1 Part 1 values", func(t *testing.T) {
		dis := Day1p1(LoadArrayD1())
		t.Log("Test run value", dis)
	})
	t.Run("Day 1 Part 2 test values", func(t *testing.T) {
		dis := Day1p2(testLeftArray, testRightArray)
		t.Log("Test run value", dis)
	})
	t.Run("Day 1 Part 2 values", func(t *testing.T) {
		dis := Day1p2(LoadArrayD1())
		t.Log("Test run value", dis)
	})

}

func TestDay2(t *testing.T) {

	var testnestedarray = [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	var nestedarray = LoadArrayD2()
	t.Run("Day 2 Part 1 test values", func(t *testing.T) {
		dis := Day2p1(testnestedarray)
		t.Log("Test run value", dis)
	})
	t.Run("Day 2 Part 1 values", func(t *testing.T) {
		dis := Day2p1(nestedarray)
		t.Log("Test run value", dis)
	})
	t.Run("Day 2 Part 2 test values", func(t *testing.T) {
		dis := Day2p2(testnestedarray)
		t.Log("Test run value", dis)
	})
	t.Run("Day 2 Part 2 test values", func(t *testing.T) {
		dis := Day2p2(nestedarray)
		t.Log("Test run value", dis)
	})

}
