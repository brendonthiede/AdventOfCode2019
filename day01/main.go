package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func requiredFuel(mass float64) float64 {
	return math.Max(math.Floor(mass/3)-2, 0)
}

func main() {
	var total, totalWithFuelForFuel float64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		fuelForMass := requiredFuel(mass)
		total += fuelForMass
		totalWithFuelForFuel += fuelForMass
		additionalFuel := requiredFuel(fuelForMass)
		for additionalFuel > 0 {
			totalWithFuelForFuel += additionalFuel
			additionalFuel = requiredFuel(additionalFuel)
		}
	}

	fmt.Printf("Total fuel requirement (without fuel):   %.0f\n", total)
	fmt.Printf("Total fuel requirement (including fuel): %.0f\n", totalWithFuelForFuel)
}
