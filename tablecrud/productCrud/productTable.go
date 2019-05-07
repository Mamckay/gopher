/*=====================================================================*\
	This


\*=====================================================================*/
package productComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Product ...
/*---------------------------------------------------------------------*\
	// ID           int      `sql:"type:int references userinfo(userid)"`
\*---------------------------------------------------------------------*/
type Product struct {
	tableName    struct{} `sql:"entproducts"`
	ID           int      `sql:"id,pk" json:"ProductID,omitempty"`
	ProductName  string   `sql:"productname"`
	ProductType  string   `sql:"producttype"`
	UserID       int      `sql:"userid"`
	SensorID     int      `sql:"sensorid"`
	OwnerName    string   `sql:"ownername"`
	FilePath     string   `sql:"filepath"`
	ProductState int      `sql:"productstate"`
	ProductPid   int      `sql:"productpid"`
	Users        []User   `pg:",many2many:product_to_users"`
}

// User ...
/*---------------------------------------------------------------------*\
	THIS IS A REFERENCE TABLE
\*---------------------------------------------------------------------*/
type User struct {
	tableName struct{}  `sql:"entusers"`
	ID        int       `sql:"id,pk" json:"UserID"`
	UserName  string    `sql:"username,unique"`
	Password  string    `sql:"password"`
	FirstName string    `sql:"firstname"`
	LastName  string    `sql:"lastname"`
	RoleName  string    `sql:"rolename"`
	JwtToken  string    `sql:"jwttoken"`
	Products  []Product `pg:",many2many:user_to_products"`
}

// CreateProductTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateProductTable() error {
	log.Printf(">===>productTable.CreateProductTable()")
	myDB := db.Connect()

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := myDB.CreateTable(&Product{}, opts)
	if createErr != nil {
		log.Printf("Error creating Gaol table products, Reason:%v\n", createErr)
		return createErr
	}
	log.Printf("Product Info Table created successfully. Only if necessary.\n")
	//LoadProductTable(myDB)
	return nil
}

// LoadProductTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadProductTable(dbRef *pg.DB) {
	log.Println(">===>productTable.LoadProductTable()")

	var testCase = 1
	MyProduct := &Product{}
	if testCase == 1 {
		InsertProductEnt(dbRef,
			//2,
			"Product One",
			"S3VC",
			2,
			2,
			"SomeOwner",
			"assets/Products/dir1/",
			1,
			1234,
		)
	} else if testCase == 2 {
		//		GetAllProducts(dbRef)
		MyProduct.GetProductsFullList(dbRef)
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}

// CreateProductWithName ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertProductEnt(dbRef *pg.DB,
	//productID int,
	productName string,
	productType string,
	userID int,
	sensorID int,
	ownername string,
	filepath string,
	productstate int,
	productpid int,
) {
	log.Println(">===>productTable.InsertProductEnt()")
	myProduct := Product{
		//ID:           productID,
		ProductName:  productName,
		ProductType:  productType,
		UserID:       userID,
		SensorID:     sensorID,
		OwnerName:    ownername,
		FilePath:     filepath,
		ProductState: productstate,
		ProductPid:   productpid,
	}
	myProduct.Create(dbRef)
}
