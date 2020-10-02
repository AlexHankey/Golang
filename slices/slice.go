package main

import "fmt"

func main()  {
	s := make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)

	fmt.Println(s[0])

	fmt.Println(len(s))

	s = append(s, 4, 5, 6, 7)
	fmt.Println(s)

	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println(s[0:])

	// concise slice definition
	t := []int{100, 200, 300}
	fmt.Println(t)

	//x := s
	//x[0] = 500
	//fmt.Println(x)
	//fmt.Println(s)

	// copy prevents changing x and s
	x := make([]int, len(s))
	copy(x, s)
	x[0] = 900
	fmt.Println(x)
	fmt.Println(s)

	//2-d slice, similar to array although lengths of slice may vary
	ss := make([][]int, 3)
	for i := 0; i< 3; i++ {
		inner_len := i +1
		ss[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			ss[i][j] = i + j
		}
	}
	fmt.Println(ss)
}