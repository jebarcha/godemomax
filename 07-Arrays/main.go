package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices[1] = 9.99

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

}

//func main() {
// var productNames [4]string = [4]string{"A Book"}
// prices := []float64{10.99, 9.99, 45.99, 20.0}
// fmt.Println(prices)

// productNames[2] = "A Carpet"
// fmt.Println(productNames)

// fmt.Println(prices[1])

// //Slices
// //featuredPrices := prices[1:3]  //start index 1 until 3 excluding 3
// featuredPrices := prices[1:]
// highligtedPrices := featuredPrices[:1]
// fmt.Println("featuredPrices:", featuredPrices)
// fmt.Println("highligtedPrices:", highligtedPrices)

// fmt.Println(len(featuredPrices), cap(featuredPrices))
// fmt.Println(len(highligtedPrices), cap(highligtedPrices))

// highligtedPrices = highligtedPrices[:3]
// fmt.Println(len(highligtedPrices), cap(highligtedPrices))
//}
