package pointers //-указатели

import "fmt"

func Pointers() {
	message := "Я тупой егор"
	MessageGay(message)
	//fmt.Println(message)
}

func MessageGay(message string) {
	message += "(Да я тупой егор сори !)"
	fmt.Println(message)
}
