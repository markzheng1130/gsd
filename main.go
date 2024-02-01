package main

import (
	"fmt"
	"gsd/hello"

	"github.com/ozgio/strutil"
)

func main() {
	s := hello.Hello("mark")
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", strutil.Reverse(s))
}
