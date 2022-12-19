package main

import "fmt"

func main() {
	// var testArray [3]int
	// var numArray = [3]int{1, 2}
	// var cityArray = [3]string{"北京", "上海", "广州"}

	// fmt.Println(testArray)
	// fmt.Println(numArray)
	// fmt.Println(cityArray)

	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "广州"}

	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Printf("Type Of numArray: %T \n", numArray)
	fmt.Println(cityArray)
	fmt.Printf("Type Of cityArray: %T \n", cityArray)

	a := [...]int{1: 3, 3: 5}
	fmt.Println(a)

	//遍历
	var cities = [...]string{"北京", "上海", "广州"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for index, value := range cities {
		fmt.Println(index, value)
	}
}
