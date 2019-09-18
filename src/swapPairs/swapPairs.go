package swapPairs

import(
	. "common"
)

// 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗 :2.1 MB, 在所有 Go 提交中击败了22.48%的用户

func SwapPairs(head *ListNode) *ListNode {
   
    beg := head
    var tag *ListNode
    if head != nil && head.Next != nil {
        tag = head.Next
    }
    
    
    var reshead *ListNode
    var preptr *ListNode
    reshead = tag
    //tmp := reshead
    for ; tag!= nil; {
        beg.Next = tag.Next
        tag.Next = beg
        if preptr != nil {
            preptr.Next = tag  
        }

        preptr = beg
        if beg.Next != nil{
            beg = beg.Next
            tag = beg.Next
        }else{
            tag = nil 
        }

    }
    
    return reshead
}