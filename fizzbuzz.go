package main

import "fmt"

func main() {
	fmt.Println("3の倍数のときfizz, 4の倍数のときbuzz, 12の倍数のときfizzbuzz")
	for i := 1; i < 20; i++ {
		if i%12 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%4 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
