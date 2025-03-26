package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Check if the file argument is provided
func main() {
	if len(os.Args) > 2 {
		log.Fatal("The program accepts only the path file as argument.")
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the file as an argument.")
	}

	fileName := os.Args[1]

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a slice to store y-values
	var yValues []float64

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert each line to a float64 and append to yValues slice
		y, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		yValues = append(yValues, y)
	}

	// Check for errors during reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Create the x-values slice
	var xValues []float64
	for i := 0; i < len(yValues); i++ {
		xValues = append(xValues, float64(i))
	}

	// Calculate the means of x and y
	var sumX, sumY float64
	for i := 0; i < len(xValues); i++ {
		sumX += xValues[i]
		sumY += yValues[i]
	}

	meanX := sumX / float64(len(xValues))
	meanY := sumY / float64(len(yValues))

	// Calculate the slop (m) and the intercept (b) for the linear regression line
	var num, den float64
	for i := 0; i < len(xValues); i++ {
		num += (xValues[i] - meanX) * (yValues[i] - meanY)
		den += (xValues[i] - meanX) * (xValues[i] - meanX)
	}

	m := num / den
	b := meanY - m*meanX

	// Calculate the Pearson Correlation Coefficient
	var sumXY, sumX2, sumY2 float64
	for i := 0; i < len(xValues); i++ {
		sumXY += xValues[i] * yValues[i]
		sumX2 += xValues[i] * xValues[i]
		sumY2 += yValues[i] * yValues[i]
	}
	pearson := (float64(len(xValues))*sumXY - sumX*sumY) / math.Sqrt((float64(len(xValues))*sumX2-sumX*sumX)*(float64(len(yValues))*sumY2-sumY*sumY))

	// // Print the linear regression line equation (y = mx + b)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)

	// Print the Pearson Correlation Coefficient
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
}
