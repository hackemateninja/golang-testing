// main.go
package main

import (
	"fmt"
	"helloworld/hello"
	"helloworld/integers"
	"helloworld/iteration"
)

func main() {
	fmt.Println(hello.Hello("Heman", "Spanish"))
	fmt.Println(integers.Add(2, 2))
	fmt.Println(iteration.Repeat("a"))
}
