package removeDuplicate

func RemoveElement(nums []int, val int) int {

	tag := -1
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] == val {
			if tag == -1 { // nums = append(nums[0:i],nums[i+1:]...)
				tag = i
			}
		} else if tag != -1 {
			nums[tag] = nums[i]
			tag++
		}
	}
	if tag == -1 {
		return length
	}
	return tag
}
