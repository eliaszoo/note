package graph

import "fmt"

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

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		g[i] = make([]int, 0, numCourses)
	}

	for _, p := range prerequisites {
		g[p[0]] = append(g[p[0]], p[1])
	}

	visit := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if len(g[i]) > 0 {
			if !dfsCanFinish(g, i, visit) {
				return false
			}
		}
	}

	return true
}

func dfsCanFinish(g [][]int, i int, visit []int) bool {
	if visit[i] == 1 {
		return false
	}
	if visit[i] == -1 {
		return true
	}

	visit[i] = 1
	for _, idx := range g[i] {
		if !dfsCanFinish(g, idx, visit) {
			return false
		}
	}

	// dfs结束，说明i节点没有环，可以完成课程
	visit[i] = -1
	fmt.Println(visit)
	return true
}

type Trie struct {
	Exist bool
	Next  [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	p := this
	for _, c := range word {
		if p.Next[c-'a'] == nil {
			p.Next[c-'a'] = &Trie{}
		}

		p = p.Next[c-'a']
	}
	p.Exist = true
}

func (this *Trie) Search(word string) bool {
	p := this
	for _, c := range word {
		if p.Next[c-'a'] == nil {
			return false
		}
		p = p.Next[c-'a']
	}
	return p.Exist
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, c := range prefix {
		if p.Next[c-'a'] == nil {
			return false
		}
		p = p.Next[c-'a']
	}
	return true
}
