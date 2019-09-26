package numRelated

func MyPow(x float64, n int) float64 {
	res := 1.0

	for index := n; index != 0; index /= 2 {
		if index%2 != 0 {
			res *= x
		}
		x *= x
	}

	if n < 0 && res != 0 {
		return 1 / res
	}

	return res
}
