package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//fmt.Println("hello")

	// output information
	fmt.Println("BMI Calculator")
	fmt.Println("-------------------")

	// prompt for user input (weight + height)
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')

	// save the user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	fmt.Print(weightInput, heightInput)
	fmt.Print

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// calculate the BMI (weight / (height * height))
	bmi := weight / (height * height)

	// output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}
