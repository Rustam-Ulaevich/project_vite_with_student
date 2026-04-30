package main

import "fmt"

func main() {
	// score := 0.0
	// text := "Get ready!!!"
	// text2 := "Your score: "
	// text3 := "You go truba!"
	// boolean := true

	// fmt.Println(boolean)
	// fmt.Println(text)
	// fmt.Println(text2, score)

	// score = score + 1

	// fmt.Println(text3)
	// fmt.Println(text2, score)

	score := 50

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

	if score != 7 {
		fmt.Println("...")
	}
}