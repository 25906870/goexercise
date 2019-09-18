package main

import (
	. "common"
	_ "swapPairs"
	. "mergeklist"
	. "reverseKGroup"
)

func main(){
	nums := []int{1,2}
	head := new(ListNode)
	tmp := head
	head.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		nt := new(ListNode)
		nt.Val = nums[i]		
		tmp.Next = nt
		tmp = tmp.Next	
		
	}
	ReverseKGroup(head,2)
	//SwapPairs(head)

	return 
	iNums := [][]int{{},{}}
	lists := make([]*ListNode, 0)
	for i:=0;i<len(iNums);i++{
		head := new(ListNode)
		head.Val = iNums[i][0]
		tmp := head
		for j := 1; j < len(iNums[i]); j++ {
			nt := new(ListNode)
			nt.Val = iNums[i][j]		
			tmp.Next = nt
			tmp = tmp.Next	
			
		}
		lists = append(lists,head)
	}

	MergeKLists(lists)

	return 
}