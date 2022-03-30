package main

import (
	"fmt"
	"testing"
)

func TestFindMaxSlice(t *testing.T) {
	q := []int{2, -3, 4, 5, -2}
	q2 := []int{7, -3, 4, -5, -2}

	fmt.Println(findMaxSlice(q))
	fmt.Println(findMaxSlice(q2))
}

func TestCheckRanges(t *testing.T) {
	q := [][]int{{10, 20}, {15, 10}, {20, 15}}
	q2 := [][]int{{10, 15}, {16, 20}, {25, 11}, {21, 17}}

	fmt.Println(checkRanges(len(q), q))
	fmt.Println(checkRanges(len(q2), q2))
}

func TestSumThreeNumberEqualZero(t *testing.T) {
	// q := []int{-1, 0, 1, 2, -1, -4}
	// q2 := []int{0, 0, 0}
	q2 := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Printf("%+v", SumThreeNumberEqualZero(q))
	fmt.Printf("%+v", SumThreeNumberEqualZero(q2))
}

func TestMaxSwap(t *testing.T) {
	fmt.Printf("%+v", MaxSwap(2736))
}
