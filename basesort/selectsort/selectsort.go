package selectsort

func Selectsort(nums []int){

	totallen := len(nums)

	if totallen == 0 {
		return 
	}
	var min,index int
	for i := 0; i < totallen; i++ {
		
		min = nums[i]
		index = i
		for j := index + 1; j < totallen; j++ {
			
			if min > nums[j] {
				min = nums[j]
				index = j
			}
		}

		nums[i],nums[index] = nums[index],nums[i]
	}

}