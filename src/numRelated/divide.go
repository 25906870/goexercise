package numRelated

func abs(num int) int {
	return num ^ (-1 >> 31) - (-1 >> 31)
}

const INT_MAX = int(^uint32(0) >> 1)

func Divide(dividend int, divisor int) int {

	tag := 0
	flag := 1
	if dividend < 0 {
		flag = -1
		dividend = abs(dividend)
		// tmp := 10
		// tmp = tmp ^ (-1 >> 31)
		// tmp = tmp - (-1 >> 31)
		// tmp = tmp ^ -1 + 1
		// fmt.Printf("%v,%b,%b,%b", tmp)
	}
	if divisor < 0 {
		flag = 0 - flag
		divisor = abs(divisor)
	}

	dividend = dividend - divisor

	for dividend > 0 {
		tag++
		dividend = dividend - divisor
	}

	if tag >= INT_MAX {
		return INT_MAX
	}

	return tag

}
