package main

import "fmt"

type User struct{
	Name string
	LastName string
	Age int
	Email string
	Password string
}

func ChangeName(name string, user *User){
	(*user).Name = name
}

func ChangeAge(age int,user *User){
	(*user).Age = age
}

func ChangeEmail(email string, user *User){
	(*user).Email = email
}

func ChangePassword(password string, user *User){
	(*user).Password = password
}

func main(){
	var user User
	ChangeName("Sol",&user)
	ChangeAge(23,&user)
	ChangeEmail("soldiazreal@gmail.com",&user)
	ChangePassword("unaclave",&user)
	
	fmt.Printf("%s %v %s %s\n",user.Name,user.Age, user.Email, user.Password )


}