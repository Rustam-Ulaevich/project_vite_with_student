package main

import (
	"fmt"
	// "go/types"
	// "study/greeting"
)

type User struct {
	Name string
	Rating float64
	Premium bool
}


func main() {

/*

	// Переменные :=

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


	// Условное ветвление if else

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


	// Цикл for

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


	// Defer

   path := filepath.Join(os.TempDir(), "defer.txt")
   f := createFile(path)
   defer closeFile(f)
   writeFile(f)


	// Указатели

	b := 5
	a := &b

	fmt.Println(a, *a)  // * разименование указателя

	var ptr *int  // nil указатель

	fmt.Print(ptr)


	// Structures

	type User struct {
		Name        string  // ""
		Age         int     // 0
		PhoneNumber string  // ""
		IsClose     bool    // false
		Rating      float64 // 0.0
	}

	func (u User) RatingUp(rating float64) {
		if u.Rating+rating <= 10 {
		u.Rating += rating
		fmt.Println("Рейтинг стал:", u.Rating)
		} else {fmt.Println("Не пройдена валидация")}
	}

	func (u *User) ChangeName(newName string) {
		if newName != "" {
		u.Name = newName
		} else {fmt.Println("Получено пустое имя")}
	}

	func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	} else {
		fmt.Println("Получено не валидное число")
	}
	}

	func (u *User) RatingUpp(value float64) {
	if u.Rating + value <= 10.0 {
		u.Rating += value
	}
	}

	func (u *User) RatingDown(value float64) {
	if u.Rating - value >= 0.0 {
		u.Rating -= value
	}
	}

	func (u *User) ChangeCloseAccount() {
	u.IsClose = !u.IsClose
	}

	func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
	) User{
	 	if name == "" {
	 		return User{}
		}
		if age <= 0 || age >= 150 {
			return User{}
		}
		if phoneNumber == "" {
			return User{}
		}
		if isClose != types.IsBoolean {
			return  User{} 
		}
		if rating < 0.0 || rating > 10.0 {
			return User{}
		}

		return User{ Name: name, Age: age, PhoneNumber: phoneNumber, IsClose: isClose, Rating: rating,}
	}
*/


	// Arrays

	// arr := [5] int {2, 3, 5, 10, 1} 

	// for i := 0; i < 4; i++ {
	// 	if arr[i]%2 == 0 {
	// 		arr[i] *= 2
	// 	}
	// 	fmt.Println( i, "-", arr[i])
	// }

	userArray := [12]User{
		User{
			Name: "Mike",
			Rating: 5.5,
			Premium: true,
		},
		User{
			Name: "Nina",
			Rating: 7.5,
			Premium: true,
		},
		User{
			Name: "Ura",
			Rating: 1.2,
			Premium: true,
		},
		User{
			Name: "Mike",
			Rating: 5.5,
			Premium: true,
		},
		User{
			Name: "Nina",
			Rating: 7.5,
			Premium: true,
		},
		User{
			Name: "Ura",
			Rating: 1.2,
			Premium: true,
		},
		User{
			Name: "Mike",
			Rating: 5.5,
			Premium: true,
		},
		User{
			Name: "Nina",
			Rating: 7.5,
			Premium: true,
		},
		User{
			Name: "Ura",
			Rating: 1.2,
			Premium: true,
		},
		User{
			Name: "Mike",
			Rating: 5.5,
			Premium: true,
		},
		User{
			Name: "Nina",
			Rating: 7.5,
			Premium: true,
		},
		User{
			Name: "Ura",
			Rating: 1.2,
			Premium: true,
		},
	}

	for index, value := range userArray {
		value.Rating += 1
		value.Premium = false
		value.Name += "-loh"
		fmt.Println(index, value)
	}
	

}	



// }

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
