package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
const mainTitle = "BMI Calculator"
const separatorLine = "-------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightPrompt = "Please enter your height (kg): "

//const reader = bufio.NewReader(os.Stdin)

func main() {
	//fmt.Println("hello")

	// output information
	fmt.Println(mainTitle)
	fmt.Println(separatorLine)

	// prompt for user input (weight + height)
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(heightPrompt)
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
