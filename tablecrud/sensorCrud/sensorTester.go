/*=====================================================================*\
	This is the UnitTest for the Backend Sensor Component


\*=====================================================================*/
package sensorComp

import (
	"log"

	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
)

/*---------------------------------------------------------------------*\
These testers are easy to call
\*---------------------------------------------------------------------*/

func GetAllSensorsByUserId() {
	log.Printf("===>sensorTester.GetAllSensorsByUserId()")

	myDB := db.Connect()

	mySensor := &Sensor{}
	mySensor.GetSensorsByOwnerID(myDB, 1)
	mySensor.GetSensorsByOwnerID(myDB, 2)
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
// func GetAllSensors(dbRef *pg.DB) {
// 	log.Printf("===>sensorTester.GetAllSensors()")

// 	mySensor := &Sensor{}
// 	mySensor.GetAllResults(dbRef)
// }
// func GetSensorByName(dbRef *pg.DB, sensorName string) {
// 	log.Printf("===>sensorTester.GetSensorByName()")

// 	mySensor := &Sensor{SensorName: sensorName}
// 	mySensor.GetByName(dbRef)
// }
// func GetSensorById(dbRef *pg.DB, id int) {
// 	log.Printf("===>sensorTester.GetSensorById()")

// 	mySensor := &Sensor{SensorID: id}
// 	mySensor.GetByID(dbRef)
// }
// func DeleteSensorWithId(dbRef *pg.DB, id int) {
// 	log.Printf("===>sensorTester.DeleteSensorWithId()")

// 	mySensor := &Sensor{SensorID: id}
// 	mySensor.Delete(dbRef)
// }
// func DeleteSensorWithName(dbRef *pg.DB, sensorName string) {
// 	log.Printf("===>sensorTester.DeleteSensorWithName()")

// 	mySensor := &Sensor{SensorName: sensorName}
// 	mySensor.Delete(dbRef)
// }
