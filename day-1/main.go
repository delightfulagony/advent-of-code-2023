package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	solution "aoc23-1/solution"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	calibrationValue, err := solution.CalibrationValueParser(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calibrationValue)

}

