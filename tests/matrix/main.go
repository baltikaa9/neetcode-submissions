package matrix

import "fmt"

type MatrixInt [][]int

func (m MatrixInt) Print() {
	for _, r := range m {
		fmt.Println(r)
	}
}