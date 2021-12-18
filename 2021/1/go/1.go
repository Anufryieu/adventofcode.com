package main

import (	
	"fmt"
	"log"
	"os"	
	"strings"
	"strconv"
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

	fmt.Println(first(values))
}

func first(values []int) int {
	var previosValue int
	count := 0

	for index, value := range values {
		if (index > 0 && value > previosValue) {
			count++
		}

		previosValue = value
	}

	return count
}