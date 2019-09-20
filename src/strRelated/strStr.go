package strRelated

func MystrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	hlen := len(haystack)
	nlen := len(needle)

	if hlen < nlen {
		return -1
	}

	for i := 0; i < hlen-nlen+1; i++ {
		if haystack[i] == needle[0] {
			j := 1
			tmp := i + 1
			for j < nlen && tmp < hlen {
				if haystack[tmp] != needle[j] {
					break
				} else {
					tmp++
					j++
				}
			}
			if j == nlen {
				return i
			}
		}
	}
	return -1
}

func MystrStrPro(haystack string, needle string) int {

	nlen := len(needle)

	for i := 0; i < len(haystack)-nlen+1; i++ {
		if needle == haystack[i:i+nlen] {
			return i
		}
	}

	return -1
}
