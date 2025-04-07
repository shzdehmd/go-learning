package main

import "fmt";

func main() {
	// Create two variables with whole numbers for learning math
	x := 10
	y := 7

	// Performing basic math operations
	sum := x + y;
	difference := x - y;
	product := x * y;
	quotient := x / y;
	remainder := x % y;

	// Printing the results of the math operations
	fmt.Println("Learning math with whole numbers:");
	fmt.Println("x =", x);
	fmt.Println("y =", y);
	fmt.Println("x + y =", sum);
	fmt.Println("x - y =", difference);
	fmt.Println("x * y =", product);
	fmt.Println("x / y =", quotient);
	// When dividing integers, the result is also an integer and Go removes the decimal part
	// For example, 10 / 7 = 1.42857... but Go will return 1
	// To get the decimal part, you need to convert one of the numbers to float by using float64()
	// for example: fmt.Println("x / y =", float64(x) / float64(y)); // 1.4285714285714286
	fmt.Println("x % y =", remainder);

	// Create two variables with decimal numbers for learning math
	fX := 10.5
	fY := 7.2

	// Performing basic math operations
	fSum := fX + fY;
	fDifference := fX - fY;
	fProduct := fX * fY;
	fQuotient := fX / fY;
	// fRemainder := fX % fY; // This will not work because fX and fY are float numbers
	// To get the remainder, you need to convert one of the numbers to int by using int()
	// for example: fmt.Println("fX / fY =", int(fX) % int(fY)); // 3
	// fRemainder := int(fX) % int(fY); // This will work because fX and fY are converted to int

	// Printing the results of the math operations
	fmt.Println("Learning math with decimal numbers:");
	fmt.Println("fX =", fX);
	fmt.Println("fY =", fY);
	fmt.Println("fX + fY =", fSum);
	fmt.Println("fX - fY =", fDifference);
	fmt.Println("fX * fY =", fProduct);
	fmt.Println("fX / fY =", fQuotient);
	// When dividing float numbers, the remainder is always 0
	// For example, 10.5 / 7.2 = 1.4583333333333333
	// To get the remainder, you need to convert one of the numbers to int by using int()
	// for example: fmt.Println("fX / fY =", int(fX) / int(fY)); // 1
}
