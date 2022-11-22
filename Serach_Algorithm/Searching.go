package main

import "fmt"

func LinearSearch(v []int, val int){
	for index,i:=range v{
		if i==val{
			fmt.Printf("%d found at %dth position.",val,index)
			fmt.Println()
			return
		}
	}
	fmt.Println("Not found..")
}
func LinearSorting(v []int) []int{
	for i:=0;i<len(v);i++{
		for j:=i+1;j<len(v);j++{
			if v[i]>v[j]{
				v[i],v[j] = v[j],v[i]
			}
		}
	}
	return v
}
func BinarySearch(v []int, key int){
	sorted_arr := LinearSorting(v)
	i:=0
	j:=len(v)-1
	for i<=j{
		mid := (i+j)/2
		if sorted_arr[mid]==key{
			fmt.Printf("%d found at %dth position.",key,mid)
			return
		}else if (key < sorted_arr[mid]){
			j = mid-1
		}else if (key > sorted_arr[mid]){
			i = mid+1
		}
	}
	fmt.Println(("Not found..!"))
}

func main(){
	arr := []int{12,98,76,54,90,123}
	BinarySearch(arr, 90)
}