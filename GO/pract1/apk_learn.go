package main

import "fmt"


func func1(){
	fmt.Println("func1 called")
}

func add(a int,b int){
	print(a+b)
	fmt.Println("\n")
	fmt.Println("a is ",a)

}

// any function is must be outside of main function
func main(){
	fmt.Println("Started")
	fmt.Println(" Program ")
	print("Hii I am ")
	print(" Deepak ")
	print("Ugale")
	fmt.Println("  Ended ")
	fmt.Println(" execution ")
	/*
C:\Users\Deepak\Desktop\GO\GO\pract1>go run apk_learn.go
Started
 Program
Hii I am  Deepak Ugale Ended
 execution
*/
print("\n\n\n\n\n")
	for i := 0; i < 10; i++ {
		print(i)
	}

	fmt.Println("\n\n\n")

	add(5,6)
	//func1()
	var x string="Hiiy"
	fmt.Println(" and data type is %T" , x)

	print("hoo %T\n",x)
}