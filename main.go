/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-12-05
 * Fileoverview: This program asks the user for the number of bolts, nuts, and washers in their purchase
 * and then calculates and prints out the total cost of the order.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// constants
	const MIN_WEIGHT float64 = 77
	const MAX_WEIGHT float64 = 105

	// variables
	var weightAsString string
	var weightAsNumber float64

	reader := bufio.NewReader(os.Stdin)

	// input weight
	fmt.Print("How much do you weigh in kg? ")
	weightAsString, _ = reader.ReadString('\n')
	weightAsString = strings.TrimSpace(weightAsString)
	weightAsNumber, _ = strconv.ParseFloat(weightAsString, 64)

	// check if weight is permissible
	if (weightAsNumber >= MIN_WEIGHT && weightAsNumber <= MAX_WEIGHT) {
		fmt.Println("You may enter the contest.")
	}

	fmt.Println("\nDone.")
}
