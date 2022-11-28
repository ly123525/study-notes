package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s): %v\n", s, len(s), cap(s))

	s1 := make([]int, 2, 10)
	fmt.Println(s1)      //[0 0]
	fmt.Println(len(s1)) //2
	fmt.Println(cap(s1)) //10

	s2 := s1 //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0]
	fmt.Println(s2) //[100 0]

	//遍历

	s3 := []int{1, 2, 3}
	for i := 0; i < len(s3); i++ {
		fmt.Println(i, s3[i])
	}

	for index, value := range s3 {
		fmt.Println(index, value)
	}

	// append
	var s4 []int
	s4 = append(s4, 1, 2, 3)
	fmt.Println(s4)

	var cities []string
	cities = append(cities, "北京")
	cities = append(cities, "上海", "广州")
	s5 := []string{"深圳", "成都", "重庆"}
	cities = append(cities, s5...)
	fmt.Println(cities) //[北京 上海 广州 深圳 成都 重庆]

	s6 := []int{1, 2, 3, 4, 5, 6, 7}
	s6 = append(s6[:2], s6[3:]...)
	fmt.Println(s6)

	var s7 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%v", i)
		fmt.Println(str)
		s7 = append(s7, str)
	}
	fmt.Println(s7, len(s7))
}
