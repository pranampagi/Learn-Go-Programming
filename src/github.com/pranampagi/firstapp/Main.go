package main

import "fmt"

// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	a2 = iota
// )

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// const myConst int = 42
	// fmt.Printf("%v, %T\n", myConst, myConst)

	// const a int = 14
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)

	// const a = 42
	// var b int = 27
	// fmt.Printf("%v, %T\n", a+b, a+b)

	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", a2)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
