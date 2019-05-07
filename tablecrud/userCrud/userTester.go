package userComp

import (
	"log"

	"github.com/go-pg/pg"
)

// TestGetUsersFullList ...
func TestGetUsersFullList(dbRef *pg.DB) {
	log.Printf("===>userTester.TestGetUsersFullList()")

	myUser := &User{}
	myUser.GetUsersFullList(dbRef)
}

// GetUserByName ...
func GetUserByName(dbRef *pg.DB, userName string) {
	log.Printf("===>userTester.GetUserByName()")

	myUser := &User{UserName: userName}
	myUser.GetByName(dbRef)
}

func GetUserById(dbRef *pg.DB, id int) {
	log.Printf("===>userTester.GetUserById()")

	myUser := &User{ID: id}
	myUser.GetByID(dbRef)
}

func DeleteUserWithId(dbRef *pg.DB, id int) {
	log.Printf("===>userTester.DeleteUserWithId()")

	myUser := &User{ID: id}
	myUser.Delete(dbRef)
}

func DeleteUserWithName(dbRef *pg.DB, userName string) {
	log.Printf("===>userTester.DeleteUserWithName()")

	myUser := &User{UserName: userName}
	myUser.Delete(dbRef)
}
