// main.go
package main

import (
	"fmt"
	"helloworld/di"
	"helloworld/hello"
	"helloworld/integers"
	"helloworld/iteration"
	pointerserrors "helloworld/pointers_errors"
	structsmethodsinterfaces "helloworld/structs_methods_interfaces"
	"os"
)

func main() {
	fmt.Println(hello.Hello("Heman", "Spanish"))
	fmt.Println(integers.Add(2, 2))
	fmt.Println(iteration.Repeat("a"))

	rectangle := structsmethodsinterfaces.Rectangle{Height: 20, Width: 30}

	circle := structsmethodsinterfaces.Circle{Radius: 10}

	fmt.Println(rectangle.Area())

	fmt.Println(circle.Area())

	wallet := pointerserrors.Wallet{}
	wallet.Deposit(50)

	fmt.Println(wallet.Balance())

	di.Greet(os.Stdout, "Elodie")

	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreetHandler)))
}
