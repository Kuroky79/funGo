package main

import (
	"fmt"
	"math"
)

/* Реализация FizzBuzz на Go */
func main() {
	var (
		a  float64
		b  float64
		c  float64
		x1 float64
		x2 float64
	)
	fmt.Println("Реши квадратное уравнение")
	fmt.Println("Введи a:")
	fmt.Scan(&a)
	fmt.Println("Введи b:")
	fmt.Scan(&b)
	fmt.Println("Введи c:")
	fmt.Scan(&c)
	D := (b * b) - 4*a*c
	if D < 0 {
		fmt.Println("Корней нет")
	} else if D == 0 {
		x1 = -b / 2 * a
		fmt.Println("Единственный корень", x1)
	} else if D > 0 {
		x1 = ((-b + math.Sqrt(D)) / 2 * a)
		x2 = ((-b - math.Sqrt(D)) / 2 * a)
		fmt.Println("Первый корень:", x1, "Второй корень:", x2)
	}

}
