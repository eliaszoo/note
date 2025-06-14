package graph

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	num := numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})
	fmt.Println(num)
}

func TestOrangesRotting(t *testing.T) {
	fmt.Println(orangesRotting([][]int{
		{2, 2},
		{1, 1},
		{0, 0},
		{2, 0},
	}))
}

func TestCanFinish(t *testing.T) {
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))

	/*trie.Insert("app")
	trie.Insert("beer")
	trie.Insert("add")
	trie.Insert("jam")
	trie.Insert("rental")

	fmt.Println(trie.Search("app"))
	fmt.Println(trie.Search("add"))
	fmt.Println(trie.Search("applepie"))
	fmt.Println(trie.Search("rest"))
	fmt.Println(trie.Search("jan"))
	fmt.Println(trie.StartsWith("app"))*/
}
