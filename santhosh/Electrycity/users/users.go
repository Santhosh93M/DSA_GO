package users

import "fmt"

type Users struct {
	username string
	userpass string
	userage  int
	next     *Users
}

type Users_list struct {
	head *Users
	len  int
}

func (u *Users_list) AddUser(name string, pass string, age int) {
	n := Users{username: name, userpass: pass, userage: age}
	if u.len == 0 {
		u.head = &n
		u.len++
		return
	}
	ptr := u.head
	for i := 0; i < u.len; i++ {
		if ptr.next == nil {
			ptr.next = ptr
			u.len++
			return
		}
		ptr = ptr.next
	}
}
func (u *Users_list) ShowUsers() {
	n := u.head
	if u.len == 0 {
		fmt.Println("No users in list..!")
		return
	}
	for i := 0; i < u.len; i++ {
		fmt.Println("---------------------------------")
		fmt.Printf("UserName = %v \nUserPass = %v \nAge = %d \n", n.username, n.userpass, n.userage)
		n = n.next
	}
	fmt.Println()
}
func (u *Users_list) GetPos(pos int) *Users {
	ptr := u.head
	if pos < 0 {
		return ptr
	}
	if pos > (u.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (u *Users_list) LoginCheck(name string, password string) bool {
	if u.len == 0 {
		fmt.Println("No users found")
		return false
	}
	temp := u.head
	for i := 0; i < u.len; i++ {
		if name == temp.username && password == temp.userpass {
			return true
		}
		temp = temp.next
	}
	return false
}
