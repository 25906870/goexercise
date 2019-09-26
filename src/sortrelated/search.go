package sortrelated

func Findtaget(nums []int, tag int) int {

	res := -1
	nl := len(nums)
	beg := 0
	for index := (nl + beg) / 2; nums[index] != tag; {
		if nums[index] > tag {
			beg = index
			index = (nl + beg) / 2
		} else {
			nl = index
			index = (nl + beg) / 2
		}

		if nums[index] == tag {
			res = index
			break
		}

		if nl == beg || nl-beg == 1 {
			break
		}
	}

	return res
}

func Search(nums []int, target int) int {
	//res := -1

	lo := 0
	rl := len(nums) - 1
	for lo < rl {
		mid := (lo + rl) / 2

		a1 := 0
		a2 := 0
		a3 := 0

		if nums[0] > target {
			a1 = 1
		}

		if nums[0] > nums[mid] {
			a2 = 1
		}

		if target > nums[mid] {
			a3 = 1
		}

		if a1^a2^a3 == 1 {
			lo = mid + 1
		} else {
			rl = mid
		}
	}

	if lo == rl && nums[lo] == target {
		return lo
	}

	return -1

	// for index := 0; true; index++ {

	// 	if nums[rl] > nums[mid] {
	// 		ns := nums[:mid]
	// 		res = Findtaget(ns, taget)
	// 	}

	// 	if res != -1 {
	// 		return res
	// 	}

	// 	if nums[index] > taget {
	// 		rl = index
	// 		index = (ll + rl) / 2
	// 	} else {
	// 		ll = index
	// 		index = (ll + rl) / 2
	// 	}

	// 	if nums[index] == taget {
	// 		res = index
	// 		break
	// 	}

	// 	if ll == rl || rl-ll == 1 {
	// 		break
	// 	}
	// }
	// // for index < nlen {

	// // }

	// return res
}
