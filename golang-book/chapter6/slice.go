package main

import "fmt"


func main(){
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
    	fmt.Println(slice1, slice2)

	slice11 := []int{1,2,3}
    	slice12 := make([]int, 2)
    	copy(slice12, slice11)
    	fmt.Println(slice11, slice12)

}
