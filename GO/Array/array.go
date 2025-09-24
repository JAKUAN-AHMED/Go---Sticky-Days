package main

import "fmt"


var arr=[3]int {1,2,3}
func main() {
	fmt.Println("Array")

	//first way
	// var arr[2] int
	
	// arr[0] = 1
	// arr[1] = 2
	
	// arr:=[2]int {1,2}
	fmt.Println("Array is: ", arr)

	//second way
	var fruitList [4]string 
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "mango"
	fruitList[3] = "grapes"
	fmt.Println("Fruits list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))
}
