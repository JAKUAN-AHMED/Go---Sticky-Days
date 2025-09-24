package main


/**
*function
#first order - a normal function is first order
#high order - a function which takes another function as a parameter
#anonymous - no name
#IIFE - immediately invoked function
#init - a function which will be executed first
#first class function - first class function is 
*
*/
import (
	"fmt"
	"example.go/mathlib"
)

//high order function
/*
1st:paramter=>function
2nd:return function
3rd:both
*/

func processOperation(f func(int, int) int, a, b int) int {
	return f(a, b)
}
func main() {
	// fmt.Println("Hello, World!")
	// //modular printed
	// sum := mathlib.Add(5, 9)
	// fmt.Println(sum)

	//high order function
	fmt.Println(processOperation(mathlib.Add, 5, 9))
	

	//anonymous function and IIFE || immediately invoked function
	// func (a int,b int){
	// 	fmt.Println("I am anonymous function")
	// 	fmt.Println(a+b)
	// }(5,7)

}


//init function
func init(){
	fmt.Println("I will be printed first")
}