package main

import (
	"fmt"
	"testing"
)

var (
	q1 = []int{1, 5, -4, 2, 0, -8, 20, 3}
)

func TestSort(t *testing.T) {

	fmt.Printf("%+v\n", InsertionSort(q1))
	fmt.Printf("%+v\n", BubbleSort(q1))
	fmt.Printf("%+v\n", SelectionSort(q1))
	q1 = []int{1, 5, -4, 2, 0, -8, 20, 3}
	fmt.Printf("%+v\n", QuickSort(q1))
	fmt.Printf("%+v\n", HeapSort(q1))
	fmt.Printf("%+v\n", MergeSort(q1))
	q1 = []int{1, 5, -4, 2, 0, -8, 20, 3}
	fmt.Printf("%+v\n", ShellSort(q1))

}
