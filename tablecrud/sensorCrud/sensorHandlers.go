/*=====================================================================*\
	This is the Request/Reply handler of Backend Sensor
	Table Component. It is responsble for receiving RestAPI
	requests from the Frontend, interacting with the
	fileStore, constructing	and returning a response.
\*=====================================================================*/
package sensorComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"github.com/rs/cors"
)

// GetUsersBySensorID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetUsersBySensorID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetUsersBySensorID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	mySensor := &Sensor{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("  paramInt == %d", paramInt)
	if paramErr == nil {
		mySensor.ID = paramInt
		userList, sqlErr := mySensor.GetUsersBySensorID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d Users found in sensorHandlers.GetUsersBySensorID() sensor=%v\n", len(userList))
			json.NewEncoder(w).Encode(userList)
			return
		}
		log.Printf("Error with database select in sensorHandlers.GetUsersBySensorID(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in sensorHandlers.GetUsersBySensorID(), Reason:  %v\n", paramErr)
}

// GetUsersBySensorID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetNotUsersBySensorID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetNotUsersBySensorID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	mySensor := &Sensor{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("  paramInt == %d", paramInt)
	if paramErr == nil {
		mySensor.ID = paramInt
		users, sqlErr := mySensor.GetNotUsersBySensorID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("Successful select in sensorHandlers.GetNotUsersBySensorID() users=%v\n", users)
			json.NewEncoder(w).Encode(users)
			return
		}
		log.Printf("Error with database select in sensorHandlers.GetNotUsersBySensorID(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in sensorHandlers.GetNotUsersBySensorID(), Reason:  %v\n", paramErr)
}

// CreateSensor ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateSensor(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>sensorHandlers.CreateSensor()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)

	var sensor Sensor
	json.Unmarshal([]byte(har), &sensor)
	fmt.Printf("%v", sensor)

	myDB := db.Connect()

	log.Printf("   sensorHandlers.Create():=%v\n", sensor)

	sqlErr := sensor.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created sensor sensorHandlers.CreateSensor(): sensor=%v\n", sensor)
		json.NewEncoder(w).Encode(sensor)
		return
	}
	log.Printf("Error creating sensor in sensorHandlers.CreateSensor(), Reason:%v\n", sqlErr)
}

// UpdateSensor ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func UpdateSensor(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.UpdateSensor()")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)

	pSensor := params["Sensor"]
	log.Printf("   Sensor=%v\n", pSensor)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var sensor Sensor
	json.Unmarshal([]byte(har), &sensor)
	fmt.Printf("%v", sensor)

	sqlErr := sensor.Update(myDB)
	if sqlErr == nil {
		log.Printf("Successful update in sensorHandlers.UpdateSensor() sensor=%v\n", sensor)
		json.NewEncoder(w).Encode(sensor)
		return
	}
	log.Printf("Error updateing sensor in sensorHandlers.UpdateSensor(), Reason:%v\n", sqlErr)
}

//DeleteSensor ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>sensorHandlers.DeleteSensor()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	paramInt, paramErr := strconv.Atoi(params["id"])
	mySensor := &Sensor{ID: paramInt}

	if paramErr == nil {
		mySensor.ID = paramInt
		sqlErr := mySensor.Delete(myDB)
		if sqlErr == nil {
			log.Printf("Successful delete in service.DeleteSensors() sensor=%v\n", mySensor)
			json.NewEncoder(w).Encode(mySensor)
			return
		}
		log.Printf("Error with database delete in service.DeleteSensors(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in sensorHandlers.DeleteSensors(), Reason:  %v\n", paramErr)
}

// GetSensor ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetSensor(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetSensor()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	mySensor := &Sensor{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		mySensor.ID = paramInt
		sensor, sqlErr := mySensor.GetByID(myDB)
		if sqlErr == nil {
			log.Printf("Successful select in service.GetSensors() sensor=%v\n", sensor)
			json.NewEncoder(w).Encode(sensor)
			return
		}
		log.Printf("Error with database select in service.GetSensor(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in sensorHandlers.GetSensor(), Reason:  %v\n", paramErr)
}

/*=====================================================================*\

	Handlers Functions completed the END 2 END naming conversion

\*=====================================================================*/
// GetSensorsFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetSensorsFullList(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetSensorsFullList()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	mySensor := &Sensor{}
	sensors, sqlErr := mySensor.GetSensorsFullList(myDB)
	log.Printf("%d sensors found inside sensorHandlers.GetSensorsFullList()\n", len(sensors))
	// for i := 0; i < len(sensors); i++ {
	// 	log.Printf("   %v\n", sensors[i])
	// }

	if sqlErr == nil {
		log.Printf("Successful select of all sensors in sensorHandlers.GetSensorsFullList()")
		json.NewEncoder(w).Encode(sensors)
		return
	}
	log.Printf("Error with database select in sensorHandlers.GetSensorsFullList(), Reason:%v\n", sqlErr)
}

// GetSensorsByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetSensorsByOwnerID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetSensorsByOwnerID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	mySensor := &Sensor{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		mySensor.ID = paramInt
		sensors, sqlErr := mySensor.GetSensorsByOwnerID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("Successful select in sensorHandlers.GetSensorsByOwnerID() sensor=%v\n", sensors)
			json.NewEncoder(w).Encode(sensors)
			return
		}
		log.Printf("Error with database select in sensorHandlers.GetSensorsByOwnerID(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in sensorHandlers.GetSensorsByOwnerID(), Reason:  %v\n", paramErr)
}

// GetSensorsByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetSensorsByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetSensorsByUserID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	mySensor := &Sensor{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		mySensor.ID = paramInt
		sensors, sqlErr := mySensor.GetSensorsByUserID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("Successful select in sensorHandlers.GetSensorsByUserID() sensor=%v\n", sensors)
			json.NewEncoder(w).Encode(sensors)
			return
		}
		log.Printf("Error with database select in sensorHandlers.GetSensorsByUserID(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in sensorHandlers.GetSensorsByUserID(), Reason:  %v\n", paramErr)
}
