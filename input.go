package main

import (
	"bufio"
	"fmt"
	"go-bmi-calculator-practice/info"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	//fmt.Print(info.WeightPrompt)
	//weightInput, _ := reader.ReadString('\n')
	//weightInput = strings.Replace(weightInput, "\n", "", -1)
	//weight, _ = strconv.ParseFloat(weightInput, 64)

	weight = getUserInput(info.WeightPrompt)

	//fmt.Print(info.HeightPrompt)
	//heightInput, _ := reader.ReadString('\n')
	//heightInput = strings.Replace(heightInput, "\n", "", -1)
	//height, _ = strconv.ParseFloat(heightInput, 64)
	height = getUserInput(info.HeightPrompt)

	return weight, height
}

func getUserInput(propmtText string) (value float64) {
	fmt.Print(propmtText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	value, _ = strconv.ParseFloat(userInput, 64)
	return value
}
