package goleetcode

func product_except_self(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	right := make([]int, size)
	res := make([]int, size)

	left[0] = nums[0]
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i]
	}

	right[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	for i := 0; i < size; i++ {
		if i == 0 {
			res[i] = right[1]
		} else if i == size-1 {
			res[i] = left[size-2]
		} else {
			res[i] = left[i-1] * right[i+1]
		}
	}
	return res

}
