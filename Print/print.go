package Print

import "fmt"

func Print(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
}