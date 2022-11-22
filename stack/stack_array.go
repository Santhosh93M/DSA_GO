package main

import "fmt"

type StackInt struct {
	s []int
}

func (s *StackInt) Len() int {
	l := len(s.s)
	return l
}

func (s *StackInt) IsEmpty() bool {
	l := len(s.s)
	return l == 0
}

func (s *StackInt) Print() {
	l := len(s.s)
	for i := 0; i < l; i++ {
		fmt.Println("Value : ",s.s[i])
	}
}
func (s *StackInt) Push(val int){
	s.s = append(s.s, val)
}

func (s *StackInt) Pop() int{
	l := len(s.s)
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *StackInt) Top() int{
	l := len(s.s)
	return s.s[l-1]
}

func main(){
	 sa:= StackInt{}
	 sa.Push(12)
	 sa.Push(23)
	 sa.Push(10)
	 sa.Print()
	 fmt.Println("Top of the array is: ", sa.Top())
	 fmt.Println(sa.Pop())
	 sa.Print()
}