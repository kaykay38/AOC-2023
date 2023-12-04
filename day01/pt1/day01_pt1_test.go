/*
Advent of Code 2023

Day 1: Trebuchet?! 
Part One Test

Author: Mia Hunt
*/

package main

import (
    "testing"
    "os"
)

func TestGetSum(t *testing.T) {
    file, err := os.Open("test_input.txt")
    if err != nil {
        t.Fatal("Error opening file:", err)
        return
    }
    defer file.Close()
    
    // Computed sum for the sample content
    got := getSum(file)

    // Expected sum for the sample content
    expected := 142

	// Compare the expected sum with the computed sum
    if got != expected {
        t.Errorf("getSum() = %d; want %d", got, expected)
    }
}
