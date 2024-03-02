package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
)

func quic(arr []int, begin, end int) {
	if begin >= end {
		return
	}

	idx := quicMove(arr, begin, end)
	quic(arr, begin, idx-1)
	quic(arr, idx+1, end)
}

// 其中一种实现方式
func quicMove(arr []int, begin, end int) int {
	mid := arr[begin] // 取出中位数之后，相当于begin/i这个位置空出来了
	i := begin
	j := end
	for i < j {
		for ; j > i; j-- {
			if arr[j] < mid {
				arr[i] = arr[j] // 执行完这句之后，相当于j的位置又空出来了，所以不是交换而是赋值
				break
			}
		}
		for ; i < j; i++ {
			if arr[i] > mid {
				arr[j] = arr[i] // 同上
				break
			}
		}
	}
	arr[i] = mid // 最后 i == j 把mid放在这个位置
	return i
}

func main() {
	arr := make([]int, 0, 100)
	arr2 := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		v := rand.IntN(1 << 16)
		arr = append(arr, v)
		arr2 = append(arr2, v)
	}
	//fmt.Println(arr)
	quic(arr, 0, len(arr)-1)

	sort.Ints(arr2)
	for i := 0; i < 100; i++ {
		if arr[i] != arr2[i] {
			fmt.Println("ERROR")
			os.Exit(1)
		}
	}

	fmt.Println("same")
	//fmt.Println(arr)
}
