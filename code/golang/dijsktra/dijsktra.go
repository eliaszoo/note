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

func (g *Graph) Dijsktra(src byte) []int {
	srcIndex, ok := g.Vertex2Index[src]
	if !ok {
		return nil
	}

	distance := make([]int, len(g.Vertices))
	for i := 0; i < len(g.Vertices); i++ {
		distance[i] = g.Matrix[srcIndex][i]
	}

	// 初始化 S 和 U
	S := make([]int, 0, len(g.Vertices))
	S = append(S, srcIndex)

	U := make(map[int]struct{})
	for i := 0; i < len(g.Vertices); i++ {
		if i != srcIndex {
			U[i] = struct{}{}
		}
	}

	for len(U) > 0 {
		min := math.MaxInt32
		minIndex := -1
		for i, _ := range U {
			if distance[i] < min {
				min = distance[i]
				minIndex = i
			}
		}

		delete(U, minIndex)
		S = append(S, minIndex)

		for i := 0; i < len(g.Vertices); i++ {
			// 从起点到选中点的最短距离 + 选中点到目标点的距离
			if distance[minIndex]+g.Matrix[minIndex][i] < distance[i] {
				distance[i] = distance[minIndex] + g.Matrix[minIndex][i]
			}
		}
	}

	return distance
}

func PrintMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

func main() {
	g := NewGraph([]byte{'a', 'b', 'c', 'd', 'e', 'f'}, false)
	g.AddEdge('a', 'b', 10)
	g.AddEdge('a', 'd', 4)
	g.AddEdge('b', 'c', 8)
	g.AddEdge('b', 'd', 2)
	g.AddEdge('b', 'e', 6)
	g.AddEdge('c', 'd', 15)
	g.AddEdge('c', 'e', 1)
	g.AddEdge('c', 'f', 5)
	g.AddEdge('d', 'e', 6)
	g.AddEdge('e', 'f', 12)

	PrintMatrix(g.Matrix)

	distance := g.Dijsktra('a')
	fmt.Println(distance)
}
