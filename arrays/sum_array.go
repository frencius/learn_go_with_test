package main

import "fmt"

func SumArray(numbers []int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func main() {
	SumArrayTest()
}

func SumArrayTest() {
	fmt.Println("Do SumArray Test")
	fmt.Println("================")
	var n, want int

	fmt.Printf("Input N: ")
	fmt.Scanf("%d", &n)
	items := make([]int, n)

	fmt.Println("Input Array items: ")

	for i := 0; i < n; i++ {
		fmt.Scan(&items[i])
	}

	fmt.Printf("Input want: ")
	fmt.Scanf("%d", &want)

	got := SumArray(items)

	if want != got {
		fmt.Println("FAIL")
	} else {
		fmt.Println("PASSED")
	}
}
