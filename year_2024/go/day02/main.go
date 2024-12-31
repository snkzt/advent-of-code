package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isSafeReport(report []int) bool {

	for i := 0; i < len(report)-2; i++ {
		// The levels are either all increasing or all decreasing.
		if !(report[i] < report[i+1] && report[i+1] < report[i+2]) &&
			!(report[i] > report[i+1] && report[i+1] > report[i+2]) {
			return false
		}

		// Any two adjacent levels differ by at least one and at most three.
		if (absInt(report[i]-report[i+1]) < 1 || absInt(report[i]-report[i+1]) > 3) ||
			(absInt(report[i+1]-report[i+2]) < 1 || absInt(report[i+1]-report[i+2]) > 3) {
			return false
		}
	}

	return true
}

func parseLine(input []byte) ([]int, error) {
	fields := bytes.Fields(input)
	var inputData []int

	if len(fields) > 0 {
		for _, field := range fields {
			val, err := strconv.Atoi(string(field))
			if err != nil {
				return nil, fmt.Errorf("invalid number: %v", err)
			}

			inputData = append(inputData, val)
		}
	}

	return inputData, nil
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func countTotalSafeReports(input []byte) (int, error) {
	totalSafeReports := 0
	lines := bytes.Split(input, []byte("\n"))

	// lines -> [][]byte, row -> []byte
	for _, row := range lines {
		if len(row) == 0 {
			continue
		}

		inputRowData, err := parseLine(row)
		if err != nil {
			return 0, err
		}

		safeReport := isSafeReport(inputRowData)
		if safeReport {
			totalSafeReports++
		}
	}

	return totalSafeReports, nil
}

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening input data failed: %v", err)
	}

	return data, nil
}

func main() {
	filePath := "../../shared/inputs/day02/input.txt"

	inputData, err := readFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	totalSafeReports, err := countTotalSafeReports(inputData)
	if err != nil {
		log.Fatalf("Error counting total safe reports: %v", err)
	}

	log.Println("Total count of safe reports:", totalSafeReports)
}
