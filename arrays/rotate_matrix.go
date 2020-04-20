package arrays

import (
	"fmt"
	"strings"
)

type Matrix [][]int

func (m Matrix) String() string {
	buf := strings.Builder{}

	for row := range m {
		_, _ = fmt.Fprintf(&buf, "%v\n", m[row])
	}

	return buf.String()
}

func RotateMatrix(matrix Matrix) {
	n := len(matrix)

	start := 0 // left upper edge

	for m := n; m > 1; m -= 2 {
		for i := 0; i < m-1; i++ {
			p1 := matrix[start][start+i]
			p2 := matrix[start+i][start+m-1]
			p3 := matrix[start+m-1][start+m-i-1]
			p4 := matrix[start+m-i-1][start]
			matrix[start][start+i] = p4
			matrix[start+i][start+m-1] = p1
			matrix[start+m-1][start+m-i-1] = p2
			matrix[start+m-i-1][start] = p3
		}

		start += 1
	}
}
