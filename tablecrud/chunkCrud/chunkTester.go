/*=====================================================================*\
	This is the UnitTest for the Backend Chunk Component


\*=====================================================================*/
package chunkComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"log"
)

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func TestUpdatePid() {
	log.Printf(">===>chunkTester.TestUpdatePid()")
	myDB := db.Connect()

	myChunk := &Chunk{ID: 2}
	aPid, err := myChunk.GetPidByID(myDB, 2)
	log.Printf("   aPid=%d\n", aPid)

	myChunk.UpdatePid(myDB, aPid*2, 10)

	thePid, err := myChunk.GetPidByID(myDB, 2)
	log.Printf("   thePid=%d\n", thePid)

	log.Printf("   err=%v\n", err)

}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
// func TestGetAllResults() {
// 	log.Printf(">===>chunkTester.TestGetAllResults()")
// 	myDB := db.Connect()
// 	myChunk := &Chunk{}
// 	myChunk.GetAllResults(myDB)
// }

func TestGetChunksByOwnerID() {
	log.Printf(">===>chunkTester.TestGetChunksByOwnerID()")
	myDB := db.Connect()
	myChunk := &Chunk{}
	myChunk.GetChunksByOwnerID(myDB, 2)
}

// func GetChunkByName(dbRef *pg.DB, chunkName string) {
// 	log.Printf(">===>chunkTester.GetChunkByName()")
// 	myChunk := &Chunk{ChunkName: chunkName}
// 	myChunk.GetByName(dbRef)
// }

// func GetChunkById(dbRef *pg.DB, id int) {
// 	log.Printf(">===>chunkTester.GetChunkById()")
// 	myChunk := &Chunk{ChunkID: id}
// 	myChunk.GetByID(dbRef)
// }

// func DeleteChunkWithId(dbRef *pg.DB, id int) {
// 	log.Printf(">===>chunkTester.DeleteChunkWithId()")
// 	myChunk := &Chunk{ChunkID: id}
// 	myChunk.Delete(dbRef)
// }

// func DeleteChunkWithName(dbRef *pg.DB, chunkName string) {
// 	log.Printf(">===>chunkTester.DeleteChunkWithName()")
// 	myChunk := &Chunk{ChunkName: chunkName}
// 	myChunk.Delete(dbRef)
// }
