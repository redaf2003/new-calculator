package introduction

import (
	calculatingvalues "calculator/calculatingValues"
	"fmt"
)

func Conduction() {
	fmt.Print("Ведите 1-ое число:")
	var num1 float64 //-Пользователь пишет число в консоль (вся функция)
	fmt.Scan(&num1)

	fmt.Print("Введите операцию:(+, -, *, /): ")
	var operator string
	fmt.Scan(&operator)

	fmt.Print("Ведите 2-ое число:")
	var num2 float64 //-Пользователь пишет 2-ое исло в консоль (вся функция)
	fmt.Scan(&num2)

	result, err := calculatingvalues.Calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n",
		num1, operator, num2, result)
}

//func Conduction2() {
//	fmt.Print("Введите операцию (+, -, *, /): ")
//	var operator float64
//	fmt.Scan(&operator)
//}

//func Conduction3() {
//	fmt.Print("Ведите 2-ое число =")
//	var num2 float64 //-Пользователь пишет 2-ое число в консоль (вся функция)
//	fmt.Scan(&num2)
//}
