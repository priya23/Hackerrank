package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var i int
	count := 0
	fmt.Scanf("%d", &i)
	arr := make([]int, i)
	for j, _ := range arr {
		fmt.Scanf("%d", &arr[j])
		count += arr[j]
	}
	fmt.Println(count)
}
