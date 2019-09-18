package reverseKGroup

import(
	. "common"
)


func ReverseKGroup(head *ListNode, k int) *ListNode {
	var ret *ListNode
	
	if k <= 0 {
		return head
	}

	que := make([]*ListNode, k)
	var ptr *ListNode
	for ;head != nil;{
		
		for i := 0; i < k ; i++ {
			que[i] = head
			if head != nil {
				head = head.Next
                if head == nil{
					if i < k -1 {
						if ptr != nil{
							ptr.Next = que[0]
						 }else{
							 ret = que[0]
						 }
						 return ret
					}		    
                }
			}
		}

		
		for j := 0; j < k; j++ {
			if ret == nil {
				ret = que[k-j-1]
				ptr = ret
			}else{
				ptr.Next = que[k-j-1]
				ptr = ptr.Next
				ptr.Next = nil
			}
		}
	}
	
	return ret 
}