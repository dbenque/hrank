package main

import "fmt"

func main() {
	var v, d, n uint32
	fmt.Scanf("%d %d", &n, &d)
	array := make([]uint32, d)
	for i := uint32(0); i < d; i++ {
		fmt.Scanf("%d", &v)
		array[i] = v
	}

	if d == 0 {
		n--
	}

	for i := d; i < n; i++ {
		fmt.Scanf("%d", &v)
		fmt.Printf("%d ", v)
	}

	if d == 0 {
		fmt.Scanf("%d", &v)
		fmt.Printf("%d", v)
		return
	}

	for i, v := range array {
		fmt.Printf("%d", v)
		if uint32(i) < d-1 {
			fmt.Printf(" ")
		}
	}
}
