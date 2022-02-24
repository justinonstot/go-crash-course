package main

import "myapp/packageone"

var myVar string = "I am myVar"

func main() {

	var blockVar string = "I am blockVar"

	secondVar := "some value"
	packageone.PrintMe(myVar, blockVar)

}
