/*=====================================================================*\
	This file contains multiple sections:
		-ChunkToVideo Table
		-ChunkToVideo Handlers
		-ChunkToVideo Entity

		-ChunkToProduct Table
		-ChunkToProduct Handlers
		-ChunkToProduct Entity
\*=====================================================================*/
package chunkComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

/*=====================================================================*\
	ChunkToVideo Table
\*=====================================================================*/
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type ChunkToVideo struct {
	//tableName struct{} `sql:"relsensoruser"`
	ChunkID int //`sql:"chunk_item_id"` //, type:int references sensor(id)"`
	VideoID int //`sql:"user_id"`       //, type:int references infuser(id)"`
}

type UserToChunk struct {
	//tableName struct{} `sql:"relsensoruser"`
	UserID  int //`sql:"user_id"`
	ChunkID int //`sql:"chunk_id"`
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertChunkToVideoRel(dbRef *pg.DB, chunkID int, videoID int) error {
	log.Printf("===>chunkRelation.InsertChunkToVideoRel(ChunkID=%d,VideoID=%d)\n", chunkID, videoID)

	myDB := db.Connect()
	chk2vid := &ChunkToVideo{ChunkID: chunkID, VideoID: videoID}

	insertErr := myDB.Insert(chk2vid)
	if insertErr != nil {
		log.Printf("Error writing to ChunkToVideo Table in chunkRelation.InsertChunkToVideoRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("ChunkToVideo Relation inserted successfully into ChunkToVideo Table")
	return nil
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateChunkToVideoTable() error {
	log.Printf("===>userTester.CreateChunkToVideoTable()")

	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &ChunkToVideo{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&ChunkToVideo{}, opts)
	if createErr != nil {
		log.Printf("Error creating ChunkToVideo table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("ChunkToVideo Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

/*=====================================================================*\
	ChunkToProduct Table
\*=====================================================================*/
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type ChunkToProduct struct {
	//tableName struct{} `sql:"relsensoruser"`
	ChunkID   int //`sql:"sensor_id, type:int references sensor(id)"`
	ProductID int //`sql:"user_id, type:int references infuser(id)"`
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertChunkToProductRel(dbRef *pg.DB, chunkID int, productID int) error {
	log.Printf("===>chunkRelation.InsertChunkToProductRel(chunkID=%d,productID=%d)\n", chunkID, productID)

	myDB := db.Connect()
	chk2prd := &ChunkToProduct{ChunkID: chunkID, ProductID: productID}

	insertErr := myDB.Insert(chk2prd)
	if insertErr != nil {
		log.Printf("Error writing to ChunkToProduct Table in chunkRelation.InsertChunkToProductRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("ChunkToProduct Relation inserted successfully into ChunkToProduct Table")
	return nil
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateChunkToProductTable() error {
	log.Printf("===>userTester.CreateChunkToProductTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &ChunkToProduct{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&ChunkToProduct{}, opts)
	if createErr != nil {
		log.Printf("Error creating ChunkToProduct table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("ChunkToProduct Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

/*=====================================================================*\

	UNUSED CHUNK RELATIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
