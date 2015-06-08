package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Print("Os is...")

	switch os:=runtime.GOOS;os {
		case "darwin":
			fmt.Println("OSX")
		case "linux":
			fmt.Println("linux")
		default :
			fmt.Printf("%s.", os)
	}

}
