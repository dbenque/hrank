package main

import "fmt"

func main() {

	var N, Q int
	var line string
	fmt.Scanf("%d", &N)
	data := map[string]int{}
	for i := 0; i < N; i++ {
		fmt.Scanln(&line)
		if v, ok := data[line]; ok {
			v++
			data[line] = v
		} else {
			data[line] = 1
		}
	}

	fmt.Scanf("%d", &Q)
	for i := 0; i < Q; i++ {
		fmt.Scanln(&line)
		if v, ok := data[line]; ok {
			fmt.Printf("%d\n", v)
		} else {
			fmt.Println("0")
		}
	}
}
