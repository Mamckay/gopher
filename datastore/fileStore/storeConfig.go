// Package fileStore is the Data Storage Config Package
/*===================================================================================*\
	The fileStore is the Sensor Capture Component


\*===================================================================================*/
package fileStore

import (
	"fmt"
	"log"
	"time"
)

const (
	winToolPath = "c:/Appl/ffmpeg/bin/"
	macToolPath = "/usr/local/bin/"

	winDevName = "c:/Users/amckay/"
	macDevName = "/Users/amckay/"
	videoFile  = "Snowing"
)

var (
	winStorePath = "c:/VIT.DEV/Proto/src/AdvDashBoard/ngAdvDashBoard/src/assets/"
	macStorePath = macDevName + "VIT.DEV/Proto/src/AdvDashBoard/ngAdvDashBoard/src/assets/"
	storePath    string
	toolPath     string
	chunkPath    string
	archivePath  string
	productPath  string
	// videoPath    string
	// tempPath     string
)

func GetNowTime() string {
	timeNow := time.Now()
	//log.Printf("timeNow=%s\n", timeNow)
	timeStr := fmt.Sprintf("%s", timeNow)
	//log.Printf("timeStr=%s\n", timeStr[0:19])
	return timeStr[0:19]
}

// SetMachine Takes a machine name and configures cmdFFmpeg component
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func SetMachine(machine string) {
	log.Printf(">===>VidCapturePkg.SetMachine(%s)", machine)

	if machine == "mac" {
		storePath = macStorePath
		toolPath = macToolPath
	} else if machine == "win" {
		storePath = winStorePath
		toolPath = winToolPath
		chunkPath = storePath + "Capture/"
		archivePath = storePath + "Archive/"
		productPath = storePath + "Products/"
		// videoPath = storePath + "Video/"
		// tempPath = storePath + "Temp/"

	} else if machine == "lin" {

	} else {
		log.Printf("BAD MACHINE !")
	}
}

/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
// GetGoPaths Returns the Tools Path and with Store Path
func GetGoPaths(machine string) (string, string) {
	return toolPath, storePath
}

// GetGoToolPath Returns the path to FFMPEG tools
func GetGoToolPath() string {
	return toolPath
}

// GetGoStorePath Returns the path to the video storage
func GetGoStorePath() string {
	return storePath
}

// GetGoCapturePath Returns the path to the video chunk storage
func GetGoCapturePath() string {
	return chunkPath
}

// GetGoProductPath Returns the path to the video chunk storage
func GetGoProductPath() string {
	return productPath
}

// GetGoArchivePath Returns the path to the video archive storage
func GetGoArchivePath() string {
	return archivePath
}

// GetNgCapturePath Returns the path to the video chunk storage
func GetNgCapturePath() string {
	chunkPath = "../../assets/" + "Capture/"
	return chunkPath
}

// GetNgProductPath Returns the path to the video chunk storage
func GetNgProductPath() string {
	productPath = "../../assets/" + "Products/"
	return productPath
}

// GetNgArchivePath Returns the path to the video archive storage
func GetNgArchivePath() string {
	archivePath = "../../assets/" + "Archive/"
	return archivePath
}
