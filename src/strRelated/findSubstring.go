package strRelated

var inStr []string

func findstr(in string, words []string) {
	if 1 == len(words) {
		in += words[0]
		inStr = append(inStr, in)
		return
	}

	findstr(in, append(words[:0], words[:1]...))
	findstr(in, append(words[:1], words[:2]...))
	//findstr(length,)
}

func FindSubstring(s string, words []string) []int {

	count := len(words)
	ls := len(s)
	wls := count * len(words[0])

	if ls < wls {
		return nil
	}

	for index := 0; index < count; index++ {

	}

	return nil
}
