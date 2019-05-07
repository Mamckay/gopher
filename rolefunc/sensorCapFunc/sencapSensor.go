// Package senCapComp is the Sensor Capture Functions
/*===================================================================================*\



\*===================================================================================*/
package senCapComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	ds "AdvDashBoard/goAdvDashBoard/datastore/fileStore"
	chk "AdvDashBoard/goAdvDashBoard/tablecrud/chunkCrud"
	sen "AdvDashBoard/goAdvDashBoard/tablecrud/sensorCrud"
	vid "AdvDashBoard/goAdvDashBoard/tablecrud/videoCrud"
	qdb "database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/go-pg/pg"
)

// LiveCaptureWithIPSen ...
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func LiveCaptureWithIPSen(aSensor sen.Sensor, mode string) {
	log.Printf(">===>sencapSensor.LiveCaptureWithIPSen()\n")
	var urlStr = aSensor.Sensorproto + "://" + aSensor.Sensorauth + "@" + aSensor.LocalUrl
	log.Printf("urlStr=%v\n", urlStr)

	myDB := db.Connect()

	if mode == "ON" {
		log.Printf("----- LIVE CAPTURE ON ---------")
		log.Printf("On Time=%s", ds.GetNowTime())

		aPid, err := LiveCaptureWithFFmpeg(aSensor.Sensorname, aSensor.Sensortype, urlStr)
		aSensor.UpdatePid(myDB, aPid, 10)
		log.Printf("   err=%v\n", err)

	} else if mode == "OFF" {
		log.Printf("----- LIVE CAPTURE OFF ---------")
		log.Printf("Off Time=%s", ds.GetNowTime())

		aPid, err := aSensor.GetPidByID(myDB, aSensor.ID)
		log.Printf("   aPid=%d\n", aPid)
		ds.KillWithPid(aPid)
		aSensor.UpdatePid(myDB, aPid, 0)
		log.Printf("   err=%v\n", err)

	} else {
		log.Printf("----- BAD LIVE MODE ---------")
	}
}

// FileCaptureWithIPSen ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func FileCaptureWithIPSen(aSensor sen.Sensor, mode string) {
	log.Printf(">===>sencapSensor.FileCaptureWithIPSen()\n")

	var urlStr = aSensor.Sensorproto + "://" + aSensor.Sensorauth + "@" + aSensor.LocalUrl
	log.Printf("urlStr=%v\n", urlStr)

	myDB := db.Connect()

	if mode == "ON" {
		log.Printf("----- FILE CAPTURE ON ---------")

		aPid, aFil, err := FileCaptureWithFFmpeg(aSensor.Sensorname, aSensor.Sensortype, urlStr)
		aSensor.UpdatePid(myDB, aPid, 20)
		log.Printf("   err=%v\n", err)

		vid.InsertVideoEnt(myDB,
			aFil,
			aSensor.Sensortype,
			aSensor.ID,
			aSensor.UserID,
			ds.GetNowTime(),
			ds.GetNowTime(),
			aSensor.LocalUrl,
			ds.GetNgArchivePath()+aFil+".mp4",
		)

	} else if mode == "OFF" {
		log.Printf("----- FILE CAPTURE OFF ---------")

		aPid, err := aSensor.GetPidByID(myDB, aSensor.ID)
		log.Printf("   aPid=%d\n", aPid)
		ds.KillWithPid(aPid)
		//ds.SignalWithPid(aPid)
		aSensor.UpdatePid(myDB, aPid, 0)
		log.Printf("   err=%v\n", err)

	} else {
		log.Printf("----- BAD FILE MODE ---------")
	}
}

// ChunkCaptureWithIPSen ...
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func ChunkCaptureWithIPSen(aSensor sen.Sensor, mode string) {
	log.Printf(">===>sencapSensor.ChunkCaptureWithIPSen()\n")

	var urlStr = aSensor.Sensorproto + "://" + aSensor.Sensorauth + "@" + aSensor.LocalUrl

	myDB := db.Connect()

	if mode == "ON" {
		log.Printf("----- CHUNK CAPTURE ON ---------")

		aPid, aFil, err := ChunkCaptureWithFFmpeg(myDB, aSensor, urlStr)
		log.Printf("   PID = %d  aFIL = %s \n", aPid, aFil)
		aSensor.UpdatePid(myDB, aPid, 30)
		log.Printf("   err=%v\n", err)
		log.Printf("   aFil=%v\n", aFil)

	} else if mode == "OFF" {
		log.Printf("----- CHUNK CAPTURE OFF ---------")

		aPid, err := aSensor.GetPidByID(myDB, aSensor.ID)
		log.Printf("   aPid=%d\n", aPid)
		ds.KillWithPid(aPid)
		//ds.SignalWithPid(aPid)
		aSensor.UpdatePid(myDB, aPid, 0)
		log.Printf("   err=%v\n", err)

		//  video, er := CreateVideo(iPid)
		var chunks []chk.Chunk
		//chunks = []chnk.Chunk{ChunkPid: aPid}
		getErr := myDB.Model(&chunks).Where("chunkpid = ?", aPid).Select()
		if getErr != nil {
			// Error Handler
			log.Printf("Error in sencapHandler.ChunkCapture()\n")
			log.Printf("Reason %v\n", getErr)
			return
		}
		log.Printf("Select successful chunkPid\n")
		log.Printf("  len(chunks)=%d\n", len(chunks))

		var filelist string
		var sb strings.Builder
		for i := 0; i < len(chunks); i++ {
			// log.Printf("   %v %v", sensor, sensor.Users[i])
			log.Printf("   %v\n", chunks[i])
			sb.WriteString("file '" + "../Capture/" + chunks[i].ChunkName + "'\n")
		}

		// Need to fix ds.GetNgArchivePath()
		dirPath := ds.GetGoArchivePath()
		fileName := chunks[0].ChunkName[0 : len(chunks[0].ChunkName)-7]
		log.Printf("dirPath=%s, fileName=%s\n", dirPath, fileName)
		filelist = dirPath + fileName + "-list.txt"
		log.Printf("filelist=%s\n", filelist)

		pwd, errrrr := os.Getwd()
		if errrrr != nil {
			fmt.Println(err)
			log.Printf(" ERROR %v", errrrr)
		}
		log.Printf("The current working dir:  %v", pwd)

		errr := ioutil.WriteFile(filelist, []byte(sb.String()), 0644)
		if errr != nil {
			log.Printf("   ERROR ioutil.WriteFile  err: %v\n", errr)
			return
		}

		pid, f, err := ChunkListToVideoWithFFmpeg(fileName, filelist)
		if err != nil {
			log.Printf("   ERROR ChunkListToVideoWithFFmpeg  err: %v\n", err)
			return
		}
		log.Printf("Chunks count to assemble %d", len(chunks))
		if len(chunks) > 0 {
			log.Printf("Chunks to concatinated %s", chunks[0].ChunkName)
			vid.InsertVideoEnt(myDB,
				fileName,
				aSensor.Sensortype,
				aSensor.ID,
				aSensor.UserID,
				ds.GetNowTime(),
				ds.GetNowTime(),
				aSensor.LocalUrl,
				fileName+".mp4",
			)
		}

		log.Printf("Constructed Video from chunks for pid = %d  file = %s ", pid, f)

		//rows, err := strQueryDB.Query("SELECT * FROM entchunks where chunkpid = %d", aPid)
		cmdName = "touch"
		cmdArgs = []string{"../ngAdvDashBoard/src/main.ts"}
		ds.ExecWithScanner(cmdName, cmdArgs)

	} else {
		log.Printf("----- BAD CHUNK MODE ---------")
	}
}

// CaptureWatcher Look for new files created and report.
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func CaptureWatcher(myDB *pg.DB, aSensor sen.Sensor, pid int) {
	log.Printf(">>====>>CaptureWatcher() called\n")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("CaptureWatcher %v \n", err)
	}
	defer watcher.Close()
	log.Printf("CaptureWatcher aSensor = %v PID = %d\n", aSensor, pid)
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Println("event:", event)
				log.Printf("Operation:%s\n", event.Op)
				log.Printf("%s\n", event.Name)
				if event.Op&fsnotify.Write == fsnotify.Write {
					stringSlice := strings.Split(event.Name, "\\")
					//relpath := ds.GetNgCapturePath()
					fileName := stringSlice[len(stringSlice)-1]
					log.Println(" chunk fileName=%s", fileName)
					chk.InsertChunkEnt(myDB, 1,
						aSensor.ID,
						fileName,
						aSensor.Sensortype,
						aSensor.Sensorname,
						fileName, 10, pid)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	path := ds.GetGoCapturePath()
	log.Println("watcher.Add(%s)\n", path)

	err = watcher.Add(path)
	if err != nil {
		log.Printf("watcher.Add %v\n", err)
	}
	<-done
}

/*=====================================================================*\

	THIS CODE DOES NOT SEEM TO BE USED AND SHOULD BE DELETED

\*=====================================================================*/
var strQueryDB qdb.DB

func join_(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pg"
	dbname   = "AdvDashBoard"
)

func sqlConnect_() (*qdb.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//"postgres:pg@tcp(127.0.0.1:5432)/AdvDashBoard"
	db, err := qdb.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Failed to connect to database.\n")
		return nil, err
		//os.Exit(100)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to connect to database.\n")
		return nil, err
	}
	log.Printf("Connection to database [%v] successful.\n", db)
	return db, err
}
