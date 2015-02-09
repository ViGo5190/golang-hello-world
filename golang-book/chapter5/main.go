package main

import "fmt"

func main(){
	i := 1
	for i <= 10 {
		fmt.Println(i)
	i += 1
	}
	
	for j:=1; j<=10;j+=1 {
		if j%2 == 0{
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}
	}
}
