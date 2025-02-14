package main

import "fmt"

func main() {
	//result := c1.Add(100).Add(200).Sub(20).Mul(30).Get()

	c1 := NewCalc(20)

	r := c1.Add(100).Add(200).Sub(220).Add(400).Mul(2).Mul(3).Div(10).Get()

	fmt.Println("Result:", r)
}

type ICalc interface {
	Add(num int) ICalc
	Sub(num int) ICalc
	Mul(num int) ICalc
	Div(num int) ICalc
	Get() int
}

type Calc struct {
	value int
}

func NewCalc(v int) *Calc {
	return &Calc{value: v}
}

func (c *Calc) Add(n int) ICalc {
	c.value += n
	return c
}

func (c *Calc) Sub(n int) ICalc {
	c.value -= n
	return c
}

func (c *Calc) Mul(n int) ICalc {
	c.value *= n
	return c
}

func (c *Calc) Div(n int) ICalc {
	c.value /= n
	return c
}

func (c *Calc) Get() int {
	return c.value
}
