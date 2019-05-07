/*=====================================================================*\


\*=====================================================================*/
package chunkComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"context"
	qdb "database/sql"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"testing"

	_ "github.com/lib/pq"
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
func TestGetAllResults(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/chunkCrud.TestGetAllResults()")
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	myDB := db.Connect()
	myUser := &Chunk{}
	myUser.GetAllResults(myDB)
}

// TestGetAllResult ...
func TestGetResultsChunksByUser(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	//log.Printf("RUNNING: ===>tablecrud/chunkCrud.TestGetAllResults()")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Chunk{}
	myUser.GetResultsByUserID(myDB, 2)
}

// TestGetUserById ...
func TestGetChunkById(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/chunkCrud.TestGetChunkById()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Chunk{ID: 1}
	myUser.GetByID(myDB)
}

// TestGetUserByName ...
func TestGetChunkByName(t *testing.T) {
	log.Printf("-------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/chunkCrud.TestGetChunkByName()")
	log.Printf("-------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Chunk{ChunkName: "ChunkID1"}
	myUser.GetByName(myDB)
}

func TestDeleteChunkWithId(t *testing.T) {
	log.Printf("----------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/chunkCrud.TestDeleteChunkWithId()")
	log.Printf("----------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Chunk{ID: 3}
	myUser.Delete(myDB)
}

func TestDeleteChunkWithName(t *testing.T) {
	log.Printf("------------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/chunkCrud.TestDeleteChunkWithName()")
	log.Printf("------------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Chunk{ChunkName: "ChunkID2"}
	myUser.Delete(myDB)
}

/*=====================================================================*\
	Functions below this point have been prepared for integrations

\*=====================================================================*/
// GetProductsByChunkID
func TestGetProductsByChunkID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myProd := &Chunk{}
	chunkwProdts, errTest := myProd.GetProductsByChunkID(myDB, 3)
	if errTest != nil {
		log.Printf("   Test Failed in GetProductsByChunkID: %v", errTest)
		log.Printf("   %v", chunkwProdts)
		return
	}
	fmt.Println("User ", chunkwProdts.ID)
	for i := 0; i < len(chunkwProdts.Products); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", chunkwProdts.Products[i])
	}
	// for i := 0; i < len(chunkwVideos.Videos); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", chunkwVideos.Videos[i].ID,
	// 		chunkwVideos.Videos[i].UserID,
	// 		chunkwVideos.Videos[i].SensorName,
	// 		chunkwVideos.Videos[i].Location,
	// 		prodchunkwVideos.Videosts[i].InputURL,
	// 		chunkwVideos.Videos[i].SecurityLevel)
	// }
}

// TestGetGetVideosByChunkID
func TestGetVideosByChunkID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myProd := &Chunk{}
	chunkwVideos, errTest := myProd.GetVideosByChunkID(myDB, 4)
	if errTest != nil {
		log.Printf("   Test Failed in GetVideosByChunkID: %v", errTest)
		log.Printf("   %v", chunkwVideos)
		return
	}
	fmt.Println("User ", chunkwVideos.ID)
	for i := 0; i < len(chunkwVideos.Videos); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", chunkwVideos.Videos[i])
	}
	// for i := 0; i < len(chunkwVideos.Videos); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", chunkwVideos.Videos[i].ID,
	// 		chunkwVideos.Videos[i].UserID,
	// 		chunkwVideos.Videos[i].SensorName,
	// 		chunkwVideos.Videos[i].Location,
	// 		prodchunkwVideos.Videosts[i].InputURL,
	// 		chunkwVideos.Videos[i].SecurityLevel)
	// }
}

var ctx = context.Background()
var tstDB qdb.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pg"
	dbname   = "AdvDashBoard"
)

func sqlConnect() (*qdb.DB, error) {
	//log.Printf(">===>pgDBMain.Connect()")
	// opts := &pg.Options{
	// 	User:         "postgres",
	// 	Password:     "pg",
	// 	Addr:         "localhost:5432",
	// 	Database:     "AdvDashBoard",
	// 	DialTimeout:  30 * time.Second,
	// 	ReadTimeout:  1 * time.Minute,
	// 	WriteTimeout: 1 * time.Minute,
	// 	IdleTimeout:  30 * time.Minute,
	// 	MaxAge:       1 * time.Minute, //do a ping to db every minute
	// 	PoolSize:     20,
	// }
	//con := qdb.DB.sql.Open()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//"postgres:pg@tcp(127.0.0.1:5432)/AdvDashBoard"
	db, err := qdb.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Failed to connect to database.\n")
		return nil, err
		//os.Exit(100)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to connect to database.\n")
		return nil, err
	}
	log.Printf("Connection to database [%v] successful.\n", db)
	return db, err
}

func TestQueryChunkID(t *testing.T) {
	log.Printf("  database query \n")
	//var ctx = context.Background()
	tstDB, err := sqlConnect()
	if err != nil {
		log.Printf("   ERROR sqlConnect(): %v\n", err)
		return
	}

	//return
	// rows, err := tstDB.Query("SELECT * FROM entusers")
	rows, err := tstDB.Query("SELECT * FROM entusers as u join entsensors as s on u.id = s.userid and u.id = 4")
	// rows, err := tstDB.Query("SELECT * FROM entusers")
	if err != nil {
		log.Printf("   Test Failed in QueryContext: %v\n", err)
		return
	}
	defer rows.Close()

	records := make([]string, 0)

	for rows.Next() {
		var a string
		var b string
		var c string
		var d string
		var e string
		var f string
		var g string

		var ba string
		var bb string
		var bc string
		var bd string
		var be string
		var bf string
		var bg string
		var bh string
		var bi string
		var bj string
		var bk string
		var bl string
		var bm string

		var err = rows.Scan(&a, &b, &c, &d, &e, &f, &g, &ba, &bb, &bc, &bd, &be, &bf, &bg, &bh, &bi, &bj, &bk, &bl, &bm)
		if err != nil {
			log.Printf("ERROR %v", err)
			return
		}

		records = append(records, a+" "+b+" "+c+" "+d+" "+e+" "+f+" "+g+" "+ba+" "+bb+" "+bc+" "+bd+" "+be+" "+bf+" "+bg+" "+bh+" "+bi+" "+bj+" "+bk+" "+bl+" "+bm+" "+"\n")
	}
	if err := rows.Err(); err != nil {
		log.Printf(" ERROR %v\n", err)
	}
	log.Printf("  Record:  %s", records)

	log.Printf("  database query  %v\n", rows)

}
