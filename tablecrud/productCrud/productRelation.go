/*=====================================================================*\
	This file contains multiple sections:
		-ProductToUser Table
		-ChunkToVideo Handlers
		-ChunkToVideo Entity

		-ChunkToProduct Table
		-ChunkToProduct Handlers
		-ChunkToProduct Entity
\*=====================================================================*/
package productComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	//"github.com/rs/cors"
)

// UserToProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type UserToProduct struct {
	tableName struct{} `sql:"user_to_products"`
	UserID    int      //`sql:"sensor_id, type:int references sensor(id)"`
	ProductID int      //`sql:"user_id, type:int references infuser(id)"`
}

// CreateUserToProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateUserToProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>productRelation.CreateUserToProduct()")
	w.Header().Add("Content-Type", "application/json")

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aUserToProduct UserToProduct
	json.Unmarshal([]byte(har), &aUserToProduct)
	fmt.Printf("unmarshalled aUserToProduct=%v", aUserToProduct)

	myDB := db.Connect()
	sqlErr := aUserToProduct.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created product productRelation.CreateUserToProduct()\n")
		log.Printf(" product=%v\n", aUserToProduct)
		json.NewEncoder(w).Encode(aUserToProduct)
		return
	}
	log.Printf("Error creating product in productRelation.CreateUserToProduct()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

// InsertUserToProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertUserToProduct(dbRef *pg.DB, userID int, productID int) error {
	log.Printf("===>productRelation.InsertUserToProduc(userID=%d,productID=%d)\n", userID, productID)
	myDB := db.Connect()
	var aUserToProduct UserToProduct
	aUserToProduct.UserID = userID
	aUserToProduct.ProductID = productID
	sqlErr := aUserToProduct.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created product productRelation.InsertUserToProduct()\n")
		log.Printf(" product=%v\n", aUserToProduct)
		return nil
	}
	log.Printf("Error creating product in productRelation.InsertUserToProduct()\n")
	log.Printf("Reason:%v\n", sqlErr)
	return sqlErr
}

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *UserToProduct) Create(db *pg.DB) error {
	log.Printf(">===>productRelation.Create()")
	log.Printf("   productRelation.Create():=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in UserToProduct.Create()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("UserToProduct %d to %d inserted successfully into table", gi.UserID, gi.ProductID)
	return nil
}

// DeleteUserToProductRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteUserToProductRel(dbRef *pg.DB, userID int, productID int) error {
	log.Printf("===>productRelation.DeleteUserToProductRel(userID=%d,productID=%d)\n", userID, productID)
	myDB := db.Connect()
	usr2prd := &UserToProduct{UserID: userID, ProductID: productID}

	insertErr := myDB.Delete(usr2prd)
	if insertErr != nil {
		log.Printf("Error writing to UserToProduct Table in prodUserRelation.DeleteUserToProductRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("UserToProduct Relation Deleted successfully into UserToProduct Table")
	return nil
}

// CreateUserToProductTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateUserToProductTable() error {
	log.Printf("===>userTester.CreateUserToProductTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &UserToProduct{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&UserToProduct{}, opts)
	if createErr != nil {
		log.Printf("Error creating UserToProduct table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("UserToProduct Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

/*=====================================================================*\

	UNUSED PRODUCT RELATIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/

// HandleUserToProductRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func HandleUserToProductRel(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>productRelation.InsertUserToProductRel()")
	w.Header().Add("Content-Type", "application/json")

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aUserToProduct UserToProduct
	json.Unmarshal([]byte(har), &aUserToProduct)
	fmt.Printf("unmarshalled aUserToProduct=%v", aUserToProduct)

	UserID := aUserToProduct.UserID
	ProductID := aUserToProduct.ProductID

	myDB := db.Connect()
	sqlErr := InsertUserToProductRel(myDB, UserID, ProductID)
	if sqlErr == nil {
		log.Printf("Successful Insert of UserID=%d to ProductID=%d\n", UserID, ProductID)
		json.NewEncoder(w).Encode("Success")
		return
	}
	log.Printf("Error Insert nsert of UserID=%d to ProductID=%d\n", UserID, ProductID)
	log.Printf("Reason:%v\n", sqlErr)
}

// InsertUserToProductRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertUserToProductRel(dbRef *pg.DB, userID int, productID int) error {
	log.Printf("===>productRelation.InsertUserToProductRel(userID=%d,productID=%d)\n", userID, productID)
	myDB := db.Connect()
	usr2prd := &UserToProduct{UserID: userID, ProductID: productID}

	insertErr := myDB.Insert(usr2prd)
	if insertErr != nil {
		log.Printf("Error writing to UserToProduct Table in prodUserRelation.InsertUserToProductRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("UserToProduct Relation inserted successfully into UserToProduct Table")
	return nil
}

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *ProductToUser) Create(db *pg.DB) error {
	log.Printf("===>prodUserModel.Create()")
	log.Printf("   prodUserModel.Create():ProductToUser=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in prodUserModel.Create(), Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("ProductToUser %s inserted successfully into table", gi.ProductID)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *ProductToUser) Delete(db *pg.DB) error {
	log.Printf("===>productModel.Delete()")

	_, deleteErr := db.Model(gi).
		Where("sensorid = ?0", gi.ProductID).
		WhereOr("userid = ?0", gi.UserID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting ProductToUser in productModel.Delete(), Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("ProductToUser %s deleted successfully from table", gi.ProductID)
	return nil
}

// ProductToUser_NOTUSED ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type ProductToUser struct {
	ProductID int //`sql:"sensor_id, type:int references sensor(id)"`
	UserID    int //`sql:"user_id, type:int references infuser(id)"`
}

// InsertProductToUserRel_NOTUSED ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertProductToUserRel(dbRef *pg.DB, productID int, userID int) error {
	log.Printf("===>productRelation.InsertProductToUserRel_NOTUSED(productID=%d,userID=%d)\n", productID, userID)
	myDB := db.Connect()
	prd2usr := &ProductToUser{ProductID: productID, UserID: userID}

	insertErr := myDB.Insert(prd2usr)
	if insertErr != nil {
		log.Printf("Error writing to ProductToUser Table in prodUserRelation.InsertProductToUserRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("ProductToUser Relation inserted successfully into ProductToUser Table")
	return nil
}

// CreateProductToUserTable_NOTUSED ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateProductToUserTable() error {
	log.Printf("===>userTester.CreateProductToUserTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &ProductToUser{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&ProductToUser{}, opts)
	if createErr != nil {
		log.Printf("Error creating ProductToUser table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("ProductToUser Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

// CreateSenUserTable_NOTUSED ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSenUserTable_NOTUSED() error {
	log.Printf("===>sensorTester.CreateSenUserTable()")
	myDB := db.Connect()

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := myDB.CreateTable(&ProductToUser{}, opts)
	if createErr != nil {
		log.Printf("Error creating table SenUser, Reason:%v\n", createErr)
		return createErr
	}
	log.Printf("SenUser table created successfully. Only if necessary.\n")
	LoadSenUserTable_NOTUSED(myDB)
	return nil
}

// LoadSenUserTable_NOTUSED ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadSenUserTable_NOTUSED(dbRef *pg.DB) {
	log.Println("===>sensorRelation.LoadSenUserTable()")
	var testCase = 0
	if testCase == 1 {
		CreateSenUserTest_NOTUSED(dbRef, 1, 1)
	} else if testCase == 2 {
		//GetAllSensors(dbRef)
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}

// CreateSenUserTest_NOTUSED ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSenUserTest_NOTUSED(dbRef *pg.DB,
	userID int,
	productID int,
) {
	log.Printf("===>sensorTester.CreateSenUserTest()")
	mySenUser := ProductToUser{
		UserID:    userID,
		ProductID: productID,
	}
	mySenUser.Create(dbRef)
}
