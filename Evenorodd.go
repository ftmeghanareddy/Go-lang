package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}
}
