package greedy

import "fmt"

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	cover := nums[0]
	old := cover
	step := 1
	i := 1
	for i < len(nums) {
		if cover >= len(nums)-1 {
			break
		}

		for ; i <= old; i++ {
			if cover >= len(nums)-1 {
				break
			}
			if i+nums[i] > cover {
				cover = i + nums[i]
			}
		}
		step++
		old = cover
	}

	return step
}

func productExceptSelf(nums []int) []int {
	a := make([]int, len(nums))
	b := make([]int, len(nums))
	a[0] = 1
	for i := 1; i < len(nums); i++ {
		a[i] = a[i-1] * nums[i-1]
	}

	fmt.Println("a:", a)
	b[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		b[i] = b[i+1] * nums[i+1]
	}

	fmt.Println("b:", b)
	for i := 0; i < len(nums); i++ {
		a[i] *= b[i]
	}

	return a
}

// 分发糖果 https://leetcode.cn/problems/candy/submissions/633517775/?envType=study-plan-v2&envId=top-interview-150
func candy(ratings []int) int {
	a := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		a[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			a[i] = a[i-1] + 1
		}
	}
	fmt.Println(a)
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && a[i] <= a[i+1] {
			a[i] = a[i+1] + 1
		}
	}
	fmt.Println(a)
	var count int
	for i := 0; i < len(a); i++ {
		count += a[i]
	}
	return count
}

// 缺失的第一个正整数：https://leetcode.cn/problems/first-missing-positive/?envType=study-plan-v2&envId=top-100-liked
// 思路：将数组原地hash
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 这里要使用for循环，要保证交换后的数据能够得到处理
		for nums[i] != i+1 {
			// 如果数据不符合要求，或者目标位置已经有正确的相同的值，则跳出循环
			if nums[i] <= 0 || nums[i] > len(nums) || nums[i] == nums[nums[i]-1] {
				break
			}

			idx := nums[i] - 1
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)
	m := len(matrix)
	n := len(matrix[0])

	x, y := 0, 0
	for m-1 > x && n-1 > y {
		i, j := x, y
		for ; j < n-1; j++ {
			ret = append(ret, matrix[i][j])
		}

		for ; i < m-1; i++ {
			ret = append(ret, matrix[i][j])
		}

		for ; j > y; j-- {
			ret = append(ret, matrix[i][j])
		}

		for ; i > x; i-- {
			ret = append(ret, matrix[i][j])
		}
		m--
		n--
		x++
		y++
	}
	if x == m {
		for i := y; i < n; i++ {
			ret = append(ret, matrix[x][i])
		}
	} else if y == n {
		for i := x; i < m; i++ {
			ret = append(ret, matrix[i][y])
		}
	} else {
		ret = append(ret, matrix[x-1][y-1])
	}

	return ret
}

func spiralOrder2(matrix [][]int) []int {
	ret := make([]int, 0)
	u := 0
	d := len(matrix) - 1
	l := 0
	r := len(matrix[0]) - 1

	for {
		for i := l; i <= r; i++ {
			ret = append(ret, matrix[u][i])
		}
		u++
		if u > d {
			break
		}

		for i := u; i <= d; i++ {
			ret = append(ret, matrix[i][r])
		}
		r--
		if l > r {
			break
		}

		for i := r; i >= l; i-- {
			ret = append(ret, matrix[d][i])
		}
		d--
		if u > d {
			break
		}

		for i := d; i >= u; i-- {
			ret = append(ret, matrix[i][l])
		}
		l++
		if l > r {
			break
		}

	}

	return ret
}

func maxSlidingWindow(nums []int, k int) []int {
	var ret []int
	if len(nums) < k {
		return ret
	}

	queue := make([]int, 0)
	push := func(val int) {
		tail := len(queue) - 1
		for tail >= 0 && val > queue[tail] {
			queue = queue[:tail]
			tail--
		}
		queue = append(queue, val)
	}

	for i := 0; i < len(nums); i++ {
		if i >= k {
			if len(queue) > 0 && queue[0] == nums[i-k] {
				queue = queue[1:]
			}
		}

		push(nums[i])
		if i >= k-1 {
			ret = append(ret, queue[0])
		}
	}

	return ret
}

func getMax(nums []int) int {
	max := nums[0]
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index
}
