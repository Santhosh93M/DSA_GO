package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
    
    for _,i:=range sandwiches{
        count := 0
        for i!=students[count]{
            students = append(students, students[count])
            students = students[1:]
        }
        students = students[1:]
        
    }
    return len(students)
}

func main(){
	st := []int{1,1,1,0,0,1}
	sa := []int {1,0,0,0,1,1}
	fmt.Println(countStudents(st, sa))
}