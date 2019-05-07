/*=====================================================================*\
	All the functions inside this file are handler functions.
	This means they are passed a Request and they interact with
	the Product database table and then generate a Response.
\*=====================================================================*/
package productComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"
)

// CreateProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>productHandlers.CreateProduct()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)

	var product Product
	json.Unmarshal([]byte(har), &product)
	fmt.Printf("product=%v\n", product)

	myDB := db.Connect()

	log.Printf("   productHandlers.Create():=%v\n", product)

	sqlErr := product.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created product productHandlers.CreateProduct()\n")
		log.Printf(" product=%v\n", product)
		json.NewEncoder(w).Encode(product)
		return
	}
	log.Printf("Error creating product in productHandlers.CreateProduct()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

// UpdateProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.UpdateProduct()")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)

	pProduct := params["Product"]
	log.Printf("   Product=%v\n", pProduct)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var product Product
	json.Unmarshal([]byte(har), &product)
	fmt.Printf("%v", product)

	sqlErr := product.Update(myDB)
	if sqlErr == nil {
		log.Printf("Successful update in productHandlers.UpdateProduct() product=%v\n", product)
		json.NewEncoder(w).Encode(product)
		return
	}
	log.Printf("Error updateing product in productHandlers.UpdateProduct()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

//DeleteProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>productHandlers.DeleteProduct()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProduct := &Product{}
	paramInt, paramErr := strconv.Atoi(params["id"])

	if paramErr == nil {
		myProduct.ID = paramInt
		sqlErr := myProduct.Delete(myDB)
		if sqlErr == nil {
			log.Printf("Successful delete in productHandlers.DeleteProducts() product=%v\n", myProduct)
			json.NewEncoder(w).Encode(myProduct)
			return
		}
		log.Printf("Error with Delete() in productHandlers.DeleteProducts()\n")
		log.Printf(", Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in productHandlers.DeleteProducts()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

// GetProduct ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.GetProduct()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProduct := &Product{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		myProduct.ID = paramInt
		product, sqlErr := myProduct.GetByID(myDB)
		if sqlErr == nil {
			log.Printf("Successful select in productHandlers.GetProducts() product=%v\n", product)
			json.NewEncoder(w).Encode(product)
			return
		}
		log.Printf("Error with GetByID() in productHandlers.GetProduct()\n")
		log.Printf(", Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in productHandlers.GetProduct()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetProductsFullList_Other(dbRef *pg.DB) {
	log.Printf(">===>productTester.GetProductsFullList_Other()")

	myProduct := &Product{}
	myProduct.GetProductsFullList(dbRef)
}

/*=====================================================================*\

	Handlers Functions completed the END 2 END naming conversion

\*=====================================================================*/

// GetProductsFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetProductsFullList(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.GetProductsFullList()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	myProduct := &Product{}
	products, sqlErr := myProduct.GetProductsFullList(myDB)
	log.Printf("%d products found inside productHandlers.GetProductsFullList()\n", len(products))
	// for i := 0; i < len(products); i++ {
	// 	log.Printf("   %v\n", products[i])
	// }

	if sqlErr == nil {
		log.Printf("Successful select of all products in productHandlers.GetProductsFullList()")
		json.NewEncoder(w).Encode(products)
		return
	}
	log.Printf("Error with GetAllResults() in productHandlers.GetProductsFullList()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

// GetProductsByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetProductsByOwnerID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.GetProductsByOwnerID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProduct := &Product{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   UserID=%d\n", paramInt)

	if paramErr == nil {
		myProduct.ID = paramInt
		products, sqlErr := myProduct.GetProductsByOwnerID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d products found inside productHandlers.GetProductsByOwnerID()\n", len(products))
			// for i := 0; i < len(products); i++ {
			// 	log.Printf("  ProductName=%s,UserID=%d", products[i].ProductName, products[i].UserID)
			// }
			log.Printf("Successful select of all products in productHandlers.GetProductsByOwnerID()")
			json.NewEncoder(w).Encode(products)
			return
		}
		log.Printf("Error with GetResultsByUserID() in productHandlers.GetProductsByOwnerID()\n")
		log.Printf("Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in productHandlers.GetProductsByOwnerID()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

// GetProductsByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetProductsByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.GetProductsByUserID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProduct := &Product{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   UserID=%d\n", paramInt)

	if paramErr == nil {
		myProduct.ID = paramInt
		products, sqlErr := myProduct.GetProductsByUserID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d products found inside productHandlers.GetProductsByUserID()\n", len(products))
			// for i := 0; i < len(products); i++ {
			// 	log.Printf("  ProductName=%s,UserID=%d", products[i].ProductName, products[i].UserID)
			// }
			log.Printf("Successful select of all products in productHandlers.GetProductsByUserID()")
			json.NewEncoder(w).Encode(products)
			return
		}
		log.Printf("Error with GetResultsByUserID() in productHandlers.GetProductsByUserID()\n")
		log.Printf("Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in productHandlers.GetProductsByUserID()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

// GetUsersByProductID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetUsersByProductID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>productHandlers.GetUsersByProductID()")

	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProduct := &Product{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("  paramInt == %d", paramInt)
	if paramErr == nil {
		myProduct.ID = paramInt
		userList, sqlErr := myProduct.GetUsersByProductID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d Users found in productHandlers.GetUsersByProductID()\n", len(userList))
			json.NewEncoder(w).Encode(userList)
			return
		}
		log.Printf("Error with database select in productHandlers.GetUsersByProductID(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in productHandlers.GetUsersByProductID(), Reason:  %v\n", paramErr)
}

// GetUsersByProductID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetNotUsersByProductID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetUsersByProductID()")

}
