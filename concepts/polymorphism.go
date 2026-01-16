package main

import "log"

type Greeter interface {
	Greet() string
}
type User struct {
	Name string
}

type Admin struct {
	Name string
	Designation string
}

     
func (u User) Greet() string {     //polymorphism
	return "Hello" + u.Name
}

func (ad Admin) Greet() string {
	return "Welcome" + ad.Name + "!" + "Your designation is " + ad.Designation  //polymorphism
}

func SayHello(g Greeter ) {
	log.Print(g.Greet())
}

func main1() {
	u := User{Name:"akangkha"}
	SayHello(u)
	log.Print("----------admin------")
	ad:= Admin{Name:"akangkha", Designation:"system designer"}
	SayHello(ad)
}