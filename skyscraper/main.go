//https://www.hackerrank.com/challenges/jim-and-the-skyscrapers
package main

import "fmt"

type sky struct {
	route  int32
	height int32
}

type stack []*sky

func (s *stack) Push(v *sky) {
	*s = append(*s, v)
}

func (s *stack) Pop() (*sky, int32) {

	l := len(*s)
	if l == 0 {
		return nil, -1
	}
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v, int32(l - 1)
}

func (s *stack) Len() int32 {
	return int32(len(*s))
}

func main() {
	var n, i, count int32
	fmt.Scanf("%d", &n)
	land := stack{}
	var h, previous, l int32
building:
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &h)
		if h < previous || land.Len() == 0 {
			land.Push(&sky{route: 0, height: h})
			previous = h
			continue
		}

		if previous == h {
			sk, _ := land.Pop()
			sk.route++
			count += sk.route
			land.Push(sk)
			previous = h
			continue
		}

		if h > previous {
			sk := &sky{route: 0, height: -1}
			for sk.height < h {
				sk, l = land.Pop()
				if l == -1 {
					land.Push(&sky{route: 0, height: h})
					previous = h
					continue building
				}
			}
			if sk.height == h {
				sk.route++
				land.Push(sk)
				count += sk.route
			} else {
				land.Push(&sky{route: 0, height: h})
			}
			previous = h
		}
	}

	fmt.Printf("%d", 2*count)
}
