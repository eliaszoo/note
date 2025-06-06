package array

func Rotate(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[m-1-i][n-j-1]
			matrix[m-1-i][n-j-1] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

func Rotate2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[m-i-1][j] = matrix[m-i-1][j], matrix[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
