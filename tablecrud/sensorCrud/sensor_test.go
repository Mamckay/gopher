package sensorComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"testing"
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
	log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestGetAllResults()")
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	myDB := db.Connect()
	myUser := &Sensor{}
	sensors, errTest := myUser.GetAllResults(myDB)
	if errTest != nil {
		log.Printf("   Test Failed: %s", errTest)
		log.Printf("   %v", sensors)
		return
	}

	log.Printf("Select successful for ID: %v\n", sensors)
	// for i := 0; i < len(sensors); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensors[i].SensorID,
	// 		sensors[i].UserID,
	// 		sensors[i].SensorName,
	// 		sensors[i].Location,
	// 		sensors[i].InputURL,
	// 		sensors[i].SecurityLevel)
	// }
}

// TestGetUserById ...
func TestGetSensorById(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestGetSensorById()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Sensor{ID: 1}
	myUser.GetByID(myDB)
}

// TestGetUserByName ...
func TestGetSensorByName(t *testing.T) {
	log.Printf("-------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestGetSensorByName()")
	log.Printf("-------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Sensor{Sensorname: "SensorID1"}
	myUser.GetByName(myDB)
}

func TestDeleteSensorWithId(t *testing.T) {
	log.Printf("----------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestDeleteSensorWithId()")
	log.Printf("----------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Sensor{ID: 3}
	myUser.Delete(myDB)
}

func TestDeleteSensorWithName(t *testing.T) {
	log.Printf("------------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestDeleteSensorWithName()")
	log.Printf("------------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Sensor{Sensorname: "SensorID2"}
	myUser.Delete(myDB)
}

/*=====================================================================*\
	Functions below this point have been prepared for integrations

\*=====================================================================*/

// TestGetUsersBySensorID ...
func TestGetSensorssByUserID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Sensor{}
	user, errTest := myUser.GetSensorssByUserID(myDB, 2)
	if errTest != nil {
		log.Printf("   Test Failed in TestGetUsersBySensorID: %v", errTest)
		log.Printf("   %v", user)
		return
	}
	fmt.Println("Sensor ", user.ID)
	for i := 0; i < len(user.Sensors); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", user.Sensors[i])
	}
	// for i := 0; i < len(sensor); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensor[i].SensorID,
	// 		sensor[i].UserID,
	// 		sensor[i].SensorName,
	// 		sensor[i].Location,
	// 		sensor[i].InputURL,
	// 		sensor[i].SecurityLevel)
	// }
}

// TestGetUsersBySensorID ...
func TestGetUsersBySensorID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Sensor{}
	sensor, errTest := myUser.GetUsersBySensorID(myDB, 2)
	if errTest != nil {
		log.Printf("   Test Failed in TestGetUsersBySensorID: %v", errTest)
		log.Printf("   %v", sensor)
		return
	}
	fmt.Println("Sensor ", sensor.ID)
	for i := 0; i < len(sensor.Users); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", sensor.Users[i])
	}
	// for i := 0; i < len(sensor); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensor[i].SensorID,
	// 		sensor[i].UserID,
	// 		sensor[i].SensorName,
	// 		sensor[i].Location,
	// 		sensor[i].InputURL,
	// 		sensor[i].SecurityLevel)
	// }
}

// TestGetAllResult ...
// func TestGetResultsSensorsByUser(t *testing.T) {
// 	log.Printf("--------------------------------------------------------")
// 	//log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestGetAllResults()")
// 	trace2()
// 	fmt.Println(nameOf((*A).Method))
// 	log.Printf("--------------------------------------------------------")
// 	myDB := db.Connect()
// 	mySensor := &Sensor{}

// 	//myUser.GetResultsByUserID(myDB, 2)
// 	userResult, errTest := mySensor.GetUsersBySensorID(myDB, 2) // myUser.GetAllResults(myDB)
// 	if errTest != nil {
// 		log.Printf("   Test Failed: %s", errTest)
// 		log.Printf("   %v", userResult)
// 		return
// 	}

// 	log.Printf("Select successful for ID: %v\n", userResult)
// 	for i := 0; i < len(userResult.Users); i++ {
// 		log.Printf("   " + userResult.Users[i].Username)
// 	}

// 	// for i := 0; i < len(sensors); i++ {
// 	// 	log.Printf("   %d %d %s %s %s %s", sensors[i].SensorID,
// 	// 		sensors[i].UserID,
// 	// 		sensors[i].SensorName,
// 	// 		sensors[i].Location,
// 	// 		sensors[i].InputURL,
// 	// 		sensors[i].SecurityLevel)
// 	// }
// }

func TestGetSensorsByOwner(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	//log.Printf("RUNNING: ===>tablecrud/sensorCrud.TestGetAllResults()")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	sensorObject := new(Sensor)

	mySensors, errTest := sensorObject.GetSensorsByOwner(myDB, 2) // myUser.GetAllResults(myDB)
	if errTest != nil {
		log.Printf("   Test Failed: %s", errTest)
		log.Printf("   %v", mySensors)
		return
	}

	log.Printf("Select successful for ID: %v\n", mySensors)
	for i := 0; i < len(mySensors); i++ {
		log.Printf("   " + mySensors[i].Sensorname)
	}

	// for i := 0; i < len(sensors); i++ {
	// 	log.Printf("   %d %d %s %s %s %s", sensors[i].SensorID,
	// 		sensors[i].UserID,
	// 		sensors[i].SensorName,
	// 		sensors[i].Location,
	// 		sensors[i].InputURL,
	// 		sensors[i].SecurityLevel)
	// }
}

func TestGetNotUsersBySensorID(t *testing.T) {
	myDB := db.Connect()
	sensorObject := new(Sensor)

	id := 1
	myUsers, errTest := sensorObject.GetNotUsersBySensorID(myDB, id)
	if errTest != nil {
		log.Printf("   Test Failed: %s", errTest)
		log.Printf("   %v", myUsers)
		return
	}

	log.Printf("Select successful for ID: %v\n", id)
	for i := 0; i < len(myUsers); i++ {
		log.Printf("   %v\n", myUsers[i])
	}

}

var ctx = context.Background()

func TestGetTheNotUsersBySensorID(t *testing.T) {
	//myDB := db.Connect()

	log.Printf("ctx=%v\n", ctx)

	var thedb *sql.DB

	rows, err := thedb.QueryContext(ctx, "select * from entusers")
	return

	if err != nil {
		log.Printf("Error with database select, Reason:  %v\n", err)
		return
	}
	log.Printf("%v\n", rows)
}
