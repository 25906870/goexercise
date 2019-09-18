package mergeklist

import (
	"container/heap"
	. "common"
)

 type MHeap []*ListNode

 func (m MHeap) Len() int {
	 return len(m)
 }
 func (m MHeap) Less(i,j int)bool{
	 return m[i].Val < m[j].Val
 }
 func (m MHeap) Swap(i,j int){
	if i < 0 || j < 0 {
		return
	}

	 m[i], m[j] = m[j], m[i]
 }
 func (m *MHeap) Push(x interface{}){
	 *m = append(*m,x.(*ListNode))
 }
 func (m *MHeap) Pop() interface{}{
	 old := *m
	 n := len(old)
	 x := old[n-1]
	 *m = old[0:n-1]
	 return x
 }
 
 func MergeKLists(lists []*ListNode) *ListNode {
    mHeap := make(MHeap,0,1)
    
    for i:=0;i<len(lists);i++{
        if lists[i] == nil{
            continue 
        }
        mHeap = append(mHeap,lists[i])
    }
    
    if len(mHeap) == 0{
		 return nil
	 }
    
    heap.Init(&mHeap)
    
    head := heap.Pop(&mHeap).(*ListNode)
    tmp := head
    if tmp == nil{
        return nil
    }
    if tmp.Next != nil{
        heap.Push(&mHeap,tmp.Next)
    }
    
    for ;len(mHeap) != 0;{
        
        ptr := heap.Pop(&mHeap).(*ListNode)
        if ptr != nil {
            tmp.Next = ptr
            tmp = tmp.Next
            
            if ptr.Next != nil{
                heap.Push(&mHeap,ptr.Next)
            }
        }
        
    }
    
    return head
 }

