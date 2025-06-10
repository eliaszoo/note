package graph

func numIslands(grid [][]byte) int {
	visit := make([][]bool, len(grid))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(grid[i]))
	}

	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !visit[i][j] {
				dfs(grid, visit, i, j, &count)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, visit [][]bool, i, j int, count *int) {
	m := len(grid)
	n := len(grid[0])

	visit[i][j] = true
	if j < n-1 && grid[i][j+1] == '1' && !visit[i][j+1] {
		dfs(grid, visit, i, j+1, count)
	}
	if j > 0 && grid[i][j-1] == '1' && !visit[i][j-1] {
		dfs(grid, visit, i, j-1, count)
	}
	if i < m-1 && grid[i+1][j] == '1' && !visit[i+1][j] {
		dfs(grid, visit, i+1, j, count)
	}
	if i > 0 && grid[i-1][j] == '1' && !visit[i-1][j] {
		dfs(grid, visit, i-1, j, count)
	}
}

func orangesRotting(grid [][]int) int {
	return bfsOrangesRotting(grid)
}

func bfsOrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := make([][2]int, 0)
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}
	if count == 0 {
		return 0
	}

	var times int
	for len(queue) > 0 {
		size := len(queue)
		for x := 0; x < size; x++ {
			head := queue[0]
			queue = queue[1:]

			tmpI := head[0]
			tmpJ := head[1]

			if tmpI > 0 && grid[tmpI-1][tmpJ] == 1 {
				grid[tmpI-1][tmpJ] = 2
				count--
				queue = append(queue, [2]int{tmpI - 1, tmpJ})
			}
			if tmpI < m-1 && grid[tmpI+1][tmpJ] == 1 {
				grid[tmpI+1][tmpJ] = 2
				count--
				queue = append(queue, [2]int{tmpI + 1, tmpJ})
			}
			if tmpJ > 0 && grid[tmpI][tmpJ-1] == 1 {
				grid[tmpI][tmpJ-1] = 2
				count--
				queue = append(queue, [2]int{tmpI, tmpJ - 1})
			}
			if tmpJ < n-1 && grid[tmpI][tmpJ+1] == 1 {
				grid[tmpI][tmpJ+1] = 2
				count--
				queue = append(queue, [2]int{tmpI, tmpJ + 1})
			}
		}
		times++
	}
	if count > 0 {
		return -1
	}
	return times - 1
}
