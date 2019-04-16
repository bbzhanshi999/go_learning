package main

import (
	"fmt"
	"reflect"
)

type User struct {
	isExist  bool   "is an valid user"
	username string "the name of the user"
	gender   int    "the sexual of user"
}

func main() {
	user := User{false, "aaaa", 1}
	for i := 0; i < 3; i++ {
		refTag(user, i)
	}
}

func refTag(user User, ix int) {
	typ := reflect.TypeOf(user)
	ixField := typ.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
