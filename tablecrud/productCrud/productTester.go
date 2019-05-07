/*=====================================================================*\
	This is the UnitTest for the Backend Product Component


\*=====================================================================*/
package productComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"log"
)

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func TestUpdatePid() {
	log.Printf(">===>productTester.TestUpdatePid()")
	myDB := db.Connect()

	myProduct := &Product{ID: 2}
	aPid, err := myProduct.GetPidByID(myDB, 2)
	log.Printf("   aPid=%d\n", aPid)

	myProduct.UpdatePid(myDB, aPid*2, 10)

	thePid, err := myProduct.GetPidByID(myDB, 2)
	log.Printf("   thePid=%d\n", thePid)

	log.Printf("   err=%v\n", err)

}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func TestGetAllResults() {
	log.Printf(">===>productTester.TestGetAllResults()")
	myDB := db.Connect()
	myProduct := &Product{}
	myProduct.GetProductsFullList(myDB)
}

func TestGetProductsByOwnerID() {
	log.Printf(">===>productTester.TestGetProductsByOwnerID()")
	myDB := db.Connect()
	myProduct := &Product{}
	myProduct.GetProductsByOwnerID(myDB, 2)
}

// func GetProductByName(dbRef *pg.DB, productName string) {
// 	log.Printf(">===>productTester.GetProductByName()")
// 	myProduct := &Product{ProductName: productName}
// 	myProduct.GetByName(dbRef)
// }

// func GetProductById(dbRef *pg.DB, id int) {
// 	log.Printf(">===>productTester.GetProductById()")
// 	myProduct := &Product{ProductID: id}
// 	myProduct.GetByID(dbRef)
// }

// func DeleteProductWithId(dbRef *pg.DB, id int) {
// 	log.Printf(">===>productTester.DeleteProductWithId()")
// 	myProduct := &Product{ProductID: id}
// 	myProduct.Delete(dbRef)
// }

// func DeleteProductWithName(dbRef *pg.DB, productName string) {
// 	log.Printf(">===>productTester.DeleteProductWithName()")
// 	myProduct := &Product{ProductName: productName}
// 	myProduct.Delete(dbRef)
// }
