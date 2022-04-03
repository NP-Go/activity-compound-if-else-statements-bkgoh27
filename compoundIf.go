package main

import "fmt"

func isLeapYear(year int) bool {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%v is a leap year", year)
		return true
	} else {
		fmt.Printf("%v is not a leap year", year)
		return false
	}
}

func main() {
	var year int
	fmt.Println("Give me the year")
	fmt.Scanln(&year)

	isLeapYear(year)
}
