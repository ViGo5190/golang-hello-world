package main

import (
	"fmt"
)

func sq(x float64) float64{
//	for i:=1;i<x;i++{
//		z:=i
//		fmt.Println(i)
//	}
	z := 1.0
	z = z-((z*z - x)/(2*z))
	return z
}

func main(){
	fmt.Println(sq(9));
}
