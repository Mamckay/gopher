// Package chunkComp ...
/*=====================================================================*\

\*=====================================================================*/
package chunkComp

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// GetPidByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetPidByID(db *pg.DB, chunkID int) (int, error) {
	log.Printf(">===>chunkItem.GetPidByID()")

	getErr := db.Model(gi).
		Where("id = ?0", chunkID).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in chunkItem.GetByOwner(), Reason %v\n", getErr)
		return gi.ChunkPid, getErr
	}
	log.Printf("Select by Owner successful for gi: %v\n", *gi)
	return gi.ChunkPid, nil
}

// UpdatePid ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) UpdatePid(db *pg.DB, pid int, state int) error {
	log.Printf(">===>chunkItem.UpdatePid()")

	_, updateErr := db.Model(gi).
		Set("chunkpid=?0,chunkstate=?1", pid, state).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in chunkItem.Update(), Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Chunk %s updated successfully in table", gi.ChunkName)
	return nil
}

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) Create(db *pg.DB) error {
	log.Printf(">===>chunkItem.Create()")
	log.Printf("   chunkItem.Create():=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in chunkItem.Create(), Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("Chunk %s inserted successfully into table", gi.ChunkName)
	return nil
}

// Update ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) Update(db *pg.DB) error {
	log.Printf(">===>chunkItem.Update()")

	_, updateErr := db.Model(gi).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in chunkItem.Update(), Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Chunk %s updated successfully in table", gi.ChunkName)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) Delete(db *pg.DB) error {
	log.Printf(">===>chunkItem.Delete()")

	_, deleteErr := db.Model(gi).
		Where("chunkname = ?0", gi.ChunkName).
		WhereOr("id = ?0", gi.ID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item in chunkItem.Delete(), Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Chunk %s deleted successfully from table", gi.ChunkName)
	return nil
}

// GetByOwner ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetByOwner(db *pg.DB) (Chunk, error) {
	log.Printf(">===>chunkItem.GetByOwner()")
	//getErr := db.Select(gi)
	getErr := db.Model(gi).
		Where("ownername = ?0", gi.OwnerName).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in chunkItem.GetByOwner(), Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select by Owner successful for gi: %v\n", *gi)
	return *gi, nil
}

// GetByName ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetByName(db *pg.DB) (Chunk, error) {
	log.Printf(">===>chunkItem.GetByName()")
	//getErr := db.Select(gi)
	getErr := db.Model(gi).
		Where("chunkname = ?0", gi.ChunkName).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in chunkItem.GetByName(), Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select by name successful for ID: %v\n", *gi)
	return *gi, nil
}

// GetByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetByID(db *pg.DB) (Chunk, error) {
	log.Printf(">===>chunkItem.GetByID(ChunkID=%d)", gi.ID)

	//getErr := db.Select(gi)
	getErr := db.Model(gi).Where("id = ?0", gi.ID).Select()
	if getErr != nil {
		log.Printf("Error while selecting item, Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select by ID successful in chunkItem.GetById() chunk=%v\n", *gi)
	return *gi, nil
}

/*=====================================================================*\

	Item Functions that are using RELATIONSHIP queries

\*=====================================================================*/

// GetProductsByChunkID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetProductsByChunkID(db *pg.DB, chunkID int) ([]Product, error) {
	log.Printf("===>chunkItem.GetProductsByChunkID()")
	log.Printf("The ChunkID == %d", chunkID)

	gi.ID = chunkID
	log.Printf("gi ChunkID == %d", gi.ID)
	getErr := db.Model(gi).
		Relation("Products").
		Where("id = ?0", chunkID).
		First()

	if getErr != nil {
		log.Printf("Error in chunkItem.GetProductsByChunkID, Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d product found inside chunkItem.GetProductsByChunkID()", len(gi.Products))
	// for i := 0; i < len(gi.Videos); i++ {
	// 	log.Printf("  %v\n", gi.Products[i])
	// }
	return gi.Products, getErr
}

// GetVideosByChunkID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetVideosByChunkID(db *pg.DB, chunkID int) ([]Video, error) {
	log.Printf("===>chunkItem.GetVideosByChunkID()")

	getErr := db.Model(gi).
		Relation("Videos", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC")
			return q, nil
		}).
		Where("id = ?0", chunkID).
		First()

	if getErr != nil {
		log.Printf("Error in chunk.GetVideosByChunkID")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d video found inside chunkItem.GetVideosByChunkID()", len(gi.Videos))
	// for i := 0; i < len(gi.Videos); i++ {
	// 	log.Printf("  %v\n", gi.Videos[i])
	// }
	return gi.Videos, getErr
}

// // GetChunksByProductID ...
// /*---------------------------------------------------------------------*\
// \*---------------------------------------------------------------------*/
// func (gi *Chunk) GetChunksByProductID(db *pg.DB, productID int) ([]Chunk, error) {
// 	log.Printf(">===>chunkItem.GetChunksByProductID()")
// 	return gi, getErr
// }

// // GetChunksByVideoID ...
// /*---------------------------------------------------------------------*\
// \*---------------------------------------------------------------------*/
// func (gi *Chunk) GetChunksByVideoID(db *pg.DB, videoID int) ([]Chunk, error) {
// 	log.Printf(">===>chunkItem.GetChunksByVideoID()")
// 	return &gi, getErr
// }

/*=====================================================================*\

	Functions completed the END 2 END naming conversion

\*=====================================================================*/

// GetChunksFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetChunksFullList(db *pg.DB) ([]Chunk, error) {
	log.Printf(">===>chunkItem.GetChunksFullList()")
	var chunks []Chunk
	getErr := db.Model(&chunks).Column("*").
		Offset(0).
		Order("id asc").
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all chunks in chunkItem.GetChunksFullList(), Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("%d chunks found inside chunkItem.GetChunksFullList()\n", len(chunks))
	// for i := 0; i < len(chunks); i++ {
	// 	log.Printf("  %v\n", chunks[i])
	// }
	return chunks, getErr
}

// GetChunksByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetChunksByOwnerID(db *pg.DB, userid int) ([]Chunk, error) {
	log.Printf(">===>chunkItem.GetChunksByOwnerID(userid=%d)", userid)
	var chunks []Chunk
	getErr := db.Model(&chunks).Column("*").
		Offset(0).
		Order("id asc").
		Where("userid = ?0", userid).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all chunks in chunkItem.GetChunksByOwnerID(), Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("%d chunks found inside chunkItem.GetChunksByOwnerID()\n", len(chunks))
	// for i := 0; i < len(chunks); i++ {
	// 	log.Printf("   %v\n", chunks[i])
	// }
	return chunks, getErr
}

// GetChunksByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Chunk) GetChunksByUserID(db *pg.DB, userid int) ([]Chunk, error) {
	log.Printf("===>ChunkItem.GetChunksByUserID()")

	var user User
	getErr := db.Model(&user).
		Relation("Chunks", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC", userid)
			return q, nil
		}).
		Where("id = ?0", userid).
		First()

	if getErr != nil {
		log.Printf("Error in ChunkItem.GetChunksByUserID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d Chunks found inside ChunkItem.GetChunksByUserID()\n", len(user.Chunks))
	// for i := 0; i < len(user.Chunks); i++ {
	// 	log.Printf("  %v\n", user.Chunks[i])
	// }
	return user.Chunks, getErr
}

/*=====================================================================*\

	UNUSED CHUNK ITEM FUNCTIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
