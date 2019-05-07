/*=====================================================================*\
	This


\*=====================================================================*/
package chunkComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Chunk ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type Chunk struct {
	tableName  struct{}  `sql:"entchunks"`
	ID         int       `sql:"id,pk" json:"ChunkID,omitempty"`
	ChunkName  string    `sql:"chunkname" json:"ChunkName,omitempty"`
	ChunkType  string    `sql:"chunktype" json:"ChunkType,omitempty"`
	SensorID   int       `sql:"sensorid" json:"SensorID,omitempty"`
	UserID     int       `sql:"userid" json:"UserID,omitempty"`
	OwnerName  string    `sql:"ownername" json:"OwnerName,omitempty"`
	FilePath   string    `sql:"filepath" json:"FilePath,omitempty"`
	ChunkState int       `sql:"chunkstate" json:"ChunkState,omitempty"`
	ChunkPid   int       `sql:"chunkpid" json:"ChunkPid,omitempty"`
	Videos     []Video   `pg:",many2many:chunk_to_videos"`
	Products   []Product `pg:",many2many:chunk_to_products"`
	Users      []User    `json:"users,omitempty" pg:",many2many:sensor_to_users,joinFK:sensor_id, joinFK:user_id"`
}

/*---------------------------------------------------------------------*\
		THIS IS NOT A REAL TABLE, COPY THIS FROM userTable.go
\*---------------------------------------------------------------------*/
type User struct {
	tableName struct{} `sql:"entusers"`
	ID        int      `sql:"id,pk" json:"UserID,omitempty"`
	UserName  string   `sql:"username,unique" json:"UserName,omitempty"`
	Password  string   `sql:"password" json:"Password,omitempty"`
	FirstName string   `sql:"firstname" json:"FirstName,omitempty"`
	LastName  string   `sql:"lastname" json:"LastName,omitempty"`
	RoleName  string   `sql:"rolename" json:"RoleName,omitempty"`
	JwtToken  string   `sql:"jwttoken" json:"JwtToken,omitempty"`
	Chunks    []Chunk  `pg:",many2many:user_to_chunks",joinFK:userowner_id, joinFK:sensor_id"`
}

// Video ...
/*---------------------------------------------------------------------*\
	CAUTION: THIS IS NOT A REAL TABLE: Copy it from VideoTable.go
\*---------------------------------------------------------------------*/
type Video struct {
	tableName struct{} `sql:"entvideos"`
	ID        int      `sql:"id,pk" json:"VideoID,omitempty"`
	VideoName string   `sql:"videoname" json:"VideoName,omitempty"`
	VideoType string   `sql:"videotype" json:"VideoType,omitempty"`
	SensorID  int      `sql:"sensorid, type:int references sensor(id)" json:"SensorID,omitempty"`
	UserID    int      `sql:"userid, type:int references infuser(id)" json:"UserID,omitempty"`
	StartTime string   `sql:"starttime" json:"StartTime,omitempty"`
	EndTime   string   `sql:"endtime" json:"EndTime,omitempty"`
	InputURL  string   `sql:"inputurl" json:"InputURL,omitempty"`
	FilePath  string   `sql:"filepath" json:"FilePath,omitempty"`
	// Users     []User   `pg:",many2many:video_to_users"`
}

/*---------------------------------------------------------------------*\
	CAUTION: THIS IS NOT A REAL TABLE: Copy it from ProductTable.go
\*---------------------------------------------------------------------*/
type Product struct {
	tableName struct{} `sql:"entproducts"`
	// ID           int      `sql:"type:int references userinfo(userid)"`
	ProductID    int    `sql:"id,pk" json: "ProductID,omitempty"`
	ProductName  string `sql:"productname" json: "ProductName,omitempty"`
	ProductType  string `sql:"producttype" json: "ProductType,omitempty"`
	UserID       int    `sql:"userid" json: "UserID,omitempty"`
	SensorID     int    `sql:"sensorid" json: "SensorID,omitempty"`
	OwnerName    string `sql:"ownername" json: "OwnerName,omitempty"`
	FilePath     string `sql:"filepath" json: "FilePath,omitempty"`
	ProductState int    `sql:"productstate" json: "ProductState,omitempty"`
	ProductPid   int    `sql:"productpid" json: "ProductPid,omitempty"`
	//Users        []*User `pg:",many2many:product_to_users"`
}

// CreateChunkTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateChunkTable() error {
	log.Printf(">===>chunkTable.CreateChunkTable()")
	myDB := db.Connect()

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := myDB.CreateTable(&Chunk{}, opts)
	if createErr != nil {
		log.Printf("Error creating Gaol table chunks, Reason:%v\n", createErr)
		return createErr
	}
	log.Printf("Chunk Info Table created successfully. Only if necessary.\n")
	//LoadChunkTable(myDB)
	return nil
}

// LoadChunkTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadChunkTable(dbRef *pg.DB) {
	log.Println(">===>chunkTable.LoadChunkTable()")

	var testCase = 1
	MyChunk := &Chunk{}
	if testCase == 1 {
		InsertChunkEnt(dbRef,
			2,

			2,
			"Chunk One",
			"S3VC",
			"SomeOwner",
			"assets/Products/dir1/",
			1,
			1234,
		)
		InsertChunkEnt(dbRef,
			3,
			2,

			"Chunk Two",
			"Avigilon",
			"OperatorUser",
			"assets/Products/dir1/",
			1,
			1234,
		)
		InsertChunkEnt(dbRef,
			4,
			2,
			"Chunk Three",
			"CudaEye",
			"SomeOwner",
			"assets/Products/dir1/",
			1,
			1234,
		)

	} else if testCase == 2 {
		//		GetAllChunks(dbRef)
		MyChunk.GetChunksFullList(dbRef)
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}

// CreateChunkWithName ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertChunkEnt(dbRef *pg.DB,
	userID int,
	sensorID int,
	chunkName string,
	chunkType string,
	ownername string,
	filepath string,
	chunkstate int,
	chunkpid int,
) {
	log.Printf(">===>chunkTable.InsertChunkEnt()")
	myChunk := Chunk{
		UserID:     userID,
		SensorID:   sensorID,
		ChunkName:  chunkName,
		ChunkType:  chunkType,
		OwnerName:  ownername,
		FilePath:   filepath,
		ChunkState: chunkstate,
		ChunkPid:   chunkpid,
	}
	myChunk.Create(dbRef)
}
