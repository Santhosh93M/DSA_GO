package main

import "fmt"

func LinearSearch(v []int, val int) {
	for index, i := range v {
		if i == val {
			fmt.Printf("%d found at %dth position.", val, index)
			fmt.Println()
			return
		}
	}
	fmt.Println("Not found..")
}

func main() {
	arr := []int{12, 98, 76, 54, 90, 123}
	LinearSearch(arr, 90)
}
