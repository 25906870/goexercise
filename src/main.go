package main

import (
	. "common"
	"fmt"
	. "mergeklist"
	. "numRelated"
	. "removeDuplicate"
	. "reverseKGroup"
	. "strRelated"
	_ "swapPairs"
)

func main() {
	//ret := Divide(-10, 3)

	ret := DividePro(-5, -1)
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
