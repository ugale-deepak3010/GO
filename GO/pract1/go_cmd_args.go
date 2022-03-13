package main

import (
	"fmt"
	"os"
	)

func main(){
	
	o:="I am o without defined"//short variable declaration
	fmt.Println(o)
	//var o="oo"// o redeclared in this block error
	//var i="ii"//i declared but not used
	fmt.Println("\n\n\n")


	for i := 0; i < len(os.Args); i++ {

		fmt.Println(os.Args[i])
	}

	//var i="ii"//i declared but not used

}