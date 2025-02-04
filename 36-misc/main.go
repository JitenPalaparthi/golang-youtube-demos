package main

import (
	"fmt"
	"math"
	"math/big"
	"sync/atomic"
)

func main() {

	var a1, b1 int64 = math.MaxInt64, math.MaxInt64
	fmt.Println(math.MaxInt64) // 9223372036854775807
	fmt.Println(math.MinInt64) // -9223372036854775808
	c1 := a1 + b1              // adding two max values of int64
	fmt.Println(c1)            // it gives wrong result

	a2 := big.NewInt(9223372036854775807)
	b2 := big.NewInt(9223372036854775807)
	//c2 := big.NewInt(0)
	c2 := new(big.Int).Add(a2, b2)
	fmt.Println("sum:", c2)

	m2 := new(big.Int).Mul(c2, b2)
	fmt.Println("mul:", m2)

	var a, b, c int64 = 10, 20, 30

	d := (a+b)/2 + a*b + a/b + c + a + b + c*10
	fmt.Println(d) // what is an expression
	// how an expression is split into multiple commands from asm perspective

	// 15 + 200 + 0 + 30+10+20 + 300
	// 15 + 200 +0 + 60 + 300
	// 565
	// add
	// mov
	// sub
	// mul
	// div

	// a+b * a-b
	var result int64
	vala := atomic.LoadInt64(&a)
	valb := atomic.LoadInt64(&b)
	localresult := (vala + valb) * (vala - valb)
	//atomic.AddInt64(&a, b)
	atomic.StoreInt64(&result, localresult)
	fmt.Println(result)

}
