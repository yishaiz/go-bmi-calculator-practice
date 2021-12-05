package main

import (
	"fmt"
	"go-bmi-calculator-practice/info"
)

func main() {
	info.PrintWelcome()

	weight, height := getUserMetrics()

	// calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
