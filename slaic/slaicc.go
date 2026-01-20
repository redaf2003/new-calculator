package slaic //-слайс

import (
	"fmt"
)

func FuxMake() {
	fmt.Println("")
	fmt.Println("Функция майк:")
	message := make([]string, 5)
	message = append(message, "6")
	//message[6] = "6" //-нельзя так делать
	fmt.Println("Длина:", len(message))   //-длина слайса, len()
	fmt.Println("Емкость:", cap(message)) //-емкость нашего слайса cap()
	fmt.Println(message)
	//fmt.Println(message)
	fmt.Println("")
}

func Slice() {
	fmt.Println("")
	fmt.Println("Слайс:")
	message := []string{"1", "7", "3"}

	fmt.Println("Обычный слайс :", message)
	fmt.Println("")
}
