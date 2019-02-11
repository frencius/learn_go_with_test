package main

import (
	"testing"
)

func TestSumArray(t *testing.T) {

	items := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumArray(items)

	if want != got {
		t.Errorf("got %d want %d given %v", got, want, items)
	}
}

func TestSumArray_2(t *testing.T) {
	t.Run("run another test case", func(t *testing.T) {
		items := []int{1, 2, 3, 4}
		want := 10
		got := SumArray(items)

		if got != want {
			t.Errorf("got %d want %d given %d", got, want, items)
		}
	})
}

func TestSumArray_3(t *testing.T) {
	assertCorrectNumbers := func(t *testing.T, want, got int, items []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, items)
		}
	}

	t.Run("try another test cases", func(t *testing.T) {
		items := []int{1, 2, 3, 2}
		want := 8
		got := SumArray(items)
		assertCorrectNumbers(t, want, got, items)
	})
}
