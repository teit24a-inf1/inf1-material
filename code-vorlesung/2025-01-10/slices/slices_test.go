package slices

import "fmt"

func Example_slice1D() {
	l1 := []int{}

	fmt.Println(len(l1))

	// Output:
	// 0
}

func Example_slice2D() {
	l2 := [][]int{
		{1},       // l2[0]
		{2, 3, 4}, // l2[1]
	}

	fmt.Println(len(l2))
	fmt.Println(len(l2[0]))
	fmt.Println(len(l2[1]))

	// Output:
	// 2
	// 1
	// 3
}

func Example_slice3D() {
	l3 := [][][]int{
		{
			{1}, // l3[0][0]
			{
				42, // l3[0][1][0]
				23,
			}, // l3[0][1]
		}, // l3[0]
		{
			{2, 3, 4},
			{5, 6, 7},
			{8, 9},
		}, // l3[1]
	}

	fmt.Println(len(l3))
	fmt.Println(len(l3[0]))
	fmt.Println(len(l3[1]))
	fmt.Println(len(l3[0][0]))

	// Output:
	// 2
	// 2
	// 3
	// 1
}
