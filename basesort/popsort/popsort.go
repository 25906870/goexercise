package popsort



func Popsort(nums []int){
	totallen := len(nums)

	if totallen == 0 {
		return 
	}
	
	for index := 0; index < totallen; index++ {
		
		for j := 0; j < totallen-index -1; j++ {
			
			if nums[j] > nums[j+1] {
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}

		}

	}


}