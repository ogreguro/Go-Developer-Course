package main

import "fmt"

func main() {
	recursiveFunc(23, 0, 1)
}

func recursiveFunc(runCount int, valFirst int, valSecond int) {
	if runCount == 0 {
		return
	}
	runCount--

	fmt.Println(valFirst)

	valNew := valFirst + valSecond
	valFirst = valSecond
	valSecond = valNew

	recursiveFunc(runCount, valFirst, valSecond)
}
