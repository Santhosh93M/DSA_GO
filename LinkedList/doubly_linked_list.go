package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}

type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (d *doublyLinkedList) AddFirst(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *doublyLinkedList) AddLast(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}
func (d *doublyLinkedList) AddAny(pos int, val string){
	n := node{data: val}
	if d.len==0{
		d.head = &n
		d.tail = &n
		d.len++
		return
	}
	if pos<0{
		println("Position cant be negative..!")
		return
	}
	if pos>d.len-1{
		println("Position cant be greater than length of the list...!")
		return
	}
	ptr := d.GetPos(pos)
	ptr.prev = &n
	n.next = ptr
	prev := d.GetPos(pos-1)
	prev.next = &n
	n.prev = prev
	d.len++
	return
}
func (cl *doublyLinkedList) GetPos(pos int) *node {
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
func (d *doublyLinkedList) RemoveFirst(){
	if d.len==0{
		fmt.Println("List is empty..!")
		return
	}
	fmt.Println("The removed element is : ", d.head.data)
	if d.head.next!=nil{
		d.head = d.head.next
		d.head.prev = nil
	}else{
		d.head = nil
		d.tail = nil
	}
	d.len--
}
func (d *doublyLinkedList) RemoveLast(){
	if d.len==0{
		fmt.Println("List is empty..!")
		return
	}
	fmt.Println("The removed element is : ", d.tail.data)
	if d.tail.prev!=nil{
		d.tail = d.tail.prev
		d.tail.next = nil
	}else{
		d.head = nil
		d.tail = nil
	}
	d.len--
}
func (d *doublyLinkedList) RemoveAny(pos int){
	if d.len==0{
		fmt.Println("List is empty..!")
		return
	}
	if pos<0{
		println("Position cant be negative..!")
		return
	}
	if pos>d.len-1{
		println("Position cant be greater than length of the list...!")
		return
	}
	if d.len==1{
		d.head=nil
		d.tail = nil
		d.len--
		return
	}
	ptr := d.GetPos(pos-1)
	fmt.Println("the removed element is : ", ptr.next.data)
	next := d.GetPos(pos+1)
	ptr.next = ptr.next.next
	next.prev = ptr
	d.len--
}
func (d *doublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("%v  ", temp.data)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkedList) Reverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkedList) Size() int {
	return d.len
}
var i string
func init(){
	fmt.Println("Doubly linked list...!")
	i = "santhosh"
}

func main() {
	doublyList := CirSingLinkedList{}
	doublyList.AddLast(0)
	doublyList.AddLast(0)
	doublyList.RemoveLast()
	doublyList.PrintAll()

	
}