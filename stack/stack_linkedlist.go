package main

import "fmt"

type SLNode struct {
	value int
	next  *SLNode
}
type StackIn struct {
	head *SLNode
	len  int
}

func (s *StackIn) Len() int {
	return s.len
}
func (s *StackIn) IsEmpty() bool {
	return s.len == 0
}
func (s *StackIn) Push(val int) {
	n := SLNode{value: val}
	if s.IsEmpty() {
		s.head = &n
		s.len++
		return
	}
	ptr := s.head
	for i := 0; i < s.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			s.len++
			return
		}
		ptr = ptr.next
	}
}
func (s *StackIn) Print() {
	n := s.head
	for i := 0; i < s.len; i++ {
		fmt.Printf("%d ",n.value)
		n = n.next
	}
	fmt.Println()
}
func (s *StackIn) Top(){
	n := s.head
	for i:=0;i<s.len-1;i++{
		n=n.next
	}
	fmt.Println("The top element is : ", n.value)
}
func (s *StackIn) Pop(){
	n := s.head
	for i:=1;i<s.len-1;i++{
		n = n.next
	}
	fmt.Println("The popped element is : ", n.next.value)
	n.next = nil
	s.len--

}

func main(){
	sa := StackIn{}
	sa.Push(12)
	sa.Push(34)
	sa.Push(45)
	sa.Push(90)
	sa.Print()
	sa.Top()
	sa.Pop()
	sa.Print()
}