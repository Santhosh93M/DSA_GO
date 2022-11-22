package main

import "fmt";import "errors"
// Node represents a node of linked list
type Node struct {
	value int
	next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len int
}
func (l * LinkedList) Reverse() {
	if l.len==0{
		fmt.Println("no elements in linked list")
		return
	}
	curr := l.head
	var Prev *Node
	var Next *Node
	Prev = nil
	Next = nil
	for curr!=nil{
		Next = curr.next
		curr.next = Prev
		Prev = curr
		curr = Next
	}
	l.head = curr
}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", ptr.value)
		ptr = ptr.next
	}
	println()
}
func (l *LinkedList) AddFirst(val int){
	n := Node{value: val}
	if l.len==0{
		l.head=&n
		l.len++
		return
	}
	ptr := l.head
	n.next = ptr
	l.head = &n
	l.len++
}
func (l *LinkedList) Insert(val int) {
	n := Node{value: val}
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}
func (l *LinkedList) Sort(){
	if l.len==0{
		fmt.Println("There is No items in linked list...!")
		return
	}
	ptr := l.head
	for ptr.next!=nil{
		temp := ptr.next
		for temp.next!=nil{
			if ptr.value>temp.value{
				t := temp.value
				temp.value = ptr.value
				ptr.value = t
			}
			temp = temp.next
		}
		if ptr.value>temp.value{
			t1 := temp.value
			temp.value = ptr.value
			ptr.value = t1
		}
		ptr = ptr.next
	}
}
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}
func (l *LinkedList) InsertAt(pos int, val int) {
	// create a new node
	newNode := Node{}
	newNode.value = val
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}
func (l *LinkedList) DeleteAt(pos int) error {
	// validate the position
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos+1)
	l.len--
	return nil
}
func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}
func (l *LinkedList) size()int{
	return l.len
}
func main(){
	l := LinkedList{}
	var choice, pos, val int
	var exit string
loop:
	fmt.Println("Choose the action below..!")
	fmt.Println("1.Insert 2.Insert_At 3.Delete 4.Display ")
	fmt.Scanln(&choice)
	switch choice{
	case 1:
		fmt.Println("Enter the value : ")
		fmt.Scanln(&val)
		l.Insert(val)
	case 2:
		fmt.Println("Enter the value : ")
		fmt.Scanln(&val)
		fmt.Println("Enter the position you want to insert : ")
		fmt.Scan(&pos)
		l.InsertAt(pos,val)
	case 3:
		fmt.Println("Enter the position you want to Delete : ")
		fmt.Scan(&pos)
		l.DeleteAt(pos)
	case 4:
		l.Print()
	default:
		fmt.Println("Entered Wrong choice..!")
	}
	fmt.Println("Do you want to continue(y/n)..")
	fmt.Scanln(&exit)
	if exit=="y"{
		goto loop
	}else{
		fmt.Println("bye...8")
	}
}