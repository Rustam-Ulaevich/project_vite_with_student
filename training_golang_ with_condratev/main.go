package main

import (
    "fmt"
)


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
	/*
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
	*/

	// Defer
	/*
    path := filepath.Join(os.TempDir(), "defer.txt")
    f := createFile(path)
    defer closeFile(f)
    writeFile(f)
	*/

	// Указатели

	b := 5
	a := &b

	fmt.Println(a, *a)
	
	var ptr *int  // nill указатель

	fmt.Print(ptr)
	
}

// func createFile(p string) *os.File {
//     fmt.Println("creating")
//     f, err := os.Create(p)
//     if err != nil {
//         panic(err)
//     }
//     return f
// }

// func writeFile(f *os.File) {
//     fmt.Println("writing")
//     fmt.Fprintln(f, "data")
// }

// func closeFile(f *os.File) {
//     fmt.Println("closing")
//     err := f.Close()

//     if err != nil {
//         panic(err)
//     }
// }





