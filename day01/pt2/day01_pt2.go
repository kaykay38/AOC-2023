/*
Advent of Code 2023

Day 1: Trebuchet?!
Part Two

Author: Mia Hunt
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read in a file line by line.
// For each line, add each digit/spelled out number occurence to a slice
// and add the combination of the first and last digit to sum.
func getSum(file *os.File) int {
	// create scanner to read file line by line
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		digits := []int{}
		n := len(line)
		for i, char := range line {
			if char >= 48 && char <= 57 {
				digits = append(digits, int(char)-48)
			} else {
				// Check if the current number forms a spelled-out number
				switch char {
				case 'o':
					if i+3 <= n && line[i:i+3] == "one" {
						digits = append(digits, 1)
						i += 2
					}
				case 't':
					if i+3 <= n && line[i:i+3] == "two" {
						digits = append(digits, 2)
						i += 2
					}
					if i+5 <= n && line[i:i+5] == "three" {
						digits = append(digits, 3)
						i += 4
					}
				case 'f':
					if i+4 <= n {
						switch line[i : i+4] {
						case "four":
							digits = append(digits, 4)
							i += 3
						case "five":
							digits = append(digits, 5)
							i += 3
						}
					}
				case 's':
					if i+3 <= n && line[i:i+3] == "six" {
						digits = append(digits, 6)
						i += 2
					}
					if i+5 <= n && line[i:i+5] == "seven" {
						digits = append(digits, 7)
						i += 4
					}
				case 'e':
					if i+5 <= n && line[i:i+5] == "eight" {
						digits = append(digits, 8)
						i += 4
					}
				case 'n':
					if i+4 <= n && line[i:i+4] == "nine" {
						digits = append(digits, 9)
						i += 3
					}
				}
			}
		}

		if len(digits) > 0 {
			sum += digits[0]*10 + digits[len(digits)-1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Sum:", sum)
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	getSum(file)
}
