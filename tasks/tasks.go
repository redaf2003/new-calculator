package tasks

import "fmt"

func Proverka1() {

	score1 := 0

	for i := 1; i <= 5; i++ {
		//score := 5
		score1++
		fmt.Println(score1)
		return //- делает стоп для всей функции

	}
	fmt.Println("Серега иди нахуй") //- это не выведится
}

func Proverka2() {

	score2 := 0

	for i := 1; i <= 5; i++ {
		//score := 5
		score2++
		fmt.Println(score2)
		break // -стоп для цикла (выход из цикла)
	}
	fmt.Println("Серега иди нахуй") //- это выведится
}

func Practic1(name string) {

	//if name == "1" {
	//	fmt.Println("Это бля цыфра ")
	//	return
	//	}

	if len(name) == 0 {
		fmt.Println("Имя не может быть пустым")
	} else {
		fmt.Println("Здравствуйте!", name)
		return
	}
}

func Practic2(age int) {

	if age < 5 {
		fmt.Println("Слишком маленький возраст ")
		return
	}

	if age > 100 {
		fmt.Println("Слишком большой возраст ")
		return
	}

	if age < 0 {
		fmt.Println("Возраст не может быть отрицательным ")
		return
	}

	if age == 0 {
		fmt.Println("Возраст не может быть равен 0")
		return
	}

	if age < 18 {
		fmt.Println("НЕ ПРОХОДИТЕ!!!!")
	} else {
		fmt.Println("Все четко проходите ")
		return
	}
}

func Practic3() {
	Practic1("2")
	Practic2(19)

}
