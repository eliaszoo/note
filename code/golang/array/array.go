package array

import "sort"

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

func findKthLargest(nums []int, k int) int {
	return quic(nums, 0, len(nums)-1, k)
}

func quic(nums []int, l, r, k int) int {
	idx := quickMove(nums, l, r)
	if idx+1 == k {
		return nums[idx]
	} else if idx+1 > k {
		return quic(nums, l, idx-1, k)
	} else {
		return quic(nums, idx+1, r, k)
	}
}

func quickMove(nums []int, l, r int) int {
	flag := nums[l]

	for l < r {
		for ; l < r; r-- {
			if nums[r] > flag {
				nums[l] = nums[r]
				break
			}
		}

		for ; l < r; l++ {
			if nums[l] < flag {
				nums[r] = nums[l]
				break
			}
		}
	}

	nums[l] = flag
	return l
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		target := 0 - nums[i]
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// 调过后面连续的重复值
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}
