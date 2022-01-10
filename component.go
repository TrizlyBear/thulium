package main

import "fmt"

type Num interface {
	float64 | int
}

func IsGreater[N Num](x N, y N) bool{
	return x > y
}

func main()  {
	fmt.Println(IsGreater(3.0, 2.1))
}
