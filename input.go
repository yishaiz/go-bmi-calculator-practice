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
	weight = getUserInput(info.WeightPrompt)
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
