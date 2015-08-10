package main
import "fmt"

type Vertex struct {
	X,Y int
}

func main(){
	v:=Vertex{1,2}
	v2:=Vertex{X:3}
	v3:=Vertex{Y:4}
	p:=&Vertex{5,6}
	fmt.Println(v,v2,v3,p)
}
