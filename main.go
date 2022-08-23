package main

import "fmt"

func main() {
	for a := 1; a <= 10; a++ {
		fmt.Print(a)
		if a%2 == 0 {
			fmt.Println(" genap")
		} else {
			fmt.Println(" ganjil")
		}
	}
}
