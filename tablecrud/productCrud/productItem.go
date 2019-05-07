// Package productComp ...
/*=====================================================================*\

\*=====================================================================*/
package productComp

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// GetPidByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) GetPidByID(db *pg.DB, productID int) (int, error) {
	log.Printf(">===>productItem.GetPidByID()")

	getErr := db.Model(gi).
		Where("id = ?0", productID).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in productItem.GetByOwner()\n")
		log.Printf("Reason %v\n", getErr)
		return gi.ProductPid, getErr
	}
	log.Printf("Select by Owner successful for gi: %v\n", *gi)
	return gi.ProductPid, nil
}

// UpdatePid ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) UpdatePid(db *pg.DB, pid int, state int) error {
	log.Printf(">===>productItem.Update()")

	_, updateErr := db.Model(gi).
		Set("productpid=?0,productstate=?1", pid, state).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in productItem.Update()\n")
		log.Printf("Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Product %s updated successfully in table", gi.ProductName)
	return nil
}

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) Create(db *pg.DB) error {
	log.Printf(">===>productItem.Create()")
	log.Printf("   productItem.Create():=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in productItem.Create()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("Product %s inserted successfully into table", gi.ProductName)
	return nil
}

// Update ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) Update(db *pg.DB) error {
	log.Printf(">===>productItem.Update()")

	_, updateErr := db.Model(gi).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in productItem.Update()\n")
		log.Printf("Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Product %s updated successfully in table", gi.ProductName)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) Delete(db *pg.DB) error {
	log.Printf(">===>productItem.Delete()")

	_, deleteErr := db.Model(gi).
		Where("productname = ?0", gi.ProductName).
		WhereOr("id = ?0", gi.ID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item in productItem.Delete()\n")
		log.Printf("Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Product %s deleted successfully from table", gi.ProductName)
	return nil
}

// GetByName ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) GetByName(db *pg.DB) (Product, error) {
	log.Printf(">===>productItem.GetByName()")
	getErr := db.Model(gi).
		Where("productname = ?0", gi.ProductName).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in product.GetByName()\n")
		log.Printf("Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select by name successful for ID: %v\n", *gi)
	return *gi, nil
}

// GetByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) GetByID(db *pg.DB) (Product, error) {
	log.Printf(">===>productItem.GetByID(ID=%d)", gi.ID)

	getErr := db.Model(gi).Where("id = ?0", gi.ID).Select()
	if getErr != nil {
		log.Printf("Error while selecting item, Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select by ID successful in productItem.GetById()\n")
	log.Printf("product=%v\n", *gi)
	return *gi, nil
}

/*=====================================================================*\

	Item Functions that are using RELATIONSHIP queries

\*=====================================================================*/

// GetProductsByUserID ...
/*---------------------------------------------------------------------*\
BAD NAME or BAD TYPE
\*---------------------------------------------------------------------*/
func (gi *Product) GetProductsByUserID(db *pg.DB, userId int) ([]Product, error) {
	log.Printf("===>productItem.GetProductsByUserID()")

	var user User
	getErr := db.Model(&user).
		Relation("Products", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC")
			return q, nil
		}).
		Where("id = ?0", userId).
		First()

	if getErr != nil {
		log.Printf("Error in productItem.GetProductsByUserID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d produts found inside productItem.GetProductsByUserID()\n", len(user.Products))
	// for i := 0; i < len(user.Products); i++ {
	// 	log.Printf("   %v\n", user.Products[i])
	// }
	return user.Products, getErr
}

// GetUsersByProductID ...
/*---------------------------------------------------------------------*\
BAD NAME or BAD TYPE
\*---------------------------------------------------------------------*/
func (gi *Product) GetUsersByProductID(db *pg.DB, userId int) ([]User, error) {
	log.Printf("===>productItem.GetUsersByProductID()")

	var prod Product
	getErr := db.Model(&prod).
		Relation("Users", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC")
			return q, nil
		}).
		Where("id = ?0", userId).
		First()

	if getErr != nil {
		// Error Handler
		log.Printf("Error in productItem.GetUsersByProductID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d Users found for product ID \n", len(prod.Users))
	// for i := 0; i < len(prod.Users); i++ {
	// 	log.Printf("   " + prod.Users[i].UserName)
	// }
	return prod.Users, getErr
}

/*=====================================================================*\

	Functions completed the END 2 END naming conversion

\*=====================================================================*/

// GetProductsFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) GetProductsFullList(db *pg.DB) ([]Product, error) {
	log.Printf(">===>productItem.GetProductsFullList()")
	var products []Product
	getErr := db.Model(&products).Column("*").
		Offset(0).
		Order("id asc").
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all products in productItem.GetProductsFullList()\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d products found inside productItem.GetProductsFullList()\n", len(products))
	// for i := 0; i < len(products); i++ {
	// 	log.Printf("   %v\n", products[i])
	// }
	return products, getErr
}

// GetProductsByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) GetProductsByOwnerID(db *pg.DB, userId int) ([]Product, error) {
	log.Printf(">===>productItem.GetProductsByOwnerID(userId=%d)", userId)
	var products []Product
	getErr := db.Model(&products).Column("*").
		Offset(0).
		Order("id asc").
		Where("userid = ?0", userId).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all products in productItem.GetProductsByOwnerID()\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d products found inside productItem.GetProductsByOwnerID()\n", len(products))
	// for i := 0; i < len(products); i++ {
	// 	log.Printf("  %v\n", products[i])
	// }
	return products, getErr
}

// GetProductsByOwnerName ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Product) GetProductsByOwnerName(db *pg.DB) (Product, error) {
	log.Printf(">===>productItem.GetProductsByOwnerName()")
	getErr := db.Model(gi).
		Where("ownername = ?0", gi.OwnerName).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in productItem.GetProductsByOwnerName()\n")
		log.Printf("Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select by Owner successful for gi: %v\n", *gi)
	return *gi, nil
}

/*=====================================================================*\

	UNUSED PRODUCT ITEM FUNCTIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
