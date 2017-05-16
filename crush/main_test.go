package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	b := bucket{start: 0, end: 2, value: 100}
	other := bucket{start: 0, end: 5, value: 100}
	_, res := b.add(other)
	for _, bb := range res {
		fmt.Printf("%#v\n", *bb)
	}

}
