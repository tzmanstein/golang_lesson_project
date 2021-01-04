package exec

import "fmt"

/**
 *8.12 二维数组应用案例
 */
func Exec01_TraverseMatrix() {
	var scores [3][5]float64

	for i,v := range scores {
		for j,_ := range v {
			fmt.Printf("i=%v, j=%v\n",i,j)
			fmt.Scanln(&scores[i][j])

		}
	}

	fmt.Println("scores=",scores)
}