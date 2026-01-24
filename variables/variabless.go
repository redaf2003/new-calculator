package main

import "fmt"

// Глобальные переменные (объявлены вне функции main)
var (
	number1  int     = 10
	text1    string  = "hello"
	float1   float64 = 1.1
	boolean1 bool    = true
)

func Peremenie() {
	// Вывод глобальных переменных
	fmt.Println("number1:", number1)
	fmt.Println("text1:", text1)
	fmt.Println("float1:", float1)
	fmt.Println("boolean1:", boolean1)

	// Локальные переменные (объявлены внутри функции main)
	var number2 int = 20
	var text2 string = "world"
	var float2 float64 = 2.2
	var boolean2 bool = false

	fmt.Println("number2:", number2)
	fmt.Println("text2:", text2)
	fmt.Println("float2:", float2)
	fmt.Println("boolean2:", boolean2)

	// Дополнительные примеры для лучшего понимания
	fmt.Println("\n--- Дополнительные примеры ---")

	// Краткое объявление переменных (только внутри функций)
	number3 := 30
	text3 := "golang"
	float3 := 3.3
	boolean3 := true

	fmt.Println("number3:", number3)
	fmt.Println("text3:", text3)
	fmt.Println("float3:", float3)
	fmt.Println("boolean3:", boolean3)

	// Арифметические операции
	sum := number1 + number2 + number3
	fmt.Printf("\nСумма всех чисел: %d + %d + %d = %d\n", number1, number2, number3, sum)

	// Конкатенация строк
	greeting := text1 + " " + text2 + " " + text3
	fmt.Println("Объединённая строка:", greeting)

	// Логические операции
	boolResult := boolean1 && boolean2 || boolean3
	fmt.Printf("Результат логического выражения (boolean1 && boolean2 || boolean3): %v\n", boolResult)

	// Вывод типов переменных (используем пакет reflect)
	fmt.Println("\n--- Типы переменных ---")
	fmt.Printf("number1: %T\n", number1)
	fmt.Printf("text1: %T\n", text1)
	fmt.Printf("float1: %T\n", float1)
	fmt.Printf("boolean1: %T\n", boolean1)

	// Константы
	const pi = 3.14159
	const appName = "Go Learning App"
	fmt.Println("\n--- Константы ---")
	fmt.Println("Число Пи:", pi)
	fmt.Println("Название приложения:", appName)
}
