package main

import "fmt"

func main() {

	diamondGenertor(6)
	fmt.Println()

	fmt.Println(ifFobidden(5, 4))
	fmt.Println(ifFobidden(-4, -7))
	fmt.Println(ifFobidden(2, 2))

	fmt.Println()

}

func diamondGenertor(n int) {

	for rows := 1; rows <= n; rows++ {
		for col := 1; col <= n-rows; col++ {
			fmt.Print(" ")
		}
		for cols := 1; cols <= rows*2-1; cols++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for rows := n - 1; rows >= 1; rows-- {
		for col := 1; col <= n-rows; col++ {
			fmt.Print(" ")
		}
		for cols := 1; cols <= rows*2-1; cols++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func ifFobidden(a, b int) string {
	var data string

	switch {
	case a > b:
		data = fmt.Sprintf("%v is greater than %v", a, b)
	case a < b:
		data = fmt.Sprintf("%v is greater than %v", b, a)
	default:
		data = fmt.Sprintf("%v is equal to %v", a, b)
	}

	return data
}
