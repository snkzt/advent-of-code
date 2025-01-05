package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func parseLine(line []byte) (int, int, error) {
	fields := bytes.Fields(line)
	var num1, num2 int
	var err error

	if len(fields) > 0 {
		num1, err = strconv.Atoi(string(fields[0]))
		if err != nil {
			return 0, 0, fmt.Errorf("invalid number in column 1: %v", err)
		}
	}

	if len(fields) > 1 {
		num2, err = strconv.Atoi(string(fields[1]))
		if err != nil {
			return 0, 0, fmt.Errorf("invalid number in column 2: %v", err)
		}
	}

	return num1, num2, nil
}

func parseInputData(inputData []byte) ([]int, []int, error) {
	var column1, column2 []int

	// Split a byte slice "data" into a slice of byte slices ([][]byte) based on delimiter"\n".
	// This splits the original data into a collection of lines.
	lines := bytes.Split(inputData, []byte("\n"))

	// Retrieve int columns
	for _, line := range lines {
		if len(line) == 0 {
			continue // Skip empty lines.
		}
		num1, num2, err := parseLine(line)
		if err != nil {
			return nil, nil, err
		}
		column1 = append(column1, num1)
		column2 = append(column2, num2)
	}

	// Sort each column in-place
	sortColumns(column1, column2)

	return column1, column2, nil
}

func sortColumns(column1, column2 []int) {
	sort.Ints(column1)
	sort.Ints(column2)

}

func calculateSimilarityScore(column1, column2 []int) int {
	countMap := make(map[int]int)

	for _, num := range column2 {
		countMap[num]++
	}

	similarityScore := 0
	for _, num := range column1 {
		similarityScore += countMap[num] * num
	}

	return similarityScore
}

func calculateTotalDistance(column1, column2 []int) (int, error) {
	// Assuming the location IDs are only integer here.
	totalDistance := 0
	// Calculate each row difference
	for i := 0; i < len(column1); i++ {
		totalDistance += absInt(column1[i] - column2[i])
	}

	return totalDistance, nil
}

func readFile(filePath string) ([]byte, error) {
	// Read the entire file into memory as the input is only 96B.
	// If it gets bigger like GB, then better read line by line with bufio.
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening input data failed: %v", err)
	}

	return data, nil
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func main() {
	filePath := "../../shared/inputs/day01/input.txt"

	// Read the file content
	inputData, err := readFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Parse input data into two []int
	column1, column2, err := parseInputData(inputData)
	if err != nil {
		log.Fatalf("Error parsing input data: %v", err)
	}

	totalDistance, err := calculateTotalDistance(column1, column2)
	if err != nil {
		log.Fatalf("Error calculating total distance: %v", err)
	}
	similarityScore := calculateSimilarityScore(column1, column2)
	if err != nil {
		log.Fatalf("Error calculating similarity score: %v", err)
	}

	fmt.Println("Total Distance:", totalDistance)
	fmt.Println("Similarity Score:", similarityScore)
}
