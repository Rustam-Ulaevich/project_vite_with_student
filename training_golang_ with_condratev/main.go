package main

import "fmt"


func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}

var s int = 5

func main() {
	// Переменные :=
	/*
	score := 0.0
	text := "Get ready!!!"
	text2 := "Your score: "
	text3 := "You go truba!"
	boolean := true

	fmt.Println(boolean)
	fmt.Println(text)
	fmt.Println(text2, score)

	score = score + 1

	fmt.Println(text3)
	fmt.Println(text2, score)

	score := 50 
	*/

	// Условное ветвление if else
	/*
	if score > 10 {
		if score > 15 {
			fmt.Println("You megasupermen!!!")
		} else {
			fmt.Println("You good men!")
		}	
	}else{
		fmt.Println("You need, very work...")
	}

	if score > 15 {
		fmt.Println("You meggasupermaen!!!")
	}else if score > 10 {
		fmt.Println("You good men!")
	}else{
		fmt.Print("You need to work...")
	}

	if score == 12 {
		fmt.Println("Дюжина")
	} else if score == 21 {
		fmt.Println("Очко")
	} else if score == 50 {
		fmt.Println("Полтинник")
	} else {
		fmt.Println("Ни куда не попал")
	}

	if score >= 10 {
		fmt.Println("You win!")
	}

	if score != 7 {fmt.Println("...")}
	*/

	// Цикл for
	/*
	score := 2
	fmt.Println(score)

	for i := 1; i <= 5; i++ {
		score := 3
		score +=1
		fmt.Println("Итерация: ", i)
		fmt.Println(score)
	}

	fmt.Println(score)

	for {     //бесконечный цикл
		score +=1

		fmt.Println("Итерация: ", score)
		// fmt.Println(score)
		if rand.Intn(2000) == 1 {
			fmt.Println("The end!")
			break
		}			
	}
	
	fmt.Println(score)

	*/

	// Function

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)

	fmt.Println("До вызова функции")
	number := sum(1, 2)
	fmt.Println("После вызова функции")

	fmt.Println("number:", number)


}

func sum(a int, b int) int{
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	s := a + b

	return s
}

func a() int{
	h := 1 + s

	return h
}


