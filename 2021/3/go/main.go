package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	fmt.Println("1. What is the power consumption of the submarine?")
	fmt.Println(getPowerConsumption(lines))

	fmt.Println("2. What is the power consumption of the submarine?")

	oxygenGeneratorRating := getOxygenGeneratorRating(lines)
	co2ScrubberRating := getCO2ScrubberRating(lines)
	fmt.Println(oxygenGeneratorRating * co2ScrubberRating)
}

func getPowerConsumption(lines []string) int {
	var offBitCount []int
	var onBitCount []int

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if i == 0 {
			offBitCount = make([]int, len(line))
			onBitCount = make([]int, len(line))
		}

		chars := []rune(line)

		for j := 0; j < len(chars); j++ {
			switch chars[j] {
			case '0':
				offBitCount[j]++
			case '1':
				onBitCount[j]++
			default:
				log.Fatal("There is not a bit value")
				log.Fatal(chars[j])
			}
		}
	}

	gammaRateString := ""
	epsilonRateString := ""

	for i := 0; i < len(offBitCount); i++ {
		if offBitCount[i] > onBitCount[i] {
			gammaRateString += "0"
			epsilonRateString += "1"
		} else if onBitCount[i] > offBitCount[i] {
			gammaRateString += "1"
			epsilonRateString += "0"
		} else {
			log.Fatal("There is no most common bit")
		}
	}

	gammaRate, err := strconv.ParseInt(gammaRateString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilonRate, err := strconv.ParseInt(epsilonRateString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gamma rate: %s is %d.\n", gammaRateString, gammaRate)
	fmt.Printf("Epsilon rate: %s is %d.\n", epsilonRateString, epsilonRate)

	return int(gammaRate) * int(epsilonRate)
}

func getOxygenGeneratorRating(lines []string) int {
	bitCount := len(lines[0])

	for i := 0; i < bitCount && len(lines) > 1; i++ {
		lines = getMostCommonBitValueLines(lines, i)
	}

	if len(lines) != 1 {
		log.Fatal("It's not possible to calculate oxygen generator rating.")
		return 0
	}

	oxygenGeneratorRating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Oxygen generator rate: %s is %d.\n", lines[0], oxygenGeneratorRating)

	return int(oxygenGeneratorRating)
}

func getCO2ScrubberRating(lines []string) int {
	bitCount := len(lines[0])

	for i := 0; i < bitCount && len(lines) > 1; i++ {
		lines = getLeastCommonBitValueLines(lines, i)
	}

	if len(lines) != 1 {
		log.Fatal("It's not possible to calculate CO2 scrubber rating.")
		return 0
	}

	co2ScrubberRating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CO2 scrubber rate: %s is %d.\n", lines[0], co2ScrubberRating)

	return int(co2ScrubberRating)
}

func getMostCommonBitValueLines(lines []string, bitIndex int) []string {
	offBitLines := make([]string, 0)
	onBitLines := make([]string, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			break
		}

		chars := []rune(line)

		switch chars[bitIndex] {
		case '0':
			offBitLines = append(offBitLines, line)
		case '1':
			onBitLines = append(onBitLines, line)
		default:
			log.Fatal("There is not a bit value")
			log.Fatal(chars[bitIndex])
		}
	}

	if len(onBitLines) >= len(offBitLines) {
		return onBitLines
	} else {
		return offBitLines
	}
}

func getLeastCommonBitValueLines(lines []string, bitIndex int) []string {
	offBitLines := make([]string, 0)
	onBitLines := make([]string, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if len(line) == 0 {
			break
		}

		chars := []rune(line)

		switch chars[bitIndex] {
		case '0':
			offBitLines = append(offBitLines, line)
		case '1':
			onBitLines = append(onBitLines, line)
		default:
			log.Fatal("There is not a bit value")
			log.Fatal(chars[bitIndex])
		}
	}

	if len(offBitLines) <= len(onBitLines) {
		return offBitLines
	} else {
		return onBitLines
	}
}
