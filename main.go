package main

import "fmt"

func main() {

	// 1 way
	// fmt.Print(mul_div(100,5));

	//2 way
	// mul, div := mul_div(100,5);
	// fmt.Println("Multiply of two number is", mul, "Division of tweo number is", div);

	//3 way
	_, div := mul_div(100, 5) // if not using the "mul" variable anywhere in the program use blank identifier, to let 'go' compiler to ignore it. Unused declared variable gives error
	fmt.Println("Division of two number is", div)

}

func mul_div(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2

}
