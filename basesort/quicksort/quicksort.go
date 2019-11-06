package quicksort


func QuicksortRecusive(nums []int)[]int{

	totallen := len(nums)

	if totallen == 0 {
		return nil
	}

	base := nums[0]

	left := make([]int,0)
	right := make([]int,0)

	for i := 1; i < totallen; i++ {
		
		if nums[i] < base {
			left = append(left,nums[i])
		}else{
			right = append(right,nums[i])
		}

	}

	left = QuicksortRecusive(left)
	right = QuicksortRecusive(right)

	res := make([]int,0)
	res = append(res,left...)
	res = append(res,base)
	res = append(res,right...)
	return res
}