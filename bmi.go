package main

import (
	"fmt"
	"go-bmi-calculator-practice/info"
	"strconv"
	"strings"
)

func main() {

	// output information
	fmt.Println(info.MainTitle)
	fmt.Println(info.SeparatorLine)

	// prompt for user input (weight + height)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	// save the user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	fmt.Print(weightInput, heightInput)
	fmt.Println()

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
