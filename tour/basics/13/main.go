package main

import (
	"fmt"
	"math"
)

func main(){
//	var i int = 42
//	var f float64 = float64(i)
//	var u uint = uint(f)

	var x,y int = 3,4
	var ff float64 = math.Sqrt(float64(x*x+y*y))
	var z int = int(ff)
	fmt.Println(ff, z)
}
