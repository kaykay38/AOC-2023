/*
Advent of Code 2023

Day 1: Trebuchet?! 
Part One

Author: Mia Hunt
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read in a file line by line.
// For each line, add each digit occurence to a slice
// and add the combination of the first and last digit to sum.
func getSum(file *os.File) int {
    // create scanner to read file line by line
    scanner := bufio.NewScanner(file)

    var sum int
    for scanner.Scan() {
        line := scanner.Text()
        digits := []int{}

        for _, val := range line {
            if (val >= 48 && val <= 57) {
                digits = append(digits, int(val)-48)
            }
        }

        sum += digits[0]*10 + digits[len(digits)-1]
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
    
    fmt.Println("Sum:", sum)
    return sum
}

func main()  {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()
    getSum(file)
}
