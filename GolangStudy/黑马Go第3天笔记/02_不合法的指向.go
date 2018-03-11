package main

import(
	"fmt"
)

func main() {
	var p *int
	p=nil
	fmt.Println("p=",p)

	//*p=666//err，因为p没有合法指向
	var a int
	p=&a
	*p=666
	fmt.Printf("a=%d",a)
}
