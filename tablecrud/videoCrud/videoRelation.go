/*=====================================================================*\



\*=====================================================================*/
package videoComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

// UserToVideo ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type UserToVideo struct {
	//tableName struct{} `sql:"relsensoruser"`
	UserID  int //`sql:"sensor_id, type:int references sensor(id)"`
	VideoID int //`sql:"user_id, type:int references infuser(id)"`
}

// HandleUserToVideoRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func HandleUserToVideoRel(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>productHandlers.InsertUserToVideoRel()")
	w.Header().Add("Content-Type", "application/json")

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aUserToVideo UserToVideo
	json.Unmarshal([]byte(har), &aUserToVideo)
	fmt.Printf("unmarshalled aUserToVideo=%v", aUserToVideo)

	UserID := aUserToVideo.UserID
	VideoID := aUserToVideo.VideoID

	myDB := db.Connect()
	sqlErr := InsertUserToVideoRel(myDB, UserID, VideoID)
	if sqlErr == nil {
		log.Printf("Successful Insert of UserID=%d to VideoID=%d\n", UserID, VideoID)
		json.NewEncoder(w).Encode("Success")
		return
	}
	log.Printf("Error Insert nsert of UserID=%d to VideoID=%d\n", UserID, VideoID)
	log.Printf("Reason:%v\n", sqlErr)
}

// InsertUserToVideoRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertUserToVideoRel(dbRef *pg.DB, userID int, videoID int) error {
	log.Printf("===>productRelation.InsertUserToVideoRel(userID=%d,videoID=%d)\n", userID, videoID)
	myDB := db.Connect()
	usr2vid := &UserToVideo{UserID: userID, VideoID: videoID}

	insertErr := myDB.Insert(usr2vid)
	if insertErr != nil {
		log.Printf("Error writing to UserToVideo Table in prodUserRelation.InsertUserToVideoRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("UserToVideo Relation inserted successfully into UserToVideo Table")
	return nil
}

// CreateUserToVideoTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateUserToVideoTable() error {
	log.Printf("===>userTester.CreateUserToVideoTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &UserToVideo{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&UserToVideo{}, opts)
	if createErr != nil {
		log.Printf("Error creating UserToVideo table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("UserToVideo Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

/*=====================================================================*\

	UNUSED VIDEO RELATIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
// VideoToUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type VideoToUser struct {
	//tableName struct{} `sql:"relsensoruser"`
	VideoID int //`sql:"sensor_id, type:int references sensor(id)"`
	UserID  int //`sql:"user_id, type:int references infuser(id)"`
}

// InsertVideoToUserRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertVideoToUserRel(dbRef *pg.DB, videoID int, userID int) error {
	log.Printf("===>productRelation.InsertVideoToUserRel(videoID=%d,userID=%d)\n", videoID, userID)

	myDB := db.Connect()
	vid2usr := &VideoToUser{VideoID: videoID, UserID: userID}

	insertErr := myDB.Insert(vid2usr)
	if insertErr != nil {
		log.Printf("Error writing to VideoToUser Table in prodUserRelation.InsertVideoToUserRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("VideoToUser Relation inserted successfully into VideoToUser Table")
	return nil
}

// CreateVideoToUserTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateVideoToUserTable() error {
	log.Printf("===>userTester.CreateVideoToUserTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &VideoToUser{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&VideoToUser{}, opts)
	if createErr != nil {
		log.Printf("Error creating VideoToUser table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("VideoToUser Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}
