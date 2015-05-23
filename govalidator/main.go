package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"-"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

func main() {

	post := Post{
		Title:    "It My Example",
		Message:  "duck2",
		AuthorIP: "127.2345.54.3",
	}

	returls, err := govalidator.ValidateStruct(post)

	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(returls)
}
