/*=====================================================================*\


\*=====================================================================*/
package productComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"testing"
)

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}
func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

type A struct{ x, y int }

func (*A) Method() {}

// TestGetAllResult ...
func TestGetAllResults2(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/productCrud.TestGetAllResults()")
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	myDB := db.Connect()
	myUser := &Product{}
	myUser.GetAllResults(myDB)
}

// TestGetAllResult ...
func TestGetResultsProductsByUser(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	//log.Printf("RUNNING: ===>tablecrud/productCrud.TestGetAllResults()")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Product{}
	myUser.GetResultsByUserID(myDB, 2)
}

// TestGetUserById ...
func TestGetProductById(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/productCrud.TestGetProductById()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Product{ProductID: 1}
	myUser.GetByID(myDB)
}

// TestGetUserByName ...
func TestGetProductByName(t *testing.T) {
	log.Printf("-------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/productCrud.TestGetProductByName()")
	log.Printf("-------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Product{ProductName: "ProductID1"}
	myUser.GetByName(myDB)
}

func TestDeleteProductWithId(t *testing.T) {
	log.Printf("----------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/productCrud.TestDeleteProductWithId()")
	log.Printf("----------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Product{ProductID: 3}
	myUser.Delete(myDB)
}

func TestDeleteProductWithName(t *testing.T) {
	log.Printf("------------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/productCrud.TestDeleteProductWithName()")
	log.Printf("------------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Product{ProductName: "ProductID2"}
	myUser.Delete(myDB)
}

/*=====================================================================*\
	Functions below this point have been prepared for integrations

\*=====================================================================*/

func TestGetProductsByOwnerID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myProd := &Product{}
	prodts, errTest := myProd.GetProductsByOwnerID(myDB, 7)
	if errTest != nil {
		log.Printf("   Test Failed in GetProductsByOwnerID: %v", errTest)
		log.Printf("   %v", prodts)
		return
	}
	fmt.Println("User ", prodts[0].ProductID)
	for i := 0; i < len(prodts); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", prodts[i])
	}
	// for i := 0; i < len(prodts); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensor[i].SensorID,
	// 		prodts[i].UserID,
	// 		prodts[i].SensorName,
	// 		prodts[i].Location,
	// 		prodts[i].InputURL,
	// 		prodts[i].SecurityLevel)
	// }
}

func TestGetProductsByUserID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myProd := &Product{}
	usr, errTest := myProd.GetProductsByUserID(myDB, 7)
	if errTest != nil {
		log.Printf("   Test Failed in GetProductsByUserID: %v", errTest)
		log.Printf("   %v", usr)
		return
	}
	fmt.Println("User ", usr.ID)
	for i := 0; i < len(usr.Products); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", usr.Products[i])
	}
	// for i := 0; i < len(sensor); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensor[i].SensorID,
	// 		sensor[i].UserID,
	// 		sensor[i].SensorName,
	// 		sensor[i].Location,
	// 		sensor[i].InputURL,
	// 		sensor[i].SecurityLevel)
	// }
}

func TestGetUsersByProductID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Product{}
	product, errTest := myUser.GetUsersByProductID(myDB, 2)
	if errTest != nil {
		log.Printf("   Test Failed in GetUsersByProductID: %v", errTest)
		log.Printf("   %v", product)
		return
	}
	fmt.Println("Product ", product.ProductID)
	for i := 0; i < len(product.Users); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", product.Users[i])
	}
	// for i := 0; i < len(sensor); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensor[i].SensorID,
	// 		sensor[i].UserID,
	// 		sensor[i].SensorName,
	// 		sensor[i].Location,
	// 		sensor[i].InputURL,
	// 		sensor[i].SecurityLevel)
	// }
}
