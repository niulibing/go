package main

import (
	"fmt"
	"os"
)

type User struct {
}

//go toString method
func (self User) String() string {
	return "this is user"
}

func main() {

	sprintf := fmt.Sprintf("float %f", 3.1415926)

	fmt.Println(sprintf)
	fmt.Fprintln(os.Stdout, "a\n")

	fmt.Printf("%v", User{})
}
