package gamelogic

import "fmt"

func TicTacToe() {
	fmt.Println("in tic tac toe")
	var rows int = 3
	var cols int = 3
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	matrix[0][0] = 0
	matrix[0][1] = 1
	matrix[1][2] = 0
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 1
	matrix[2][0] = 1
	matrix[2][1] = 1
	matrix[2][2] = 0
	fmt.Println(matrix)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (matrix[i][0] == matrix[i][1] && matrix[i][1] == matrix[i][2]) || (matrix[0][j] == matrix[1][j] && matrix[1][j] == matrix[2][j]) {
				fmt.Println("win")
			}
		}
	}
}
