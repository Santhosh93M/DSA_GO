package main

import "fmt"

type QLNode struct{
	value int
	next *QLNode
}
type QueueList struct{
	head *QLNode
	len int
}
func (q *QueueList) Len() int{
	return q.len
}
func (q *QueueList) IsEmpty() bool{
	return q.len==0
}
func (q *QueueList) Enqueue(val int){
	n := QLNode{value: val}
	if q.IsEmpty(){
		q.head = &n
		q.len++
		return
	}
	ptr := q.head
	for i:=0;i<q.len;i++{
		if ptr.next==nil{
			ptr.next = &n
			q.len++
			return
		}
		ptr = ptr.next
	}
}
func (q *QueueList) Dequeue(){
	if q.IsEmpty(){
		fmt.Println("Queue is Empty..!")
		return
	}
	fmt.Println("The removed element is : ", q.head.value)
	q.head = q.head.next
	q.len--
}
func (q *QueueList) Top(){
	fmt.Println("The top element is : ",q.head.value)
}
func (q *QueueList) Print(){
	n := q.head
	for i:=0;i<q.len;i++{
		fmt.Printf("%d ",n.value)
		n = n.next
	}
	fmt.Println()
}

func main(){
	q := QueueList{}
	q.Enqueue(12)
	q.Enqueue(23)
	q.Enqueue(45)
	q.Enqueue(56)
	q.Enqueue(70)
	q.Print()
	q.Dequeue()
	q.Print()
	q.Top()
}