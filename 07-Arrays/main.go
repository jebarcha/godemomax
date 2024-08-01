package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//userNames := []string{}
	userNames := make([]string, 2, 5)

	userNames[0] = "Ringo"

	userNames = append(userNames, "Paul")
	userNames = append(userNames, "John")

	userNames = append(userNames, "George")
	userNames = append(userNames, "Beatles")

	//fmt.Println(userNames)      //[Ringo  Paul John]
	//fmt.Println(cap(userNames)) //5

	//courseRatings := map[string]float64{}
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.5
	courseRatings["node"] = 4.9

	//courseRatings.output()

	//fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}

}
