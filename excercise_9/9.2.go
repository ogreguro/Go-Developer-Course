package main

import "fmt"

func main() {
	v := []int{1, 2, 3}

	for _, v1 := range v {
		fmt.Println("v1:", v1)
	BREAK_POINT:
		for _, v2 := range v {
			fmt.Println("\tv2:", v2)
			for _, v3 := range v {
				fmt.Println("\t\tv3:", v3)
				for _, v4 := range v {
					fmt.Println("\t\t\tv4:", v4)
					if v4 == 2 {
						break BREAK_POINT
					}
				}
			}
		}
	}
}
