package matrices

import "fmt"

// Definition eines eigenen Datentyps für Matrizen.
// "Matrix ist eine 2D-Float-Slice".
type Matrix [][]float64

// String ist eine Methode auf dem Typ Matrix.
// Gibt eine String-Repräsentation zurück.
func (m Matrix) String() string {
	result := ""

	for _, line := range m {
		result += fmt.Sprintf("%v\n", line)
	}

	return result
}

func Example_matrix() {
	m := Matrix{
		{1, 3},
		{2, 4},
	}

	fmt.Println(m)

	// Output:
	// [[1 3] [2 4]]
}
