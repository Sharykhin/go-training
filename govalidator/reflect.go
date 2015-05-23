package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `gender:"male" age:"18"`
}

func (person *Person) GetName() string {
	return person.Name
}

func main() {

	s := Person{Name: "Katy"}
	fmt.Println(s.GetName())
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("age"), field.Tag.Get("gender"))
	value := reflect.ValueOf(s)
	fmt.Println(value.Kind())
	fmt.Println(st.Kind())

}
