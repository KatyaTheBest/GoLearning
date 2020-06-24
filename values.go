//пример работы с переменными
package main

import (
	"fmt"
	"math"
)

func main() {

	var v, m int32 = 22, 10
	fmt.Println("v*m=", v*m)

	var s float64 = math.Sqrt(121)
	fmt.Println("s=", s)

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
