package sortrelated

func SearchInsert(nums []int, target int) int {

	left := 0
	right := len(nums)
	tag := 0
	for {

		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid + 1
		}

		if mid == 0 || left >= right {
			if mid == 0 && right == 0 {
				return 0
			}
			tag = mid
			break
		}

	}

	if nums[tag] < target {
		if tag == len(nums)-1 {
			return tag + 2
		}

	}
	return tag + 1

}
