/*===================================================================================*\
	The file contains the Request-Response Handlers for the Sensor Capture Component

			// cmdName = "touch"
			// cmdArgs = []string{"../ngAdvDashBoard/src/main.ts"}
			// ds.ExecWithScanner(cmdName, cmdArgs)

\*===================================================================================*/
package senCapComp

import (
	sen "AdvDashBoard/goAdvDashBoard/tablecrud/sensorCrud"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LiveCapture ...
/*---------------------------------------------------------------------*\
		//rtsp://admin:admin@192.168.0.16/defaultPrimary?streamType=u")

\*---------------------------------------------------------------------*/
func LiveCapture(w http.ResponseWriter, r *http.Request) {
	log.Println("===>senCapHandler.LiveCapture()")
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	var mode = params["mode"]

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aSensor sen.Sensor
	json.Unmarshal([]byte(har), &aSensor)
	fmt.Printf("unmarshalled aSensor=%v\n", aSensor)

	LiveCaptureWithIPSen(aSensor, mode)
}

// FileCapture ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func FileCapture(w http.ResponseWriter, r *http.Request) {
	log.Println("===>senCapHandler.FileCapture()")
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	var mode = params["mode"]

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aSensor sen.Sensor
	json.Unmarshal([]byte(har), &aSensor)
	fmt.Printf("unmarshalled aSensor=%v\n", aSensor)

	FileCaptureWithIPSen(aSensor, mode)
}

// ChunkCapture ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func ChunkCapture(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sencapHandler.ChunkCapture()")
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	var mode = params["mode"]

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aSensor sen.Sensor
	json.Unmarshal([]byte(har), &aSensor)
	fmt.Printf("unmarshalled aSensor=%v\n", aSensor)

	ChunkCaptureWithIPSen(aSensor, mode)
}
