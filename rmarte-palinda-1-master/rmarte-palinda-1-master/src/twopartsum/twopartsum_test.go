package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// odd length
func TestSumConcurrentlySumsOddArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expected := 66

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// negative numbers
func TestSumConcurrentlySumsNegativeNumbers(t *testing.T) {
	arr := []int{-1, 1, -2, 5, 4, -4, -5}
	expected := -2

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// TODO add at least two more test cases!
