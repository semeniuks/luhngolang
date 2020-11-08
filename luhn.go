package main

import (
	"fmt"
	"strconv"
)

func checkIfLuhnNumber(str string) {
	stringLenght := len(str)
	sum := 0
	for i := 0; i <= stringLenght-1; i++ {
		sNumber := string(str[stringLenght-1-i])
		number, _ := strconv.Atoi(sNumber)
		if (i+1)%2 == 0 {
			number = number * 2
		}
		if number >= 10 {
			strNumber := strconv.Itoa(number)
			numbeOne, _ := strconv.Atoi(string(strNumber[0]))
			numberTwo, _ := strconv.Atoi(string(strNumber[1]))
			number = numbeOne + numberTwo
		}
		sum += number
	}
	if sum%10 == 0 {
		fmt.Println("Valid Luhn number")
	} else {
		fmt.Println("Not Valid Luhn number")
	}

}
func main() {
	checkIfLuhnNumber("4978")
}
