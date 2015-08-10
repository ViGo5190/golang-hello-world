package main

import "fmt"

func main(){
	el := make(map[string]string)
	el["A"] = "a"
	el["B"] = "b"

	fmt.Println(el)

	name, ok := el["Z"]

	fmt.Println(name, ok)

	if name, ok := el["A"]; ok {
		fmt.Println(name, ok)
	}
}
