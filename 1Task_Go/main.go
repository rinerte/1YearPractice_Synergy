package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	result, err := computeSumBetweenMaxAndMin(os.Args)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func computeSumBetweenMaxAndMin(args []string) (int, error) {
	var max = math.MinInt64
	var min = math.MaxInt64

	var totalSum, minSum, maxSum int
	var right = true

	for i := 1; i < len(args); i++ {
		number, err := strconv.Atoi(args[i])
		if err != nil {
			return 0, err
		}

		if number < min {
			min = number
			minSum = totalSum
			right = false
		}
		if number > max {
			max = number
			maxSum = totalSum
			right = true
		}
		totalSum += number
	}
	var result = 0
	if right {
		result = maxSum - minSum - min
	} else {
		result = minSum - maxSum - max
	}

	return result, nil
}
