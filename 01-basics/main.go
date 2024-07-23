package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var inventmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&inventmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := inventmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println((futureValue))
	fmt.Println(futureRealValue)
}
