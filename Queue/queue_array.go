package main

import "fmt"

type QueueInt struct{
	s []int
}

func (q *QueueInt) Len() int{
	return len(q.s)
}
func (q *QueueInt) IsEmpty() bool{
	return len(q.s)==0
}

func (q * QueueInt) Enqueue(val int){
	q.s = append(q.s, val)
}

func (q *QueueInt) Top(){
	if q.IsEmpty(){
		fmt.Println("Queue Is Empty")
		return
	}
	fmt.Println("The top element is ",q.s[0])
}
func (q *QueueInt) Dequeue(){
	if q.IsEmpty(){
		fmt.Println("Queue is empty...!")
		return
	}
	fmt.Println("The removed element is : ",q.s[0])
	q.s = q.s[1:]
}
func (q *QueueInt) Print(){
	if q.IsEmpty(){
		fmt.Println("queue is empty...!")
	}
	for i:=0;i<q.Len();i++{
		fmt.Printf("%d ",q.s[i])
	}
	fmt.Println()
}

func main(){
	q := QueueInt{}
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