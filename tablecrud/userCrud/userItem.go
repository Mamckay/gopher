// Package userComp ...
/*=====================================================================*\

\*=====================================================================*/
package userComp

import (
	"log"

	"github.com/go-pg/pg"
)

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *User) Create(db *pg.DB) error {
	log.Printf("===>userItem.Create()")
	log.Printf("   user.Create():=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in userItem.Create()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("User %s inserted successfully into table", gi.UserName)
	return nil
}

// Update ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *User) Update(db *pg.DB) error {
	log.Printf("===>userItem.Update()")

	_, updateErr := db.Model(gi).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in userItem.Update()\n")
		log.Printf("Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Product %s updated successfully in table", gi.UserName)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *User) Delete(db *pg.DB) error {
	log.Printf("===>userItem.Delete()")

	_, deleteErr := db.Model(gi).
		Where("username = ?0", gi.UserName).
		WhereOr("id = ?0", gi.ID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item in user.Delete()\n")
		log.Printf("Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Product %s deleted successfully from table", gi.UserName)
	return nil
}

// GetByName ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *User) GetByName(db *pg.DB) (User, error) {
	log.Printf("===>userItem.GetByName()")
	//getErr := db.Select(gi)
	getErr := db.Model(gi).
		Where("username = ?0", gi.UserName).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in userItem.GetByName()\n")
		log.Printf("Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select successful for ID: %v\n", *gi)
	return *gi, nil
}

// GetByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *User) GetByID(db *pg.DB) (User, error) {
	log.Printf("===>userItem.GetByID(UserID=%d)", gi.ID)

	//getErr := db.Select(gi)
	getErr := db.Model(gi).Where("userid = ?0", gi.ID).Select()
	if getErr != nil {
		log.Printf("Error while selecting item\n")
		log.Printf("Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select successful in userItem.GetById() user=%v\n", *gi)
	return *gi, nil
}

/*=====================================================================*\

	Item Functions that are using RELATIONSHIP queries

\*=====================================================================*/

/*=====================================================================*\

	Functions completed the END 2 END naming conversion

\*=====================================================================*/
// GetUsersFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *User) GetUsersFullList(db *pg.DB) ([]User, error) {
	log.Printf("===>userItem.GetUsersFullList()")
	var users []User
	getErr := db.Model(&users).Column("*").
		Offset(0).
		Order("username asc").
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all users in userItem.GetUsersFullList()\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d users found inside userItem.GetUsersFullList()\n", len(users))
	// for i := 0; i < len(users); i++ {
	// 	log.Printf("   username=" + users[i].UserName)
	// }
	return users, getErr
}

/*=====================================================================*\

	UNUSED USER ITEM FUNCTIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
