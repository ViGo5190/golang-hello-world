package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	j:=0

	if (sum>10){
		for z:=1;z<100000000;z++ {
			j += z
		}
	}


	c <- sum
}

func main(){

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	fmt.Println("end");
}
