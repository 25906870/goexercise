package removeDuplicate

func RemoveDuplicates(nums []int) int {

	//sort.Ints(nums)
	beg := 0
	tag := beg + 1
	length := len(nums)
	for i := 0; i < length && tag < length; i++ {
		if nums[beg] == nums[tag] {
			tag++
		} else {
			nums[beg+1] = nums[tag]
			beg++
			tag++
		}
	}
	return beg + 1
}
