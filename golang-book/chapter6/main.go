package main

import "fmt"


func main(){
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var marks [5]float64
	marks[0] = 81
	marks[1] = 82
	marks[2] = 82
	marks[3] = 83
	marks[4] = 84
	
	var total float64 = 0
	
	for i:=1;i<len(marks);i++ {
		total += marks[i]
	}

	fmt.Println(total / float64(len(marks)))	

	marks2 := [5]float64{92,93,94,95,96}
	var total2 float64 = 0

	for _,value := range marks2 {
		total2 += value
	}

	fmt.Println(total2 / float64(len(marks2)) )
}
