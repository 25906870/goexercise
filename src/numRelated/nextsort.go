package numRelated

//5 4 2 3 1 -> 5 4 3 2 1
func nextPermutation(nums []int) {
	numslength := len(nums)

	for index := 0; index < numslength; index++ {

		if nums[index] < nums[index+1] {
			nums[index], nums[index+1] = nums[index+1], nums[index]
			return
		}

	}

}
