package main

import "fmt"
//import "os"//if not used it show error

func main(){
	var a="hii";// variable declaration  same as js
	fmt.Println(" this is a addition program \n  \" this is quote \"  "+a+" \nstatement is end \n")

	var number=2
	var b=5
	var add=number+b

	fmt.Println(b)
	fmt.Println("+")
	fmt.Println(number);
	//fmt.Println("number ="+number)	
	// error =cannot use "number =" (type untyped string) as type int

	fmt.Println("_______\t\tAddition is")
	fmt.Print(add)
	
}