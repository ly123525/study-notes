package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := intSum(a, b)
	fmt.Println(c)

	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	var d calculation
	d = add

	fmt.Printf("type of d:%T\n", c) // type of d:main.calculation
	fmt.Println(d(1, 2))            // 像调用add一样调用d

	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	ff := adder()
	fmt.Println(ff(40)) //40
	fmt.Println(ff(50)) //90

	fmt.Println("=============================")

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}

func intSum(x, y int) int {
	return x + y
}

func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

// func someFunc(x string) []int {
// 	if x == "" {
// 		return nil // 没必要返回[]int{}
// 	}

// }

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type calculation func(int, int) int

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
