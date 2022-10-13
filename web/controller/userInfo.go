package controller

import "github.com/hydrogen1999/grooo-network/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName string
	Password  string
	IsAdmin   string
}

var users []User

func init() {

	admin1 := User{LoginName: "admin1", Password: "12345", IsAdmin: "T"}
	admin2 := User{LoginName: "admin2", Password: "12345", IsAdmin: "T"}
	admin3 := User{LoginName: "admin3", Password: "12345", IsAdmin: "T"}
	admin4 := User{LoginName: "admin4", Password: "12345", IsAdmin: "T"}
	admin5 := User{LoginName: "admin5", Password: "12345", IsAdmin: "T"}
	admin6 := User{LoginName: "admin6", Password: "12345", IsAdmin: "T"}
	admin7 := User{LoginName: "admin7", Password: "12345", IsAdmin: "T"}
	admin8 := User{LoginName: "admin8", Password: "12345", IsAdmin: "T"}
	admin9 := User{LoginName: "admin9", Password: "12345", IsAdmin: "T"}
	admin10 := User{LoginName: "admin10", Password: "12345", IsAdmin: "T"}
	wenbo := User{LoginName: "thinhnc", Password: "123456", IsAdmin: "T"}

	users = append(users, admin1)
	users = append(users, admin2)
	users = append(users, admin3)
	users = append(users, admin4)
	users = append(users, admin5)
	users = append(users, admin6)
	users = append(users, admin7)
	users = append(users, admin8)
	users = append(users, admin9)
	users = append(users, admin10)
	users = append(users, wenbo)

}
