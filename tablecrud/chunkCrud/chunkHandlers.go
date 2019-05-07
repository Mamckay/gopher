/*=====================================================================*\
	All the functions inside this file are handler functions.
	This means they are passed a Request and they interact with
	the Chunk database table and then generate a Response.
\*=====================================================================*/
package chunkComp

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

// CreateChunk ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateChunk(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>chunkHandlers.CreateChunk()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)

	var chunk Chunk
	json.Unmarshal([]byte(har), &chunk)
	fmt.Printf("%v", chunk)

	myDB := db.Connect()

	log.Printf("   chunkHandlers.Create():=%v\n", chunk)

	sqlErr := chunk.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created chunk chunkHandlers.CreateChunk(): chunk=%v\n", chunk)
		json.NewEncoder(w).Encode(chunk)
		return
	}
	log.Printf("Error creating chunk in chunkHandlers.CreateChunk(), Reason:%v\n", sqlErr)
}

// UpdateChunk ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func UpdateChunk(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers.UpdateChunk()")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)

	pChunk := params["Chunk"]
	log.Printf("   Chunk=%v\n", pChunk)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var chunk Chunk
	json.Unmarshal([]byte(har), &chunk)
	fmt.Printf("%v", chunk)

	sqlErr := chunk.Update(myDB)
	if sqlErr == nil {
		log.Printf("Successful update in chunkHandlers.UpdateChunk() chunk=%v\n", chunk)
		json.NewEncoder(w).Encode(chunk)
		return
	}
	log.Printf("Error updateing chunk in chunkHandlers.UpdateChunk(), Reason:%v\n", sqlErr)
}

//DeleteChunk ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteChunk(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>chunkHandlers.DeleteChunk()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myChunk := &Chunk{}
	paramInt, paramErr := strconv.Atoi(params["id"])

	if paramErr == nil {
		myChunk.ID = paramInt
		sqlErr := myChunk.Delete(myDB)
		if sqlErr == nil {
			log.Printf("Successful delete in service.DeleteChunks() chunk=%v\n", myChunk)
			json.NewEncoder(w).Encode(myChunk)
			return
		}
		log.Printf("Error with database delete in service.DeleteChunks(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in chunkHandlers.DeleteChunks(), Reason:  %v\n", paramErr)
}

// GetChunk ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetChunk(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers.GetChunk()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myChunk := &Chunk{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		myChunk.ID = paramInt
		chunk, sqlErr := myChunk.GetByID(myDB)
		if sqlErr == nil {
			log.Printf("Successful select in service.GetChunks() chunk=%v\n", chunk)
			json.NewEncoder(w).Encode(chunk)
			return
		}
		log.Printf("Error with database select in service.GetChunk(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in chunkHandlers.GetChunk(), Reason:  %v\n", paramErr)
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetAllChunks(dbRef *pg.DB) {
	log.Printf(">===>chunkHandlers.GetAllChunks()")

	myChunk := &Chunk{}
	myChunk.GetChunksFullList(dbRef)
}

// VideosByChunkID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func VideosByChunkID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers.VideosByChunkID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProd := &Chunk{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   ChunkID=%d\n", paramInt)

	if paramErr == nil {
		videoList, sqlErr := myProd.GetVideosByChunkID(myDB, paramInt)

		if sqlErr == nil {
			log.Printf("%d videos found inside chunkHandlers.ProductsByChunkID()\n", len(videoList))
			// for i := 0; i < len(videoList); i++ {
			// 	log.Printf("  %v\n", videoList)
			// }
			log.Printf("Successful select of all chunks in service.GetChunksByOwner()")
			json.NewEncoder(w).Encode(videoList)
			return
		}
		log.Printf("Error with chunk.GetAllResults() in chunkHandlers.GetChunksByOwner(), Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in chunkHandlers.GetChunksByOwner(), Reason:  %v\n", paramErr)
}

// ProductsByChunkID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func ProductsByChunkID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers.ProductsByChunkID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProd := &Chunk{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   ChunkID=%d\n", paramInt)

	if paramErr == nil {
		productList, sqlErr := myProd.GetProductsByChunkID(myDB, paramInt)

		if sqlErr == nil {
			log.Printf("%d products found inside chunkHandlers.ProductsByChunkID()\n", len(productList))
			// for i := 0; i < len(productList); i++ {
			// 	log.Printf("  %v\n", productList)
			// }
			log.Printf("Successful select of all chunks in service.GetChunksByOwner()")
			json.NewEncoder(w).Encode(productList)
			return
		}
		log.Printf("Error with chunk.GetAllResults() in chunkHandlers.GetChunksByOwner(), Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in chunkHandlers.GetChunksByOwner(), Reason:  %v\n", paramErr)
}

/*=====================================================================*\

	Handlers Functions completed the END 2 END naming conversion

\*=====================================================================*/

// GetChunksFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetChunksFullList(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers.GetChunksFullList()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	myChunk := &Chunk{}
	chunks, sqlErr := myChunk.GetChunksFullList(myDB)
	log.Printf("%d chunks found inside chunkHandlers.GetChunksFullList()\n", len(chunks))
	// for i := 0; i < len(chunks); i++ {
	// 	log.Printf("   %v\n", chunks[i])
	// }

	if sqlErr == nil {
		log.Printf("Successful select of all chunks in service.GetChunksFullList()")
		json.NewEncoder(w).Encode(chunks)
		return
	}
	log.Printf("Error with database select in chunkHandlers.GetChunksFullList(), Reason:%v\n", sqlErr)
}

// GetChunksByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetChunksByOwnerID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers..GetChunksByOwnerID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myChunk := &Chunk{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   UserID=%d\n", paramInt)

	if paramErr == nil {
		myChunk.ID = paramInt
		chunks, sqlErr := myChunk.GetChunksByOwnerID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d chunks found inside chunkHandlers.GetChunksByOwnerID()\n", len(chunks))
			// for i := 0; i < len(chunks); i++ {
			// 	log.Printf("  %v\n", chunks[i])
			// }
			log.Printf("Successful select of all chunks in service.GetChunksByOwnerID()")
			json.NewEncoder(w).Encode(chunks)
			return
		}
		log.Printf("Error with chunk.GetAllResults() in chunkHandlers.GetChunksByOwnerID(), Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in chunkHandlers.GetChunksByOwnerID(), Reason:  %v\n", paramErr)
}

// GetChunksByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetChunksByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>chunkHandlers..GetChunksByUserID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myChunk := &Chunk{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   UserID=%d\n", paramInt)

	if paramErr == nil {
		myChunk.ID = paramInt
		chunks, sqlErr := myChunk.GetChunksByUserID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d chunks found inside chunkHandlers.GetChunksByUserID()\n", len(chunks))
			// for i := 0; i < len(chunks); i++ {
			// 	log.Printf("  %v\n", chunks[i])
			// }
			log.Printf("Successful select of all chunks in service.GetChunksByUserID()")
			json.NewEncoder(w).Encode(chunks)
			return
		}
		log.Printf("Error with chunk.GetAllResults() in chunkHandlers.GetChunksByUserID(), Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in chunkHandlers.GetChunksByUserID(), Reason:  %v\n", paramErr)
}
