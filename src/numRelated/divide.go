package numRelated

import "fmt"

func abs(num int) int {
	return num ^ -1 - (-1 >> 31)
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

func DividePro(dividend int, divisor int) int {

	tag := 0
	flag := 1
	if dividend < 0 {
		flag = -1
		dividend = abs(dividend)
	}
	if divisor < 0 {
		flag = 0 - flag
		divisor = abs(divisor)
	}

	count := 1
	if dividend > divisor {
		tmp := divisor
		for dividend > tmp {
			tmp <<= 1
			if dividend >= tmp {
				divisor = tmp
				count++
			}
		}
	}

	for index := 0; index < count; index++ {
		if dividend < divisor {
			tag <<= 1
		}else if dividend != 0  {
			dividend = dividend - divisor
			tag <<= 1
			tag |= 0x01
		}
		divisor >>= 1
	}

	if flag == -1 {
		return -tag
	}

	if tag >= INT_MAX  {
        return INT_MAX 
    }

	fmt.Printf("%b\n", tag)
	return tag

}
