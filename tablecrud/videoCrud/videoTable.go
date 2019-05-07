// package videoComp ...
/*=====================================================================*\


\*=====================================================================*/
package videoComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"log"

	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

// Video ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type Video struct {
	tableName struct{} `sql:"entvideos"`
	ID        int      `sql:"id,pk" json:"VideoID,omitempty"`
	VideoName string   `sql:"videoname"`
	VideoType string   `sql:"videotype"`
	SensorID  int      `sql:"sensorid, type:int references entsensors(id)"`
	UserID    int      `sql:"userid, type:int references entusers(id)"`
	StartTime string   `sql:"starttime"`
	EndTime   string   `sql:"endtime"`
	InputURL  string   `sql:"inputurl"`
	FilePath  string   `sql:"filepath"`
	Users     []User   `pg:",many2many:video_to_users"`
}

// User ...
/*---------------------------------------------------------------------*\
	THIS IS A REFERENCE TABLE, COPY FROM ACTUAL inside userTable.go
\*---------------------------------------------------------------------*/
type User struct {
	tableName struct{} `sql:"entusers"`
	ID        int      `sql:"id,pk"`
	UserName  string   `sql:"username,unique"`
	Password  string   `sql:"password"`
	FirstName string   `sql:"firstname"`
	LastName  string   `sql:"lastname"`
	RoleName  string   `sql:"rolename"`
	JwtToken  string   `sql:"jwttoken"`
	Videos    []Video  `pg:",many2many:user_to_videos"`
}

// CreateVideoTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateVideoTable() error {
	log.Printf("===>videoTable.CreateVideoTable()")
	myDB := db.Connect()
	delopt := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &Video{}, delopt)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&Video{}, opts)
	if createErr != nil {
		log.Printf("Error creating Video table, Reason:%v\n", createErr)
		return createErr
	}
	log.Printf("Video Info Table created successfully. Only if necessary.\n")
	//LoadVideoTable(myDB)
	return nil
}

// LoadVideoTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadVideoTable(dbRef *pg.DB) {
	log.Println("===>videoTable.LoadVideoTable()")

	var testCase = 1
	if testCase == 1 {
		InsertVideoEnt(dbRef,
			"Melonga Dancers",
			"Avigilon",
			1,
			1,
			"2018/07/24 11:18:49",
			"2018/07/24 12:18:49 ",
			"198.162.0.10",
			"../../assets/Melonga_160x120.mp4",
		)
		InsertVideoEnt(dbRef,
			"Tango Dancers",
			"Avigilon",
			1,
			1,
			"2018/07/24 09:18:49",
			"2018/07/24 10:18:49 ",
			"198.162.0.10",
			"../../assets/Tango_160x120.mp4",
		)
	} else if testCase == 2 {
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}

// InsertVideoEnt ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertVideoEnt(dbRef *pg.DB,
	videoName string,
	videoType string,
	sensorid int,
	userid int,
	startTime string,
	endTime string,
	inputurl string,
	filePath string,
) {
	log.Printf("===>videoTable.InsertVideoEnt()")
	myVideo := Video{
		VideoName: videoName,
		VideoType: videoType,
		SensorID:  sensorid,
		UserID:    userid,
		StartTime: startTime,
		EndTime:   endTime,
		InputURL:  inputurl,
		FilePath:  filePath,
	}
	myVideo.Create(dbRef)
}
