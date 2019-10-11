package sortrelated

func SearchRange(nums []int, target int) []int {

	res := []int{-1, -1}
	rl := len(nums)
	ll := 0
	tmpindex := -1
	tmpmid := -1
	for rl > ll {
		mid := (rl + ll) / 2

		if nums[mid] > target {
			rl = mid
		} else {
			ll = mid
		}

		if nums[mid] == target {
			tmpindex = mid
			break
		}
		if tmpmid != mid {
			tmpmid = mid
		} else {
			break
		}
	}

	if tmpindex != -1 {
		ri := tmpindex
		li := tmpindex

		lfinish := 1
		rfinish := 1
		if li == 0 {
			res[0] = li
			lfinish--
		}
		if ri == len(nums)-1 {
			res[1] = ri
			rfinish--
		}
		for true {
			if li >= 0 {

				if lfinish == 1 {
					if nums[li] == target {
						res[0] = li
						li--
					} else {
						lfinish--
					}

				}

			} else {
				res[0] = li + 1
				lfinish = 0
			}

			if ri < len(nums) {

				if rfinish == 1 {
					if nums[ri] == target {
						res[1] = ri
						ri++
					} else {
						rfinish--
					}
				}

			} else {
				res[1] = ri - 1
				rfinish = 0
			}

			if rfinish == 0 && lfinish == 0 || (li < 0 && ri > len(nums)-1) {
				break
			}
		}

	}
	return res
}
