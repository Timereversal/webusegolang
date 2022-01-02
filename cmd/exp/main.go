package main

import (
	"errors"
	"html/template"
	"os"
)

type User struct {
	Name    string
	Age     int
	Height  float32
	Weight  float32
	Friends []string
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name:    "Berzoc",
		Age:     36,
		Height:  1.80,
		Weight:  70,
		Friends: []string{"enver", "barnie", "Roger"},
	}
	t.Execute(os.Stdout, user)
}

// type Error struct {
//     Path string
//     User string
// }

// func (e *Error) Is(target error) bool {
//     t, ok := target.(*Error)
//     if !ok {
//         return false
//     }
//     return (e.Path == t.Path || t.Path == "") &&
//            (e.User == t.User || t.User == "")
// }

// if errors.Is(err, &Error{User: "someuser"}) {
//     // err's User field is "someuser".
// }

func Connect() error {
	return errors.New("Connection failed")
}
