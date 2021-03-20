package main

import (
	"fmt"
)

func main() {

	var number string
	fmt.Println("Enter number:")
	for {

		fmt.Scanln(&number)
		if number == "exit" {
			break
		}

		switch number {
		case "1":
			fmt.Println("~  not")
		case "2":
			fmt.Println("V  or")
		case "3":
			fmt.Println("__")
			fmt.Println("  )  if...then")
			fmt.Println("--")
		case "4":
			fmt.Println("__  there is an")
			fmt.Println("-|")
			fmt.Println("--")
		case "5":
			fmt.Println("=  equals")
		case "6":
			fmt.Println("0  zero")
		case "7":
			fmt.Println("S  the succesor of")
		case "8":
			fmt.Println("(  punctuation mark")
		case "9":
			fmt.Println(")  punctuation mark")
		case "10":
			fmt.Println(",  punctuation mark")

		}

	}
}
