package lists

import "fmt"

func Example_matrix() {
	m1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(m1[0])
	fmt.Println(m1[0][1])

	//m2 ist eine 3D Liste
	m2 := [][][]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}
	//Output:

}
