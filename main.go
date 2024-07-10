package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Println("Введите знак")
	fmt.Scanln(&a)
	if a == "+" || a == "-" {
		sub()
	} else if a == "*" || a == "/" {
		mis()
	}

}
func sub() {
	var a, b, c string

	fmt.Println("Введите первое значение")
	fmt.Scanln(&a)

	fmt.Println("Введите знак")
	fmt.Scanln(&b)

	fmt.Println("Введите второе значение")
	fmt.Scanln(&c)
	result1 := (a + c)
	result2 := strings.ReplaceAll(a, c, "")

	if b == "+" {
		fmt.Println(result1)
	} else if b == "-" {
		fmt.Println(result2)
	}
}
func mis() {
	var a, b string
	var c int

	fmt.Println("Введите слово")
	fmt.Scanln(&a)

	fmt.Println("Введите знак")
	fmt.Scanln(&b)

	fmt.Println("Введите цифру")
	fmt.Scanln(&c)

	result3 := strings.Repeat(a, c)
	_, result4 := fmt.Printf("%s\n", a[0:c])

	if b == "*" {
		fmt.Println(result3)
	} else if b == "/" {
		fmt.Println(result4)
	}
}
