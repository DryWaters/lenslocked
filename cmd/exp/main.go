package main

import (
	"html/template"
	"os"
	"time"
)

type User struct {
	Name     string
	Age      int
	Meta     UserMeta
	Birthday time.Time
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:     "John Doe",
		Age:      111,
		Birthday: time.Now(),
		Meta: UserMeta{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
