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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rainArea(arr []int) int {
	stack := make([]int, 0)
	var area int
	for i, num := range arr {
		n := len(stack)
		if n == 0 {
			stack = append(stack, i)
		} else {
			top := stack[n-1]
			topVal := arr[top]

			for num > topVal && len(stack) > 0 {
				stack = stack[:n-1]
				n--

				if len(stack) == 0 {
					break
				}

				top := stack[n-1]
				h := min(num, arr[top]) - topVal
				w := i - top - 1
				area += h * w

				topVal = arr[top]
			}

			stack = append(stack, i)
		}
	}

	return area
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, byte(c))
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if c == ')' && top == '(' {
				stack = stack[:len(stack)-1]
			} else if c == ']' && top == '[' {
				stack = stack[:len(stack)-1]
			} else if c == '}' && top == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[mid] > nums[l] {
				if target >= nums[l] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] < nums[r] {
				if target <= nums[r] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

func heapInsertMax(heap []int, num int) []int {
	h := make([]int, 0, len(heap)+1)
	h = append(h, num)
	h = append(h, heap...)

	for i := 0; i < len(h)/2; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < len(h) && h[i] < h[l] {
			h[i], h[l] = h[l], h[i]
		}
		if r < len(h) && h[i] < h[r] {
			h[i], h[r] = h[r], h[i]
		}
	}

	return h
}

func heapPopMax(h []int) []int {
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]

	for i := 0; i < len(h)/2; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < len(h) && h[i] < h[l] {
			h[i], h[l] = h[l], h[i]
		}
		if r < len(h) && h[i] < h[r] {
			h[i], h[r] = h[r], h[i]
		}
	}
	return h
}

func heapInsertMin(heap []int, num int) []int {
	h := make([]int, 0, len(heap)+1)
	h = append(h, num)
	h = append(h, heap...)

	for i := 0; i < len(h)/2; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < len(h) && h[i] > h[l] {
			h[i], h[l] = h[l], h[i]
		}
		if r < len(h) && h[i] > h[r] {
			h[i], h[r] = h[r], h[i]
		}
	}

	return h
}

func heapPopMin(h []int) []int {
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]

	for i := 0; i < len(h)/2; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < len(h) && h[i] > h[l] {
			h[i], h[l] = h[l], h[i]
		}
		if r < len(h) && h[i] > h[r] {
			h[i], h[r] = h[r], h[i]
		}
	}
	return h
}

type MedianFinder struct {
	maxHeap []int
	minHeap []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: make([]int, 0),
		minHeap: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxHeap) != len(this.minHeap) {
		if len(this.minHeap) > 0 && num > this.minHeap[0] {
			this.minHeap = heapInsertMin(this.minHeap, num)
		} else {
			this.maxHeap = heapInsertMax(this.maxHeap, num)
			top := this.maxHeap[0]
			this.maxHeap = heapPopMax(this.maxHeap)
			this.minHeap = heapInsertMin(this.minHeap, top)
		}
	} else {
		if len(this.minHeap) > 0 && num < this.minHeap[0] {
			this.maxHeap = heapInsertMax(this.maxHeap, num)
		} else {
			this.minHeap = heapInsertMin(this.minHeap, num)
			top := this.minHeap[0]
			this.minHeap = heapPopMin(this.minHeap)
			this.maxHeap = heapInsertMax(this.maxHeap, top)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (len(this.maxHeap)+len(this.minHeap))%2 == 1 {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2
	}
}

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func findDuplicate2(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	p := 0
	for p != slow {
		p = nums[p]
		slow = nums[slow]
	}
	return p
}
