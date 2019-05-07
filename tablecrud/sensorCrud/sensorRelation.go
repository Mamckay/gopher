package sensorComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	//"github.com/rs/cors"
)

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type UserToSensor struct {
	//tableName struct{} `sql:"relsensoruser"`
	UserID   int //`sql:"sensor_id, type:int references sensor(id)"`
	SensorID int //`sql:"user_id, type:int references infuser(id)"`
}

// HandleUserToSensorRel ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func HandleUserToSensorRel(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>productHandlers.InsertUserToSensorRel()")
	w.Header().Add("Content-Type", "application/json")

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aUserToSensor UserToSensor
	json.Unmarshal([]byte(har), &aUserToSensor)
	fmt.Printf("unmarshalled aUserToSensor=%v", aUserToSensor)

	UserID := aUserToSensor.UserID
	SensorID := aUserToSensor.SensorID

	myDB := db.Connect()
	sqlErr := InsertUserToSensorRel(myDB, UserID, SensorID)
	if sqlErr == nil {
		log.Printf("Successful Insert of UserID=%d to SensorID=%d\n", UserID, SensorID)
		json.NewEncoder(w).Encode("Success")
		return
	}
	log.Printf("Error Insert nsert of UserID=%d to SensorID=%d\n", UserID, SensorID)
	log.Printf("Reason:%v\n", sqlErr)
}

// CreateUserToSensorTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateUserToSensorTable() error {
	log.Printf("===>senUserRelation.CreateUserToSensorTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &UserToSensor{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&UserToSensor{}, opts)
	if createErr != nil {
		log.Printf("Error creating UserToSensor table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("UserToSensor Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertUserToSensorRel(dbRef *pg.DB, userID int, sensorID int) error {
	log.Printf("===>senUserRelation.InsertUserToSensorRel()")
	myDB := db.Connect()
	usr2sen := &UserToSensor{UserID: userID, SensorID: sensorID}

	insertErr := myDB.Insert(usr2sen)
	if insertErr != nil {
		log.Printf("Error writing to UserToSensor Table in prodUserRelation.InsertUserToSensorRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("UserToSensor Relation inserted successfully into UserToSensor Table")
	return nil
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteUserToSensorRel(dbRef *pg.DB, userID int, sensorID int) error {
	log.Printf("===>senUserRelation.DeleteUserToSensorRel()")
	myDB := db.Connect()
	usr2sen := &UserToSensor{UserID: userID, SensorID: sensorID}

	insertErr := myDB.Insert(usr2sen)
	if insertErr != nil {
		log.Printf("Error writing to UserToSensor Table in prodUserRelation.InsertUserToSensorRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("UserToSensor Relation inserted successfully into UserToSensor Table")
	return nil
}

/*=====================================================================*\

	THESE ARE NOW USED AGAIN

\*=====================================================================*/

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type SensorToUser struct {
	//tableName struct{} `sql:"relsensoruser"`
	SensorID int //`sql:"sensor_id, type:int references sensor(id)"`
	UserID   int //`sql:"user_id, type:int references infuser(id)"`
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertSensorToUserRel(dbRef *pg.DB, sensorID int, userID int) error {
	log.Printf("===>senUserRelation.InsertSensorToUserRel(sensorID=%d,userID=%d)\n", sensorID, userID)
	myDB := db.Connect()
	sen2usr := &SensorToUser{SensorID: sensorID, UserID: userID}

	insertErr := myDB.Insert(sen2usr)
	if insertErr != nil {
		log.Printf("Error writing to SendorToUser Table in senUserRelation.InsertSendorToUserRel()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("SendorToUser Relation inserted successfully into SendorToUser Table")
	return nil
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSensorToUserTable() error {
	log.Printf("===>senUserRelation.CreateSensorToUserTable()")
	myDB := db.Connect()
	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &SensorToUser{}, delopts)
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&SensorToUser{}, opts)
	if createErr != nil {
		log.Printf("Error creating SensorToUser table, Reason:%v\n", createErr)
		return createErr
	}

	log.Printf("SensorToUser Table created successfully. Only if necessary.\n")
	// LoadUserTable(myDB)
	return nil
}

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *SensorToUser) Create(db *pg.DB) error {
	log.Printf("===>senUserModel.Create()")
	log.Printf("   senUserModel.Create():SensorToUser=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in senUserModel.Create(), Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("SenUserIterm %s inserted successfully into table", gi.SensorID)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *SensorToUser) Delete(db *pg.DB) error {
	log.Printf("===>senUserRelation.Delete()")

	_, deleteErr := db.Model(gi).
		Where("sensorid = ?0", gi.SensorID).
		WhereOr("userid = ?0", gi.UserID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting SensorToUser in senUserRelation.Delete(), Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("SensorToUser %s deleted successfully from table", gi.SensorID)
	return nil
}

/*=====================================================================*\

	THESE ARE STILL UNUSED SENSOR RELATIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
// CreateSenUserTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSenUserTable_NOTUSED() error {
	log.Printf("===>sensorTester.CreateSenUserTable()")
	myDB := db.Connect()

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := myDB.CreateTable(&SensorToUser{}, opts)
	if createErr != nil {
		log.Printf("Error creating table SenUser, Reason:%v\n", createErr)
		return createErr
	}
	log.Printf("SenUser table created successfully. Only if necessary.\n")
	LoadSenUserTable_NOTUSED(myDB)
	return nil
}

// LoadSensorTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadSenUserTable_NOTUSED(dbRef *pg.DB) {
	log.Println("===>sensorRelation.LoadSenUserTable()")
	var testCase = 0
	if testCase == 1 {
		CreateSenUserTest_NOTUSED(dbRef, 1, 1)
	} else if testCase == 2 {
		//GetAllSensors(dbRef)
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}

// CreateSenUserTest ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSenUserTest_NOTUSED(dbRef *pg.DB,
	userID int,
	sensorID int,
) {
	log.Printf("===>sensorTester.CreateSenUserTest()")
	mySenUser := SensorToUser{
		UserID:   userID,
		SensorID: sensorID,
	}
	mySenUser.Create(dbRef)
}
