// Package senCapComp is the Sensor Capture Functions
/*===================================================================================*\



\*===================================================================================*/
package senCapComp

import (
	ds "AdvDashBoard/goAdvDashBoard/datastore/fileStore"
	sen "AdvDashBoard/goAdvDashBoard/tablecrud/sensorCrud"

	"log"

	"github.com/go-pg/pg"
)

var (
	cmdUUID string
	cmdName string
	cmdArgs []string
)

// LiveCaptureWithFFmpeg takes a URL and changes the size
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func LiveCaptureWithFFmpeg(sensorName string, sensorType string, urlStr string) (int, error) {
	log.Printf(">===>sencapFFmpeg.LiveCaptureWithFFmpeg(sensorType=%s)", sensorType)
	log.Printf("   urlStr=%s)", urlStr)

	if sensorType == "Avigilon" {
		urlStr = urlStr + "/defaultPrimary?streamType=u"
		log.Printf("   urlStr=%s)", urlStr)

		cmdName = ds.GetGoToolPath() + "ffplay"
		cmdArgs = []string{
			"-fflags", "nobuffer",
			"-x", "640", "-y", "480",
			urlStr,
		}

		pid, err := ds.ExecReturnPid(cmdName, cmdArgs)
		return pid, err

	} else if sensorType == "SV3C" {
		urlStr = urlStr + ":554/11"
		log.Printf("   urlStr=%s)", urlStr)

		cmdName = ds.GetGoToolPath() + "ffplay"
		cmdArgs = []string{
			"-fflags", "nobuffer",
			"-x", "640", "-y", "480",
			urlStr,
		}

		pid, err := ds.ExecReturnPid(cmdName, cmdArgs)
		return pid, err
	} else if sensorType == "Webcam" {
		urlStr = urlStr
		log.Printf("   urlStr=%s)", urlStr)

		cmdName = ds.GetGoToolPath() + "ffplay"
		cmdArgs = []string{"-hide_banner",
			"-f", "dshow",
			// "-i", "video=Logitech HD Webcam C615",
			// "-i", "video=HP TrueVision HD",
			"-i", "video=" + sensorName,
		}

		pid, err := ds.ExecReturnPid(cmdName, cmdArgs)
		return pid, err
	} else if sensorType == "R2D2" {

	} else if sensorType == "CudaEye" {

	} else {
		log.Printf("BAD sensorType=%s", sensorType)
	}
	return 0, nil
}

// FileCaptureWithFFmpeg takes a cameraName and urlStr
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func FileCaptureWithFFmpeg(sensorName string, sensorType string, urlStr string) (int, string, error) {
	log.Printf(">===>sencapFFmpeg.FileCaptureWithFFmpeg(sensorType=%s)", sensorType)
	log.Printf("   urlStr=%s)", urlStr)

	fileName := ds.GenBetterGUID()
	cmdName = "ffmpeg"
	if sensorType == "Avigilon" {
		urlStr = urlStr + "/defaultPrimary?streamType=u"
		log.Printf("   urlStr=%s)", urlStr)
		cmdArgs = []string{
			"-i", urlStr,
			"-max_muxing_queue_size", "8192",
			"-sn",
			"-vcodec", "copy",
			"-r", "25", "-t", "30",
			"-y", ds.GetGoArchivePath() + fileName + ".mp4",
		}
		pid, err := ds.ExecReturnPid(cmdName, cmdArgs)

		return pid, fileName, err
	} else if sensorType == "SV3C" {
		urlStr = urlStr + ":554/11"
		log.Printf("   urlStr=%s)", urlStr)
		cmdArgs = []string{
			"-i", urlStr,
			"-max_muxing_queue_size", "8192",
			"-sn",
			"-vcodec", "copy",
			"-r", "25", "-t", "30",
			"-y", ds.GetGoArchivePath() + fileName + ".mp4",
		}
		pid, err := ds.ExecReturnPid(cmdName, cmdArgs)

		return pid, fileName, err
	} else if sensorType == "Webcam" {
		urlStr = urlStr
		log.Printf("   urlStr=%s)", urlStr)
		//Most recent change
		cmdArgs = []string{
			// "-i", "video=Logitech HD Webcam C615",
			// "-i", "video=HP TrueVision HD",
			"-i", "video=" + sensorName,
			"-max_muxing_queue_size", "8192",
			"-sn",
			"-vcodec", "copy",
			"-r", "25", "-t", "30",
			"-y", ds.GetGoArchivePath() + fileName + ".mp4",
		}
		pid, err := ds.ExecReturnPid(cmdName, cmdArgs)

		return pid, fileName, err
	} else if sensorType == "R2D2" {

	} else if sensorType == "CudaEye" {

	} else {
		log.Printf("BAD sensorType=%s", sensorType)
	}
	return 0, fileName, nil
}

// ChunkCaptureWithFFmpeg takes a cameraName and urlStr
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func ChunkCaptureWithFFmpeg(myDB *pg.DB, aSensor sen.Sensor, urlStr string) (int, string, error) {
	log.Printf(">===>sencapFFmpeg.ChunkCaptureWithFFmpeg(sensorType=%v)", aSensor)
	log.Printf("   urlStr=%s)", urlStr)
	cmdUUID = ds.GenBetterGUID()
	// chk.InsertChunkEnt(
	// 	myDB, 2, 2, cmdUUID,
	// 	aSensor.Sensortype, aSensor.Sensorname,
	// 	ds.GetNgArchivePath(), 1, 1234)

	cmdName = ds.GetGoToolPath() + "ffmpeg"
	if aSensor.Sensortype == "Avigilon" {
		urlStr = urlStr + "/defaultPrimary?streamType=u"
		log.Printf("   urlStr=%s)", urlStr)
		cmdArgs = []string{
			"-hide_banner",
			"-i", urlStr,
			"-max_muxing_queue_size", "8182",
			"-vcodec", "copy", //"-acodec", "libfdk_aac",
			"-an", "-sn", //"-c:v",
			// "copy", "-b:v", "128k",
			//	"-map", "0",
			"-threads", "4",
			"-f", "segment",
			"-segment_time", "5",
			"-segment_format", "mp4",
			ds.GetGoCapturePath() + cmdUUID + "%03d.mp4",
		}

	} else if aSensor.Sensortype == "SV3C" {
		urlStr = urlStr + ":554/11"
		log.Printf("   urlStr=%s)", urlStr)

		cmdArgs = []string{
			"-hide_banner",
			"-i", urlStr,
			"-max_muxing_queue_size", "8182",
			"-vcodec", "copy", //"-acodec", "libfdk_aac",
			"-an", "-sn", //"-c:v",
			// "copy", "-b:v", "128k",
			//	"-map", "0",
			"-threads", "4",
			"-f", "segment",
			"-segment_time", "5",
			"-segment_format", "mp4",
			ds.GetGoCapturePath() + cmdUUID + "%03d.mp4",
		}

	} else if aSensor.Sensortype == "Webcam" {
		urlStr = urlStr
		log.Printf("   urlStr=%s)", urlStr)

		cmdArgs = []string{
			"-hide_banner",
			// "-i", "video=Logitech Webcam HD C615",
			// "-i", "video=HP TrueVision HD",
			"-i", "video=" + aSensor.Sensorname,
			"-max_muxing_queue_size", "8182",
			"-vcodec", "copy", //"-acodec", "libfdk_aac",
			"-an", "-sn", //"-c:v",
			// "copy", "-b:v", "128k",
			//	"-map", "0",
			"-threads", "4",
			"-f", "segment",
			"-segment_time", "5",
			"-segment_format", "mp4",
			ds.GetGoCapturePath() + cmdUUID + "%03d.mp4",
		}

	} else if aSensor.Sensortype == "R2D2" {

	} else if aSensor.Sensortype == "CudaEye" {

	} else {
		log.Printf("BAD sensorType=%s", aSensor.Sensortype)
	}

	pid, err := ds.ExecReturnPid(cmdName, cmdArgs)
	if err == nil {
		go CaptureWatcher(myDB, aSensor, pid)
		return pid, cmdUUID, err
	} else {
		log.Printf("Error with NOT starting Watcher\n")
		log.Printf("Reason:%v\n", err)
		return pid, cmdUUID, err
	}
	return 0, cmdUUID, nil
}

// ChunkListToVideoWithFFmpeg creates a video from chunks ident by GUID
/*-----------------------------------------------------------------------------------*\
\*-----------------------------------------------------------------------------------*/
func ChunkListToVideoWithFFmpeg(filename string, tempfile string) (int, string, error) {
	log.Printf(">===>sencapSensor.ChunkListToVideoWithFFmpeg\n")
	// log.Printf("   ds.GetGoVideoPath() + filenam GUID=%s \n", ds.GetGoVideoPath()+filename+".mp4")
	log.Printf("   tempfile=%s \n", tempfile)
	// ffmpeg -f concat -safe 0 -i mylist.txt -c copy output.mp4

	cmdName = ds.GetGoToolPath() + "ffmpeg"
	cmdArgs = []string{
		"-f", "concat",
		"-safe", "0",
		"-i",
		tempfile, "-c",
		//"copy", "c:/VIT.DEV/Proto/src/AdvDashBoard/ngAdvDashBoard/src/assets/Archive/" + filename + ".mp4",
		"copy", ds.GetGoArchivePath() + filename + ".mp4",
	}
	pid, err := ds.ExecReturnPid(cmdName, cmdArgs)
	if err != nil {
		log.Printf("   ERROR filename GUID=%s  err: %v\n", filename, err)
		return 0, filename, err
	}
	return pid, filename, nil

}
