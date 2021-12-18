package main

import (
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
