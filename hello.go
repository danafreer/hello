// hello.go
package main

import (
	"fmt"
	"github.com/danafreer/stringutil"
)

func main() {
	fmt.Println("Hi, Dana")
	fmt.Printf(stringutil.Reverse("hello, world xxx\n"))
}
