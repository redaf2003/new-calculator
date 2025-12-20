package introduction

import "fmt"

func Conduction1() {
	fmt.Print("Ведите 1-ое число =")
	var num1 float64 //-Пользователь пишет число в консоль (вся функция)
	fmt.Scan(&num1)
}
func Conduction2() {
	fmt.Print("Введите операцию (+, -, *, /): ")
	var operator float64
	fmt.Scan(&operator)
}

func Conduction3() {
	fmt.Print("Ведите 2-ое число =")
	var num2 float64 //-Пользователь пишет 2-ое число в консоль (вся функция)
	fmt.Scan(&num2)
}
