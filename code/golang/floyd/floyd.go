package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Matrix       [][]int      // 邻接矩阵
	Vertices     []byte       // 顶点
	Vertex2Index map[byte]int // 记录顶点的下标
	IsDirected   bool
}

func NewGraph(vertices []byte, directed bool) *Graph {
	g := &Graph{
		Matrix:       make([][]int, len(vertices)),
		Vertices:     vertices,
		Vertex2Index: make(map[byte]int),
		IsDirected:   directed,
	}

	for i := 0; i < len(vertices); i++ {
		g.Matrix[i] = make([]int, len(vertices))
		for j := 0; j < len(vertices); j++ {
			// 默认距离为无限大
			g.Matrix[i][j] = math.MaxInt32
			if i == j {
				g.Matrix[i][j] = 0 // 自己到自己的距离为0
			}
		}

		g.Vertex2Index[vertices[i]] = i
	}

	return g
}

func (g *Graph) AddEdge(src, dst byte, weight int) {
	if src == dst {
		return
	}

	srcIndex := g.Vertex2Index[src]
	dstIndex := g.Vertex2Index[dst]
	g.Matrix[srcIndex][dstIndex] = weight
	if !g.IsDirected {
		g.Matrix[dstIndex][srcIndex] = weight
	}
}

func (g *Graph) Floyd() ([][]int, [][]int) {
	var (
		n        int = len(g.Vertices)
		distance     = make([][]int, n)
		path         = make([][]int, n)
	)

	for i := 0; i < n; i++ {
		distance[i] = make([]int, n)
		path[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance[i][j] = g.Matrix[i][j]
			path[i][j] = -1
		}
	}

	for k := 0; k < len(g.Vertices); k++ {
		for i := 0; i < len(g.Vertices); i++ {
			for j := 0; j < len(g.Vertices); j++ {
				if distance[i][j] > distance[i][k]+distance[k][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
					path[i][j] = k // 记录中转站
				}
			}
		}
	}

	return distance, path
}

func getPath(path [][]int, i, j int, s *[]int) {
	k := path[i][j]
	if k == -1 {
		return
	}

	getPath(path, i, k, s) // 获取 i ~ k 的路径
	*s = append(*s, k)
	getPath(path, k, j, s) // 获取 k ~ j 的路径
}

func PrintMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

func main() {
	vertices := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	g := NewGraph(vertices, true)
	g.AddEdge('a', 'b', 10)
	g.AddEdge('b', 'd', 8)
	g.AddEdge('d', 'f', 2)
	g.AddEdge('a', 'c', 6)
	g.AddEdge('c', 'e', 5)
	g.AddEdge('e', 'f', 4)
	g.AddEdge('e', 'd', 1)
	distance, path := g.Floyd()
	PrintMatrix(distance)
	PrintMatrix(path)

	af := make([]int, 0)
	getPath(path, 0, 5, &af)

	fmt.Print("a -> f经过：")
	for i := 0; i < len(af); i++ {
		fmt.Print(fmt.Sprintf("%c ", vertices[af[i]]))
	}
}
