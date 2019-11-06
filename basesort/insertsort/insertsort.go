package insertsort

func Insertsort(nums []int){

	totallen := len(nums)

	if totallen == 0 {
		return 
	}

	var key,pos int

	for i := 0; i < totallen; i++ {
		
		key = nums[i]
		pos = i

		for pos > 0 && nums[pos - 1] < key {
			nums[pos] = nums[pos - 1]// 挨个移动
			pos --
		}
		nums[pos] = key

	}


}