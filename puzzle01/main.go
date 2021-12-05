package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//To do this, count the number
//of times a depth measurement increases from the previous measurement.
//(There is no measurement before the first measurement.) In the example above, the changes are as follows:
func countDepthMeasurenmentIncrements(input []int) int {
	var (
		count           int
		previousMeasure int
	)

	for i, actual := range input {
		if i == 0 {
			previousMeasure = actual
			continue
		}

		if previousMeasure < actual {
			count++
		}

		previousMeasure = actual
	}
	return count
}

func countDepthMeasurenmentIncrementsWindow(input []int) int {
	counter := 0
	previousGroupMeasure := input[0] + input[1] + input[2]

	for i := range input[2:] {
		if len(input[i:]) < 3 {
			break
		}

		nextGroupSum := input[i] + input[i+1] + input[i+2]

		if nextGroupSum > previousGroupMeasure {
			counter++
		}

		previousGroupMeasure = nextGroupSum
	}

	return counter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var input []int

	for scanner.Scan() {
		measure, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		input = append(input, measure)
	}

	log.Println(countDepthMeasurenmentIncrements(input))
	log.Println(countDepthMeasurenmentIncrementsWindow(input))
}
