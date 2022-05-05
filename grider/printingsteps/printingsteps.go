package printingsteps

import "fmt"

func PrintingSteps(n int) {
	steps := ""

	for row := 1; row <= n; row++ {
		for col := 1; col <= row; col++ {
			steps = steps + "#"
		}

		if row != n {
			steps = steps + "\n"
		}
	}

	fmt.Println(steps)
}

func PrintingSteps2(n int) {
	steps := ""

	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			if col > n-row {
				steps = steps + "#"
				continue
			}

			steps = steps + " "
		}

		if row != n {
			steps = steps + "\n"
		}
	}

	fmt.Println(steps)
}

// Pyramid
// func PrintingSteps2(n int) {
// 	steps := ""

// 	for row := 1; row <= n; row++ {
// 		for col := n; col >= 1; col-- {
// 			steps = steps + " "

// 			if col <= row {
// 				steps = steps + "#"
// 			}
// 		}

// 		if row != n {
// 			steps = steps + "\n"
// 		}
// 	}

// 	fmt.Println(steps)
// }
