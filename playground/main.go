package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		switch i {
		case 2:
			fmt.Println("Chegou no 2")
			continue
		case 3:
			fmt.Println("Quase lá")
		}
		fmt.Println(i)
	}
}
