package main

import (
	_ "common"
	"fmt"
	_ "mergeklist"
	. "numRelated"
	. "producer"
	_ "removeDuplicate"
	_ "reverseKGroup"
	. "sortrelated"
	. "strrelated"
	_ "swapPairs"
	"time"
)

func main() {

	RunChCache()

	return
	nums := []int{1, 3, 5}
	tag := 2
	res := SearchInsert(nums, tag)
	fmt.Printf("SearchRange %v", res)
}

func findmain() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	tag := 0
	res := Findtaget(nums, tag)
	fmt.Printf("Findtaget %v", res)
}

func MyPowmain() {

	in := 20.0
	n := 5

	lvp := MyPow(in, n)

	fmt.Printf("LongestValidParentheses %v", lvp)
}

func LongestValidParenthesesmain() {
	s := "(()(()"
	lvp := LongestValidParentheses(s)

	fmt.Printf("LongestValidParentheses %v", lvp)
}

func FindSubstringmain() {
	str1 := "foobarfoobar"
	words1 := []string{"foo", "bar"}
	timestamp := time.Now().Unix()
	//str := "pjzkrkevzztxductzzxmxsvwjkxpvukmfjywwetvfnujhweiybwvvsrfequzkhossmootkmyxgjgfordrpapjuunmqnxxdrqrfgkrsjqbszgiqlcfnrpjlcwdrvbumtotzylshdvccdmsqoadfrpsvnwpizlwszrtyclhgilklydbmfhuywotjmktnwrfvizvnmfvvqfiokkdprznnnjycttprkxpuykhmpchiksyucbmtabiqkisgbhxngmhezrrqvayfsxauampdpxtafniiwfvdufhtwajrbkxtjzqjnfocdhekumttuqwovfjrgulhekcpjszyynadxhnttgmnxkduqmmyhzfnjhducesctufqbumxbamalqudeibljgbspeotkgvddcwgxidaiqcvgwykhbysjzlzfbupkqunuqtraxrlptivshhbihtsigtpipguhbhctcvubnhqipncyxfjebdnjyetnlnvmuxhzsdahkrscewabejifmxombiamxvauuitoltyymsarqcuuoezcbqpdaprxmsrickwpgwpsoplhugbikbkotzrtqkscekkgwjycfnvwfgdzogjzjvpcvixnsqsxacfwndzvrwrycwxrcismdhqapoojegggkocyrdtkzmiekhxoppctytvphjynrhtcvxcobxbcjjivtfjiwmduhzjokkbctweqtigwfhzorjlkpuuliaipbtfldinyetoybvugevwvhhhweejogrghllsouipabfafcxnhukcbtmxzshoyyufjhzadhrelweszbfgwpkzlwxkogyogutscvuhcllphshivnoteztpxsaoaacgxyaztuixhunrowzljqfqrahosheukhahhbiaxqzfmmwcjxountkevsvpbzjnilwpoermxrtlfroqoclexxisrdhvfsindffslyekrzwzqkpeocilatftymodgztjgybtyheqgcpwogdcjlnlesefgvimwbxcbzvaibspdjnrpqtyeilkcspknyylbwndvkffmzuriilxagyerjptbgeqgebiaqnvdubrtxibhvakcyotkfonmseszhczapxdlauexehhaireihxsplgdgmxfvaevrbadbwjbdrkfbbjjkgcztkcbwagtcnrtqryuqixtzhaakjlurnumzyovawrcjiwabuwretmdamfkxrgqgcdgbrdbnugzecbgyxxdqmisaqcyjkqrntxqmdrczxbebemcblftxplafnyoxqimkhcykwamvdsxjezkpgdpvopddptdfbprjustquhlazkjfluxrzopqdstulybnqvyknrchbphcarknnhhovweaqawdyxsqsqahkepluypwrzjegqtdoxfgzdkydeoxvrfhxusrujnmjzqrrlxglcmkiykldbiasnhrjbjekystzilrwkzhontwmehrfsrzfaqrbbxncphbzuuxeteshyrveamjsfiaharkcqxefghgceeixkdgkuboupxnwhnfigpkwnqdvzlydpidcljmflbccarbiegsmweklwngvygbqpescpeichmfidgsjmkvkofvkuehsmkkbocgejoiqcnafvuokelwuqsgkyoekaroptuvekfvmtxtqshcwsztkrzwrpabqrrhnlerxjojemcxel"
	//words := []string{"dhvf", "sind", "ffsl", "yekr", "zwzq", "kpeo", "cila", "tfty", "modg", "ztjg", "ybty", "heqg", "cpwo", "gdcj", "lnle", "sefg", "vimw", "bxcb"}
	// str = words[0]
	// words[0] = str
	rev := FindSubstring(str1, words1)

	for index := 0; index < len(rev); index++ {
		fmt.Println(rev[index])
	}
	fmt.Printf("string %v", time.Now().Unix()-timestamp)
}

/*
func Dividemain() {
	//ret := Divide(-10, 3)
	//-2147483648
	//-1
	num := 2147483648
	ret := DividePro(num, -1)
	fmt.Printf("num %b | %b\n", num, ret)
	fmt.Println(ret)
}

func MystrStrmain() {
	str := "mississippi"
	strs1 := "issipi"
	//ret := RemoveDuplicates(nums)
	ret := MystrStr(str, strs1)

	fmt.Println(ret)
}

func Remvoemain() {
	nums := []int{2, 2, 0, 2, 1, 2, 3, 2}
	//ret := RemoveDuplicates(nums)
	ret := RemoveElement(nums, 2)
	for i := 0; i < ret; i++ {
		fmt.Println(nums[i])
	}
}

func ReverseKGroupmain() {
	nums := []int{1, 2}
	head := new(ListNode)
	tmp := head
	head.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		nt := new(ListNode)
		nt.Val = nums[i]
		tmp.Next = nt
		tmp = tmp.Next

	}
	ReverseKGroup(head, 2)
	//SwapPairs(head)

	return
	iNums := [][]int{{}, {}}
	lists := make([]*ListNode, 0)
	for i := 0; i < len(iNums); i++ {
		head := new(ListNode)
		head.Val = iNums[i][0]
		tmp := head
		for j := 1; j < len(iNums[i]); j++ {
			nt := new(ListNode)
			nt.Val = iNums[i][j]
			tmp.Next = nt
			tmp = tmp.Next

		}
		lists = append(lists, head)
	}

	MergeKLists(lists)

	return
}
*/
