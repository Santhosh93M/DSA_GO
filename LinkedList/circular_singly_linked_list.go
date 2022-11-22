package main

import "fmt"

// import "fmt";import "errors"
// Node represents a node of linked list

type CNode struct {
	value int
	next *CNode
}

// LinkedList represents a linked list
type CirSingLinkedList struct {
	head *CNode
	tail *CNode
	len int
}

func (cl *CirSingLinkedList) IsEmpty() bool{
	if(cl.len==0){
		return true
	}else{
		return false
	}
}
func (cl *CirSingLinkedList) Addfirst(v int){
	n := CNode{value: v}
	if(cl.len==0){
		n.next = &n
		cl.head = &n
		cl.tail = &n
	}else{
		n.next = cl.head
		cl.head = &n
		cl.tail.next = &n

	}
	cl.len++
	return
}
func (cl *CirSingLinkedList) AddLast(v int){
	n := CNode{value: v}
	if(cl.len==0){
		n.next = &n
		cl.head = &n
		cl.tail = &n
	}else{
		n.next = cl.tail.next
		cl.tail.next = &n
	}
	cl.tail = &n
	cl.len++
	return
}
func (cl *CirSingLinkedList) AddAny(pos int, val int){
	n := CNode{value: val}
	if pos < 0 {
		return
	}
	if pos == 0 {
		cl.head = &n
		cl.len++
		return
	}
	if pos > cl.len {
		return
	}
	new := cl.GetPos(pos)
	n.next = new
	prevNode := cl.GetPos(pos - 1)
	prevNode.next = &n
	cl.len++
}
func (cl *CirSingLinkedList) PrintAll(){
	if(cl.len==0){
		fmt.Println("No nodes there")
		return
	}
	n := cl.head
	i:=0	
	for(i<cl.len){
		fmt.Println(n.value)
		n = n.next
		i++
	}
	
}
func (cl *CirSingLinkedList) GetPos(pos int) *CNode {
	ptr := cl.head
	if pos < 0 {
		return ptr
	}
	if pos > (cl.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}
func (cl *CirSingLinkedList) Size() int{
	return cl.len
}
func (cl *CirSingLinkedList) RemoveFirst(){
	if(cl.IsEmpty()){
		fmt.Println("No elements is there to remove....")
		return
	}
	fmt.Println("The removed elem is -", cl.head.value )
	cl.tail.next = cl.head.next
	cl.head = cl.head.next
	if(cl.IsEmpty()){
		cl.head = nil
		cl.tail = nil
	}
	cl.len--
}
func (cl *CirSingLinkedList) RemoveLast(){
	if(cl.IsEmpty()){
		fmt.Println("No elements is there to remove....")
		return
	}
	fmt.Println("The removed elem is -", cl.tail.value )
	n := cl.head
	i:=0	
	for(i<cl.len-1){
		n = n.next
		i++
	}
	cl.tail = n
	cl.tail.next = cl.head

	cl.len--
}
func (cl *CirSingLinkedList) RemoveAny(pos int){
	n := cl.head
	i := 1
	for(i<pos-1){
		n = n.next
		i++
	}
	fmt.Println("The removed elemnt is- ",n.next.value)
	n.next = n.next.next
	cl.len--
}

func main(){
	cll := CirSingLinkedList{}
	fmt.Println(cll.IsEmpty())
	cll.Addfirst(10)
	cll.Addfirst(9)
	cll.Addfirst(8)
	cll.Addfirst(7)
	cll.Addfirst(6)
	cll.Addfirst(5)
	cll.Addfirst(4)
	cll.Addfirst(3)
	cll.Addfirst(2)
	cll.Addfirst(1)
	fmt.Println()
	fmt.Println(cll.Size())
	cll.RemoveFirst()
	cll.RemoveLast()
	cll.RemoveAny(5)
	cll.PrintAll()
	fmt.Println(cll.Size())

}

