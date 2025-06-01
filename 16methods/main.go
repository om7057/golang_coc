package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")
	om := User{"Om", "kulkarniom7057@gmail.com",true, 21}
	fmt.Println("Om's details are: ", om)
	fmt.Printf("Details are %+v\n", om)
	fmt.Printf("Name is :%v and Email is %v\n", om.Name, om.Email)
	om.GetStatus()
	om.NewMail()
	fmt.Printf("Name is :%v and Email is %v\n", om.Name, om.Email)
}

type User struct {
	Name string 
	Email string
	Status bool
	Age int 
}


func (u User) GetStatus() {
	fmt.Println("Is user active ? ", u.Status)
}

func (u User) NewMail(){
	u.Email = "tokyo35173@gmail.com"
	fmt.Println("Email of this user is :\n", u.Email)
}