package app

import "fmt"


//scoped function
func test() {
	fmt.Print("hi im the main test still")
}

//package methods 
func Test() {
	test()
	fmt.Print("hi im the BIG Test")
}