package main

import "fmt"

// type bucket struct {
// 	start, end int
// 	value      int
// }

// func (b *bucket) add(other bucket) (bool, []*bucket) {
// 	if b.start >= other.end || other.start >= b.end { // bucket not joined
// 		return false, []*bucket{b}
// 	}

// 	if other.start >= other.end {
// 		return false, []*bucket{b}
// 	}

// 	if other.start < b.start {
// 		other.start = b.start
// 	}

// 	if other.end > b.end {
// 		other.end = b.end
// 	}

// 	res := []*bucket{}

// 	if other.start != b.start {
// 		res = append(res, &bucket{start: b.start, end: other.start, value: b.value})
// 	}

// 	res = append(res, &bucket{start: other.start, end: other.end, value: b.value + other.value})

// 	if other.end != b.end {
// 		res = append(res, &bucket{start: other.end, end: b.end, value: b.value})
// 	}
// 	return true, res
// }

// func main() {
// 	var N, M, a, b, k, i int
// 	fmt.Scanf("%d %d", &N, &M)

// 	l := list.New()
// 	l.PushBack(&bucket{start: 0, end: N, value: 0})
// 	for i = 0; i < M; i++ {
// 		fmt.Scanf("%d %d %d", &a, &b, &k)
// 		a--
// 		other := bucket{start: a, end: b, value: k}
// 		for e := l.Front(); e != nil; e = e.Next() {
// 			b := (e.Value).(*bucket)
// 			if ok, res := b.add(other); ok {
// 				esave := e
// 				for _, r := range res {
// 					e = l.InsertAfter(r, e)
// 				}
// 				l.Remove(esave)
// 			}
// 		}
// 		fmt.Printf("Line %d, count bucket:%d\n", i, l.Len())
// 	}

// 	max := 0
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		b := (e.Value).(*bucket)
// 		if b.value > max {
// 			max = b.value
// 		}
// 	}
// 	fmt.Printf("%d", max)
// }

func main() {
	var N, M, a, b, k, i int64
	fmt.Scanf("%d %d", &N, &M)
	res := make([]int64, N)
	for i = 0; i < M; i++ {
		fmt.Scanf("%d %d %d", &a, &b, &k)
		a--
		res[a] = res[a] + k
		if b < N {
			res[b] = res[b] - k
		}
	}
	max := int64(0)
	current := int64(0)
	for _, v := range res {
		current += v
		if current > max {
			max = current
		}
	}
	fmt.Printf("%d", max)
}
