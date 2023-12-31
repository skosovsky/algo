package main

import "fmt"

func main() {
	var numN, numK int
	fmt.Scan(&numN, &numK) // standard input

	fmt.Print(combination(numN, numK)) // standard output
}

func combination(numN, numK int) int {
	return factorial(numN) / (factorial(numK) * factorial(numN-numK))
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

// by Python
// def combination(numN, numK):
//	return factorial(numN) // (factorial(numK) * factorial(numN-numK))
//
// def factorial(num):
//	if num <= 1:
//		return 1
//	return num * factorial(num-1)
//
// numN, numK = map(int, input().split())
// print(combination(numN, numK))
