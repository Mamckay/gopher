// Package vidPlayComp is the Data Storage Config Package
/*===================================================================================*\
	The vidPlayComp is the Sensor Capture Component


\*===================================================================================*/
package vidPlayComp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ChunkPlay ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func ChunkPlay(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>vidPlayHandler.ChunkPlay()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aSensor sen.Sensor
	json.Unmarshal([]byte(har), &aSensor)
	fmt.Printf("unmarshalled aUser=%v", aSensor)

	mySensor := &sen.Sensor{SensorName: aSensor.SensorName}
	myDB := db.Connect()
	sensor, sqlErr := mySensor.GetByName(myDB)

	if sqlErr == nil {
		log.Printf("Successful GetByName() call in sencapRegResp.ChunkCapture() sensor=%v\n", sensor)
		json.NewEncoder(w).Encode(sensor)
		return

		log.Printf("Error with database select in sencapRegResp.ChunkCapture(), Reason:  %v\n", sqlErr)
		return
	}
}
