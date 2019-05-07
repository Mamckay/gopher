// package sensorComp ...
/*=====================================================================*\


\*=====================================================================*/
package sensorComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	// usr "AdvDashBoard/goAdvDashBoard/userComp"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Sensor ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
type Sensor struct {
	tableName   struct{} `sql:"entsensors"`
	ID          int      `json:"SensorID,omitempty" sql:"id,pk"`
	UserID      int      `json:"UserID,omitempty" sql:"userid"`
	Sensorname  string   `json:"SensorName,omitempty"`
	Sensortype  string   `json:"SensorType,omitempty"`
	Sensormode  string   `json:"SensorMode,omitempty"`
	Location    string   `json:"Location,omitempty"`
	Sensorowner string   `json:"SensorOwner,omitempty" sql:"sensorowner"`
	Sensorproto string   `json:"SensorProto,omitempty"`
	Sensorauth  string   `json:"SensorAuth,omitempty"`
	LocalUrl    string   `json:"LocalUrl,omitempty" sql:"localurl"`
	Security    string   `json:"Security,omitempty" sql:"securelevel"`
	SensorState int      `json:"SensorState,omitempty" sql:"sensorstate"`
	SensorPid   int      `json:"SensorPid,omitempty" sql:"sensorpid" `
	Users       []User   `json:"users,omitempty" pg:",many2many:sensor_to_users,joinFK:sensor_id, joinFK:user_id"`
	//`pg:",many2many:relsensoruser,joinFK:user_id, joinFK:sensor_id"`
	// `pg:"many2many:relsensoruser,joinFK:user_id"` // json:"useritems,omitempty"    :sen_user_item,joinFK:user_id
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
	Sensors   []Sensor `pg:",many2many:user_to_sensors",joinFK:userowner_id, joinFK:sensor_id"`
}

// CreateSensorTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSensorTable() error {
	log.Printf("===>sensorTester.CreateSensorTable()")
	myDB := db.Connect()

	delopts := &orm.DropTableOptions{
		IfExists: false,
	}
	orm.DropTable(myDB, &Sensor{}, delopts)
	orm.DropTable(myDB, &SensorToUser{}, delopts)

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := myDB.CreateTable(&Sensor{}, opts)
	if createErr != nil {
		log.Printf("Error creating sensor table Sensor, Reason:%v\n", createErr)
		return createErr
	}
	log.Printf("Sensor Info Table created successfully. Only if necessary.\n")
	//LoadSensorTable(myDB)
	return nil
}

// LoadSensorTable ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func LoadSensorTable(dbRef *pg.DB) {
	log.Println("===>sensorTable.LoadSensorTable()")

	var testCase = 1
	if testCase == 1 {
		InsertSensorEnt(dbRef,
			1,
			"SensorName",
			"SensorType",
			"SensorMode",
			"Location",
			"SensorOwner",
			"rtsp",
			"192.168.1.10",
			"InputURL",
			"Security",
			6,
			8,
		)
	} else if testCase == 2 {
		//GetAllSensors(dbRef)
	} else if testCase == 3 {
	} else if testCase == 4 {
	}
}

// CreateSensorWithName ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func InsertSensorEnt(dbRef *pg.DB,
	userID int,
	sensorName string,
	sensorType string,
	sensorMode string,
	location string,
	sensorowner string,
	sensorproto string,
	sensorauth string,
	localurl string,
	security string,
	state int,
	pid int,
) {
	log.Println(">===>sensorTable.InsertSensorEnt()")
	mySensor := Sensor{
		UserID:      userID,
		Sensorname:  sensorName,
		Sensortype:  sensorType,
		Sensormode:  sensorMode,
		Location:    location,
		Sensorowner: sensorowner,
		Sensorproto: sensorproto,
		Sensorauth:  sensorauth,
		LocalUrl:    localurl,
		Security:    security,
		SensorState: state,
		SensorPid:   pid,
	}
	mySensor.Create(dbRef)
}
