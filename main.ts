/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-12-05
 * @fileoverview This program asks the user for the number of bolts, nuts, and washers in their purchase
 * and then calculates and prints out the total cost of the order.
 */

// constants
const MIN_WEIGHT: number = 77;
const MAX_WEIGHT: number = 105;

// variable
let weightAsString: string;
let weightAsNumber: number;

// input weight
weightAsString = prompt("How much do you weigh in kg? ") || "0";
weightAsNumber = parseFloat(weightAsString);

// check if weight is permitted
if (weightAsNumber >= 77 && weightAsNumber <= 105) {
  console.log("You may enter the contest");
}

console.log("\nDone.");
