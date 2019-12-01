package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	runTestCases()

	modules := loadEntriesFromFile("input.txt")
	fmt.Printf("Calculating sum of fuel requirements for %d modules.\n", len(modules))

	total, totalPrime := 0, 0
	for i := 0; i < len(modules); i++ {
		total += calculateFuel(modules[i])
		totalPrime += calculateFuelPrime(modules[i])
	}

	fmt.Printf("Total fuel: %d\n", total)
	fmt.Printf("Total fuel prime: %d\n", totalPrime)
}

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0) - 2)
}

func calculateFuelPrime(mass int) int {
	fuel, result := mass, 0
	for {
		fuel = calculateFuel(fuel)
		if fuel > 0 {
			result += fuel
		} else {
			break
		}
	}

	return result
}

func loadEntriesFromFile(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	massEntries := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		massEntries = append(massEntries, m)
	}

	return massEntries
}

func runTestCases() {
	inputs := []int{12, 14, 1969, 100756}
	expectedOutputs := []int{2, 2, 654, 33583}
	expectedOutputsPrime := []int{2, 2, 966, 50346}
	for i := 0; i < len(inputs); i++ {
		output := calculateFuel(inputs[i])
		outputPrime := calculateFuelPrime(inputs[i])
		result := output == expectedOutputs[i] && outputPrime == expectedOutputsPrime[i]

		fmt.Printf("in: %d, out: %d, outPrime: %d, result = %v\n", inputs[i], output, outputPrime, result)
	}
}
