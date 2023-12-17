// author: delightfulagony<agony@disroot.org> year: 2023

package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"flag"
	solution "aoc23-1/solution"
)

func main() {

	var improved *bool = flag.Bool("improved", false,
			"Selects between normal and improved mode")
	flag.Parse()

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

	var calibrationValue int
	if *improved {
		calibrationValue, err = solution.ImprovedCalibrationValueParser(input)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		calibrationValue, err = solution.CalibrationValueParser(input)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(calibrationValue)

}

