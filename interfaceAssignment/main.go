package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea(float64, float64) float64
}

type triangle struct{}
type square struct{}

func main() {
	t := triangle{}
	s := square{}
	printArea(t)
	printArea(s)
}

func printArea(b shape) {
	theType := reflect.TypeOf(b)
	fmt.Println(b.getArea(2, 20), "and is type", theType)
}

func (t triangle) getArea(base float64, height float64) float64 {
	return base * height / 2
}

func (s square) getArea(base float64, height float64) float64 {
	return base * height
}
