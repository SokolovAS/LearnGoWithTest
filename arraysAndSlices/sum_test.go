package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		expected := 6

		if expected != got {
			t.Errorf("Expected %d got %d", expected, got)
		}
	})
}

func checkSums(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v got %v", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("take the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{4, 5})
		expected := []int{2, 5}

		checkSums(t, got, expected)
	})
	t.Run("take the sum of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		expected := []int{0, 11}

		checkSums(t, got, expected)
	})
}
