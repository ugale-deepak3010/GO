package main

import "fmt"

func main(){
	var i string// data type is optional
	i=" Deepak "
	fmt.Println("hii"+i)

	var a=21

	if(a==2){
		fmt.Println("2")
	}else{
		fmt.Println(" false")
	}

	switch a{
		case 21:
		fmt.Println("a is switch 21")
	default:
		fmt.Println("default")
	}
}