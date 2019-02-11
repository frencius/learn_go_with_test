package array_slices

import "testing"

func TestSum(t *testing.T) {

	t.Run("using array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d, but want %d, given %v", got, want, numbers)
		}
	})
}
