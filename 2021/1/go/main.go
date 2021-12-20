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
	items := strings.Fields(content)
	values := make([]int, len(items))

	for index, item := range items {
		value, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}

		values[index] = value
	}

	fmt.Println("How many measurements are larger than the previous measurement?")
	fmt.Println(count_of_values_that_larger_than_previous(values))

	fmt.Println("How many sums are larger than the previous sum?")
	fmt.Println(count_of_values_that_larger_than_previous(sums(values)))
}

func count_of_values_that_larger_than_previous(values []int) int {
	var previosValue int
	count := 0

	for index, value := range values {
		if index > 0 && value > previosValue {
			count++
		}

		previosValue = value
	}

	return count
}

func sums(values []int) []int {
	sums := make([]int, len(values)-2)

	for i := 0; i < len(values)-2; i++ {
		sums[i] = values[i] + values[i+1] + values[i+2]
	}

	return sums
}
