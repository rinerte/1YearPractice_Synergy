package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var max = math.MinInt64
	var min = math.MaxInt64

	var totalSum, minSum, maxSum int

	for i := 1; i < len(os.Args); i++ {
		var number, error = strconv.Atoi(os.Args[i])
		if error != nil {
			panic(error)
		}

		
		if number < min {
			min = number
			minSum = totalSum
		}
		if number > max {
			max = number
			maxSum = totalSum
		}
		totalSum += number
	}

	var result = maxSum - minSum
	if result<0 {
		result = -result - max
	}
	fmt.Println(result)
}
