package main

import (
	"fmt"
	//"math/rand"
	"os"
	"strconv"
)

func main() {
	var samp int
	fmt.Scanf("%d", &samp)

	for i := 0; i < samp; i++ {
		var x, y int64
		count := 0
		fmt.Fscanln(os.Stdin, &x, &y)
		//x1 := strconv.FormatInt(x, 2)
		//y1 := strconv.FormatInt(y, 2)
		z := (x ^ y)
		for y, x := range strconv.FormatInt(z, 2) {
			fmt.Println("\n ", x, y)
			if x == 49 {
				fmt.Println("inside if")
				count += 1
			}
		}
		fmt.Println("VAL IS %T %T", x, y, z, count)
	}
}
