package main

import "fmt"

func main() {

	for x := 1; x <= 10; x++ {
		fmt.Printf("%d\t", x)
	}

	fmt.Println("")

	y := 1
	for y <= 10 {
		fmt.Printf("%d\t", y)
		y++
	}

	fmt.Println("")

	z := 1
	for {
		if z > 10 {
			break
		}
		fmt.Printf("%d\t", z)
		z++

	}
	fmt.Println("\nDone.")

}
