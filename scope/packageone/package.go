package packageone

import "fmt"

var PackageVar string = "I am PackageVar"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
