package main

import "fmt"


//pointer with recursion

func printArray(arr *[]int, n int) {

	if n==0{
		return
	}

	printArray(arr, n-1)
	fmt.Println((*arr)[n-1])
	
}
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	printArray(&arr, n)

}
