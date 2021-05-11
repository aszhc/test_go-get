package main

import (
	"./mypkg"
	"fmt"

)

func main() {
	mypkg.CustomPkgFunc()
	fmt.Println("hello world!")
}

func SayHello()  {
	fmt.Println("Hello mother fucker...")
}