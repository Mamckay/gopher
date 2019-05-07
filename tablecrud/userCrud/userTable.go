/*=====================================================================*\
The user model component defines the User table and provides
the create, read, update, and delete functions.
\*=====================================================================*/

package userComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"

	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// User ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type User struct {
	tableName struct{} `sql:"entusers"`
	ID        int      `sql:"id,pk" json:"UserID,omitempty"`
	UserName  string   `sql:"username,unique" json:"UserName,omitempty"`
	Password  string   `sql:"password" json:"Password,omitempty"`
	FirstName string   `sql:"firstname" json:"FirstName,omitempty"`
	LastName  string   `sql:"lastname" json:"LastName,omitempty"`
	RoleName  string   `sql:"rolename" json:"RoleName,omitempty"`
	JwtToken  string   `sql:"jwttoken" json:"JwtToken,omitempty"`
	// Sensor []*sen.Sensor `pg:",many2many:sensor_item"`
}

// CreateUserWithName ...
func InsertUserEnt(dbRef *pg.DB,
	userName string,
	password string,
	firstName string,
	lastName string,
	roleName string,
	jwttoken string,
) {
	log.Printf("===>userTester.CreateUser()")
	myUser := User{
		UserName:  userName,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		RoleName:  roleName,
		JwtToken:  jwttoken,
	}
	myUser.Create(dbRef)
}

// CreateUserTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateUserTable() error {
	log.Printf("===>userTester.CreateUserTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: true,
	}
	orm.DropTable(myDB, &User{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&User{}, opts)
	if createErr != nil {
		log.Printf("Error creating User table products, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("User Table created successfully. Only if necessary.\n")
	//LoadUserTable(myDB)
	return nil
}

// LoadVideoTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadUserTable(dbRef *pg.DB) {
	log.Println("===>userTable.LoadUserTable()")
	var testCase = 1
	if testCase == 1 {
		InsertUserEnt(dbRef, "AdminUser", "123456", "John", "Doe", "Admin", "Token")
		InsertUserEnt(dbRef, "OperateUser", "123456", "Jean", "Smith", "Operator", "Token")
		InsertUserEnt(dbRef, "ReviewUser", "123456", "Mike", "Johnson", "Viewer", "Token")
		InsertUserEnt(dbRef, "ProdUser", "123456", "Kim", "Hann", "Creator", "Token")
	} else if testCase == 2 {
		TestGetUsersFullList(dbRef)
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}
