package pointers //-указатели

import (
	"fmt"
)

func Pointers() {
	message := "Я тупой егор"
	MessageGay(&message)
	fmt.Println(message)
}

func MessageGay(message *string) {
	*message += "(Да я тупой егор сори!)"
	fmt.Println(message)
}

func Numbersss() {
	fmt.Println("Указатели:")
	number := 5
	p := &number
	fmt.Println(number)
	fmt.Println(p)

	*p = 10
	fmt.Println(number)
	fmt.Println("______________________")

}
