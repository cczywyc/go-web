package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// u is struct, it's data can not be changed
	u := User{
		Name: "Tom",
		Age:  10,
	}
	u.ChangeName("Tom changed")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)

	// up is pointer, it's data cen be changed
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	up.ChangeName("Jerry changed")
	up.ChangeAge(120)
	fmt.Printf("%v \n", up)
}

func (u User) ChangeName(newName string) {
	u.Name = newName
}

func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}
