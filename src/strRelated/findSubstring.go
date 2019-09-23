package strrelated

import "fmt"

var inStr []string
var inStrmap map[string]string

// fbs all str
func findstr(in string, words []string) {
	if 1 == len(words) {
		in += words[0]
		inStr = append(inStr, in)
		inStrmap[in] = in
		return
	} else if len(words) == 0 {
		return
	}
	wlen := len(words)

	for i := 0; i < wlen; i++ {
		tmpin := in + words[i]
		nwds := make([]string, wlen)
		copy(nwds, words)
		// nwds = append(nwds[:i], nwds[i+1:]...)
		nwds = nwds[:i+copy(nwds[i:], nwds[i+1:])]
		findstr(tmpin, nwds)
	}
}

// FailedFindSubstring time out.
func FailedFindSubstring(s string, words []string) []int {
	inStr = make([]string, 0)
	inStrmap = make(map[string]string)

	count := len(words)
	ls := len(s)
	wls := count * len(words[0])
	resp := make([]int, 0)
	if count <= 0 {
		return resp
	}
	if ls < wls {
		return resp
	}

	strstr := func(in string, need string) []int {
		nl := len(need)
		res := make([]int, 0)
		for i := 0; i < len(in)-nl+1; i++ {
			if need == in[i:i+nl] {
				res = append(res, i)
			}
		}

		return res
	}

	findstr("", words)
	//respmap := make(map[int]int)
	for _, v := range inStrmap {

		ret := strstr(s, v)
		if len(ret) != 0 {
			resp = append(resp, ret...)
			//respmap[ret] = ret
		}
		fmt.Println(v, ret)
	}

	return resp
}

// CopyMap .
func CopyMap(m map[string]int) map[string]int {
	cp := make(map[string]int)
	for k, v := range m {
		cp[k] = v
	}

	return cp
}

// FindSubstring from s get vaild wlength .
func FindSubstring(s string, words []string) []int {
	inWorldmap := make(map[string]int)
	beStrmap := make(map[int]string)
	resp := make([]int, 0)

	count := len(words)
	ls := len(s)

	if count <= 0 || ls <= 0 {
		return resp
	}

	wlength := len(words[0])
	wls := count * wlength

	if ls < wls {
		return resp
	}
	// insert hash map

	for index := 0; index < count; index++ {
		if _, ok := inWorldmap[words[index]]; !ok {
			inWorldmap[words[index]] = 1
		} else {
			inWorldmap[words[index]]++
		}

	}

	for index := 0; index < ls-wls+1; index++ {
		tmpW := s[index : index+wlength]
		//tmpEd := s[index+wls-wlength : index+wls]
		_, ok := inWorldmap[tmpW]
		//_, ok1 := inWorldmap[tmpEd]&& ok1
		if ok {
			tmpwls := s[index : index+wls]
			beStrmap[index] = tmpwls
		}
	}

	for k, v := range beStrmap {

		//tmpMap := make(map[string]int)
		tmpMap := CopyMap(inWorldmap)
		for index := 0; index < wls; index += wlength {
			tmpW := v[index : index+wlength]
			if _, ok := tmpMap[tmpW]; ok {
				if tmpMap[tmpW] == 1 {
					delete(tmpMap, tmpW)
				} else if tmpMap[tmpW] > 0 {
					tmpMap[tmpW]--
				} else {
					break
				}
			} else {
				break
			}
		}
		if len(tmpMap) == 0 {
			resp = append(resp, k)
		}
	}

	if false {
		FailedFindSubstring(s, words)
	}
	return resp
}
