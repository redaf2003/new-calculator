package variables

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

	// Инкремент и декремент (score++ и score--)
	fmt.Println("\n--- Инкремент и декремент ---")

	score := 5
	fmt.Println("Исходный score:", score)

	// Инкремент (увеличение на 1)
	score++
	fmt.Println("После score++:", score)

	// Декремент (уменьшение на 1)
	score--
	fmt.Println("После score--:", score)

	// Множественный инкремент
	score++
	score++
	score++
	fmt.Println("После трех score++:", score)

	// Использование в цикле
	fmt.Println("\nИнкремент в цикле:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Итерация %d, значение i: %d\n", i+1, i)
	}

	// Арифметические операции
	sum := number1 + number2 + number3
	fmt.Printf("\nСумма всех чисел: %d + %d + %d = %d\n", number1, number2, number3, sum)

	// Конкатенация строк
	greeting := text1 + " " + text2 + " " + text3
	fmt.Println("Объединённая строка:", greeting)

	// Логические операции
	boolResult := boolean1 && boolean2 || boolean3
	fmt.Printf("Результат логического выражения (boolean1 && boolean2 || boolean3): %v\n", boolResult)

	// Вывод типов переменных
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

	// Примеры с разными типами инкремента/декремента
	fmt.Println("\n--- Дополнительные примеры с инкрементом ---")

	counter := 0
	fmt.Println("Начальное значение counter:", counter)

	// Префиксный и постфиксный инкремент
	counter = 10
	fmt.Println("counter до инкремента:", counter)
	counter++ // постфиксный инкремент
	fmt.Println("counter после counter++:", counter)

	// Важное замечание: В Go нет префиксного ++counter или --counter
	// ++ и -- могут использоваться только как постфиксные операторы

	// Пример с float
	floatCounter := 0.5
	fmt.Println("\nFloat счетчик:", floatCounter)
	// floatCounter++ // ОШИБКА: в Go инкремент/декремент работает только с целыми числами
	floatCounter += 1.0 // Правильный способ для float
	fmt.Println("После floatCounter += 1.0:", floatCounter)

	// Практический пример: подсчет очков в игре
	fmt.Println("\n--- Пример: Игровые очки ---")
	gameScore := 100
	fmt.Println("Начальный счет:", gameScore)

	// Игрок получает очки
	gameScore += 25 // Получил бонус
	fmt.Println("После бонуса (+25):", gameScore)

	gameScore++ // За выполненное задание
	fmt.Println("За задание (+1):", gameScore)

	gameScore += 10 // Победа над врагом
	fmt.Println("За победу (+10):", gameScore)

	gameScore-- // Штраф за смерть
	fmt.Println("Штраф за смерть (-1):", gameScore)

	gameScore -= 5 // Покупка предмета
	fmt.Println("Покупка предмета (-5):", gameScore)

	fmt.Println("Финальный счет:", gameScore)
}
