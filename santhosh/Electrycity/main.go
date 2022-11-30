package main

import (
	"Electrycity/users"
	"fmt"
)

var namecheck string
var passcheck string

func main() {
	var mainoption int
	var ext string
	us := users.Users_list{}
	fmt.Print("Welcome to the application!...\n")
	for ok := true; ok; ok = ext == "y" {
		print("1.Login 2.Sign Up 3.show Users\n")
		println("Select the option Above : ")
		fmt.Scanln(&mainoption)
		switch mainoption {
		case 1:
			login(&us)
		case 2:
			signup(&us)

		case 3:
			us.ShowUsers()
		default:
			fmt.Println("select the correct option!..")
		}
		fmt.Print("Do you want to go mainmenu (y/n)..: ")
		fmt.Scanln(&ext)
	}

}
func login(user *users.Users_list) {
loop:
	print("Enter Your Name : ")
	fmt.Scanln(&namecheck)
	print("Enter Your Password : ")
	fmt.Scanln(&passcheck)
	if user.LoginCheck(namecheck, passcheck) {
		fmt.Println("Logged in")
	} else {
		fmt.Println("Entered wrong details enter again...!")
		goto loop
	}
}

func signup(users *users.Users_list) {
	var pass string
	var name string
	var age int
	print("Enter Your Name : ")
	fmt.Scanln(&name)
	print("Enter you Paasword : ")
	fmt.Scanln(&pass)
	print("Enter Your Age : ")
	fmt.Scanln(&age)
	users.AddUser(name, pass, age)
}
