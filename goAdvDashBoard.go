/*=====================================================================*\


\*=====================================================================*/
package main

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	ds "AdvDashBoard/goAdvDashBoard/datastore/fileStore"
	cap "AdvDashBoard/goAdvDashBoard/rolefunc/sensorCapFunc"
	chk "AdvDashBoard/goAdvDashBoard/tablecrud/chunkCrud"
	prd "AdvDashBoard/goAdvDashBoard/tablecrud/productCrud"
	sen "AdvDashBoard/goAdvDashBoard/tablecrud/sensorCrud"
	usr "AdvDashBoard/goAdvDashBoard/tablecrud/userCrud"
	vid "AdvDashBoard/goAdvDashBoard/tablecrud/videoCrud"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func usage() {
	log.Printf("Usage: goAdvDashBoard [--machine <Machine>] [--test <TestName>] [--database <DataBaseType>] [--sensor <SensorName>]")
	log.Printf("	Where: <Machine> = {Local | DemoVm-2 | DemoVM-5 }")
	log.Printf("	Where: <DataBaseType> = {Empty | Starter | Simple | Complex }")
	log.Printf("	Where: <SensorName> = {SV3C | Avigilon | CudaEye | R2D2 }")
	log.Printf("	Where: <TestName> = {Live | File | Chunks }")
}

var hostAddr string
var hostPort string
var muxPort string
var hostURL string
var machine string
var testName string
var dataBase string
var sensorName string

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func main() {
	// if len(os.Args) < 3 {
	// 	usage()
	// 	return
	// }
	log.Println(">===>goPgMultComp.main()")

	machine = "local"
	testName = "None"
	dataBase = "None"
	sensorName = "None"

	idx := 1
	for _, a := range os.Args[1:] {
		fmt.Printf("   %d.%s\n", idx, a)
		if a == "--test" {
			testName = os.Args[idx+1]
		} else if a == "--machine" {
			machine = os.Args[idx+1]
		} else if a == "--sensor" {
			sensorName = os.Args[idx+1]
		} else if a == "--database" {
			dataBase = os.Args[idx+1]
		} else if a == "--help" {
			usage()
			return
		}
		idx = idx + 1
	}

	SetGoMachine(machine)
	SetMuxPortAddr(muxPort)

	log.Println("   muxPort=", muxPort)
	log.Println("   hostAddr=", hostAddr)
	log.Println("   hostPort=", hostPort)
	log.Println("   hostURL=", hostURL)
	log.Println("   testName=", testName)
	log.Println("   sensorName=", sensorName)
	log.Println("   dataBase=", dataBase)

	//return
	CreateOptionDataBase(dataBase)
	ExecuteOptionalTest(testName)

	if testName == "None" {
		SensorCaptureRouter(hostURL)
	} else {
	}
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func SetGoMachine(machine string) {
	if machine == "local" {
		muxPort = ":8062"
		hostAddr = "localhost"
		hostPort = "6200"
	} else if machine == "DevOps" {
		muxPort = ":8000"
		hostAddr = "18.191.141.108"
		hostPort = "4200"
	} else if machine == "Demo" {
		muxPort = ":8000"
		hostAddr = "18.218.204.94"
		hostPort = "4200"
	} else {
		log.Println("BAD machine!")
	}

	hostURL = "http://" + hostAddr + ":" + hostPort

}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func ExecuteOptionalTest(testName string) {
	// var urlStr = ""
	// if sensorName == "Avigilon" {
	// 	urlStr = "rtsp://admin:admin@192.168.0.2/defaultPrimary?streamType=u"
	// } else if sensorName == "SV3C" {

	// } else if sensorName == "R2D2" {

	// } else if sensorName == "CudaEye" {

	// } else {
	// 	log.Printf("BAD sensorName=%s", sensorName)
	// 	usage()
	// 	return
	// }

	// if testName == "Live" {
	// 	//cap.LiveCaptureWithIPSen(sensorName, urlStr)
	// } else if testName == "File" {
	// 	//cap.FileCaptureWithIPSen(sensorName, urlStr)
	// } else if testName == "Chunks" {
	// 	// dbRef := db.Connect()
	// 	// cap.ChunkCaptureWithIPSen(dbRef, sensorName, urlStr)
	// } else {
	// 	log.Printf("BAD testName=%s", testName)
	// 	usage()
	// 	return
	// }
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateOptionDataBase(dataBase string) {
	if dataBase == "Empty" {
		CreatEmptyTables()
		return
	} else if dataBase == "Starter" {
		CreatEmptyTables()
		CreatStarterEntities()
		return
	} else if dataBase == "Simple" {
		CreatEmptyTables()
		//CreatStarterEntities()
		LoadSimpleDatabase()
		return
	} else if dataBase == "Complex" {
		CreatEmptyTables()
		CreatStarterEntities()
		LoadSimpleDatabase()
		return
	}
}

// SensorCaptureRouter ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func SensorCaptureRouter(hostAddr string) {
	log.Println(">===>goAdvDashBoard.SensorCaptureRouter()")

	SetHost(hostAddr)
	ds.SetMachine("win")

	// Init the Router
	r := mux.NewRouter().StrictSlash(true)

	// Route Handlers for Sensor Capture Endpoints
	r.HandleFunc("/api/SenCap/Live/{mode}", cap.LiveCapture).Methods("POST")
	r.HandleFunc("/api/SenCap/File/{mode}", cap.FileCapture).Methods("POST")
	r.HandleFunc("/api/SenCap/Chunk/{mode}", cap.ChunkCapture).Methods("POST")

	r.HandleFunc("/api/Users/ToSensor", sen.HandleUserToSensorRel).Methods("POST")
	r.HandleFunc("/api/Users/ToVideo", vid.HandleUserToVideoRel).Methods("POST")
	r.HandleFunc("/api/Users/ToProduct", prd.HandleUserToProductRel).Methods("POST")

	// Advanced Relationship endpoints
	r.HandleFunc("/api/Products/ByChunkID/{id}", chk.ProductsByChunkID).Methods("GET")
	r.HandleFunc("/api/Videos/ByChunkID/{id}", chk.VideosByChunkID).Methods("GET")

	// Route Handlers for User CRUD Endpoints
	r.HandleFunc("/api/Users/auth", usr.AuthUser).Methods("GET")
	r.HandleFunc("/api/Users/register", usr.RegisterUser).Methods("POST")
	r.HandleFunc("/api/Users/authenticate", usr.GetAuthUser).Methods("PUT")
	r.HandleFunc("/status", usr.StatusHandler).Methods("GET")
	r.HandleFunc("/api/Users", usr.GetUsersFullList).Methods("GET")
	r.HandleFunc("/api/Users", usr.CreateUser).Methods("POST")
	r.HandleFunc("/api/Users/{id}", usr.GetUser).Methods("GET")
	r.HandleFunc("/api/Users/{id}", usr.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/Users/{id}", usr.DeleteUser).Methods("DELETE")

	// Advanced Relationship endpoints
	r.HandleFunc("/api/Users/BySensorID/{id}", sen.GetUsersBySensorID).Methods("GET")
	r.HandleFunc("/api/Users/NotBySensorID/{id}", sen.GetNotUsersBySensorID).Methods("GET")
	r.HandleFunc("/api/Users/ByVideoID/{id}", vid.GetUsersByVideoID).Methods("GET")
	r.HandleFunc("/api/Users/NotByVideoID/{id}", vid.GetNotUsersByVideoID).Methods("GET")
	r.HandleFunc("/api/Users/ByProductID/{id}", prd.GetUsersByProductID).Methods("GET")
	r.HandleFunc("/api/Users/NotByProductID/{id}", prd.GetNotUsersByProductID).Methods("GET")

	// Route Handlers for Sensor CRUD Endpoints
	r.HandleFunc("/api/Sensors", sen.CreateSensor).Methods("POST")
	r.HandleFunc("/api/Sensors", sen.GetSensorsFullList).Methods("GET")
	r.HandleFunc("/api/Sensors/{id}", sen.GetSensor).Methods("GET")
	r.HandleFunc("/api/Sensors/{id}", sen.UpdateSensor).Methods("PUT")
	r.HandleFunc("/api/Sensors/{id}", sen.DeleteSensor).Methods("DELETE")
	r.HandleFunc("/api/Sensors/ByOwnerID/{id}", sen.GetSensorsByOwnerID).Methods("GET")
	r.HandleFunc("/api/Sensors/ByUserID/{id}", sen.GetSensorsByUserID).Methods("GET")

	// Route Handlers for Product CRUD Endpoints
	r.HandleFunc("/api/Products", prd.CreateProduct).Methods("POST")
	r.HandleFunc("/api/Products", prd.GetProductsFullList).Methods("GET")
	r.HandleFunc("/api/Products/{id}", prd.GetProduct).Methods("GET")
	r.HandleFunc("/api/Products/{id}", prd.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/Products/{id}", prd.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/Products/ByOwnerID/{id}", prd.GetProductsByOwnerID).Methods("GET")
	r.HandleFunc("/api/Products/ByUserID/{id}", prd.GetProductsByUserID).Methods("GET")

	// Route Handlers for Chunk CRUD Endpoints
	r.HandleFunc("/api/Chunks", chk.CreateChunk).Methods("POST")
	r.HandleFunc("/api/Chunks", chk.GetChunksFullList).Methods("GET")
	r.HandleFunc("/api/Chunks/{id}", chk.GetChunk).Methods("GET")
	r.HandleFunc("/api/Chunks/{id}", chk.UpdateChunk).Methods("PUT")
	r.HandleFunc("/api/Chunks/{id}", chk.DeleteChunk).Methods("DELETE")
	r.HandleFunc("/api/Chunks/ByOwnerID/{id}", chk.GetChunksByOwnerID).Methods("GET")
	r.HandleFunc("/api/Chunks/ByUserID/{id}", chk.GetChunksByUserID).Methods("GET")

	// Route Handlers for Video CRUD Endpoints
	r.HandleFunc("/api/Videos", vid.CreateVideo).Methods("POST")
	r.HandleFunc("/api/Videos", vid.GetVideosFullList).Methods("GET")
	r.HandleFunc("/api/Videos/{id}", vid.GetVideo).Methods("GET")
	r.HandleFunc("/api/Videos/{id}", vid.UpdateVideo).Methods("PUT")
	r.HandleFunc("/api/Videos/{id}", vid.DeleteVideo).Methods("DELETE")
	r.HandleFunc("/api/Videos/ByOwnerID/{id}", vid.GetVideosByOwnerID).Methods("GET")
	r.HandleFunc("/api/Videos/ByUserID/{id}", vid.GetVideosByUserID).Methods("GET")

	var c = cors.AllowAll()
	handler := c.Handler(r)

	log.Println("	vid.ListenAndServe on [" + muxRouterPort + "]")
	log.Fatal(http.ListenAndServe(muxRouterPort, handler))
}

var (
	hostName string
)

// SetHost ...
func SetHost(host string) { hostName = host }

var (
	muxRouterPort string
)

// SetMuxPortAddr ...
func SetMuxPortAddr(addr string) {
	muxRouterPort = addr
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreatEmptyTables() {
	log.Printf("\n")
	log.Printf("================================================\n")
	log.Printf("    CreateEmptyTables()\n")
	log.Printf("================================================\n")

	chk.CreateChunkTable()
	prd.CreateProductTable()
	sen.CreateSensorTable()
	usr.CreateUserTable()
	vid.CreateVideoTable()

	chk.CreateChunkToVideoTable()
	chk.CreateChunkToProductTable()

	vid.CreateUserToVideoTable()
	prd.CreateUserToProductTable()
	sen.CreateUserToSensorTable()

	vid.CreateVideoToUserTable()
	prd.CreateProductToUserTable()
	sen.CreateSensorToUserTable()
}

/*---------------------------------------------------------------------*\
[]-Create empty Entity and Relation Tables
[]-Create 4 users with Roles: Admin-1, Operator-2, Reviewer-3, Creator-4
[]-Create a sensor-1 owned by the Operator-2 user
[]-Create chunks-1,2 captured on sensor-1
[]-Create a video-1,2,3 created on the sensor-1 with at least one Reviewer-3
[]-Create a product-1 owned by Creator-4 that contains the chunks-1,2,3
[]-Create a video-3 owned by Operator-3 that contains the chunks-1,2,3
[]-Make Product-1 the owner of Chunks-1,2,3
[]-Make Video-3 the owner of Chunks-1,2,3
[]-Make Reviewer-3 have access to Video-1, and Product-1
\*---------------------------------------------------------------------*/
func LoadSimpleDatabase() {
	log.Printf("\n")
	log.Printf("================================================\n")
	log.Printf("    LoadSimpleDatabase()\n")
	log.Printf("================================================\n")
	dbRef := db.Connect()

	usr.InsertUserEnt(dbRef, "AdminUser", "123456", "John", "Doe", "Admin", "Token")
	usr.InsertUserEnt(dbRef, "OperateUser", "123456", "Jean", "Smith", "Operator", "Token")
	usr.InsertUserEnt(dbRef, "ReviewUser", "123456", "Mike", "Johnson", "Viewer", "Token")
	usr.InsertUserEnt(dbRef, "ProdUser", "123456", "Kim", "Hann", "Creator", "Token")
	usr.InsertUserEnt(dbRef, "AdmSki", "123456", "Adm", "Ski", "Operator", "Token")
	usr.InsertUserEnt(dbRef, "JcrSki", "123456", "Scr", "Ski", "Reviewer", "Token")

	//-Create a sensor-1 owned by user-2, which has operator role
	sen.InsertSensorEnt(dbRef, 3, "HP TrueVision HD", "Webcam", "Live", "Bi7", "SensorOwner", "rtsp", "admin:admin", "n/a", "High", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "HP TrueVision HD", "Webcam", "File", "Bi7", "SensorOwner", "rtsp", "admin:admin", "n/a", "Medium", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "HP TrueVision HD", "Webcam", "Chunk", "Bi7", "SensorOwner", "rtsp", "admin:admin", "n/a", "Low", 1, 1)

	sen.InsertSensorEnt(dbRef, 3, "Integrated Webcam", "Webcam", "Live", "VIT", "SensorOwner", "rtsp", "admin:admin", "n/a", "High", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "Integrated Webcam", "Webcam", "File", "VIT", "SensorOwner", "rtsp", "admin:admin", "n/a", "Medium", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "Integrated Webcam", "Webcam", "Chunk", "VIT", "SensorOwner", "rtsp", "admin:admin", "n/a", "Low", 1, 1)

	sen.InsertSensorEnt(dbRef, 2, "3EA1", "Avigilon", "Live", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.0.2", "High", 1, 1)
	sen.InsertSensorEnt(dbRef, 2, "3EA1", "Avigilon", "File", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.0.2", "Medium", 1, 1)
	sen.InsertSensorEnt(dbRef, 2, "3EA1", "Avigilon", "Chunk", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.0.2", "Low", 1, 1)

	sen.InsertSensorEnt(dbRef, 2, "A5B3", "SV3C", "Live", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.0.21", "High", 1, 1)
	sen.InsertSensorEnt(dbRef, 2, "A5B3", "SV3C", "File", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.0.21", "Medium", 1, 1)
	sen.InsertSensorEnt(dbRef, 2, "A5B3", "SV3C", "Chunk", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.0.21", "Low", 1, 1)

	sen.InsertSensorEnt(dbRef, 3, "Integrated Webcam", "Webcam", "Live", "JCR", "SensorOwner", "rtsp", "admin:admin", "n/a", "High", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "Integrated Webcam", "Webcam", "File", "JCR", "SensorOwner", "rtsp", "admin:admin", "n/a", "Medium", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "Integrated Webcam", "Webcam", "Chunk", "JCR", "SensorOwner", "rtsp", "admin:admin", "n/a", "Low", 1, 1)

	sen.InsertSensorEnt(dbRef, 2, "GreenSensor", "SV3C", "Live", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.1.11", "High", 1, 1)
	sen.InsertSensorEnt(dbRef, 2, "BlueSensor", "SV3C", "File", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.1.11", "Medium", 1, 1)
	sen.InsertSensorEnt(dbRef, 2, "RedSensor", "SV3C", "Chunk", "Location", "SensorOwner", "rtsp", "admin:admin", "192.168.1.11", "Low", 1, 1)

	//-Create a video-1,2 created on the sensor-1 owned bu Operator-2
	vid.InsertVideoEnt(dbRef, "Melonga Dancers", "Avigilon", 1, 2, "2018/07/24 11:18:49", "2018/07/24 12:18:49 ", "198.162.0.10", "Melonga_160x120.mp4")
	vid.InsertVideoEnt(dbRef, "Tango Dancers", "Avigilon", 1, 2, "2018/07/24 09:18:49", "2018/07/24 10:18:49 ", "198.162.0.10", "Tango_160x120.mp4")
	vid.InsertVideoEnt(dbRef, "Test Video 3", "SV3C", 1, 2, "2018/07/24 09:18:49", "2018/07/24 10:18:49 ", "198.162.0.10", "Tango_160x120.mp4")
	vid.InsertVideoEnt(dbRef, "Test Video 4", "ABC", 1, 2, "2018/07/24 09:18:49", "2018/07/24 10:18:49 ", "198.162.0.10", "Tango_160x120.mp4")
	vid.InsertVideoEnt(dbRef, "Test Video 5", "DEF", 1, 2, "2018/07/24 09:18:49", "2018/07/24 10:18:49 ", "198.162.0.10", "Tango_160x120.mp4")

	//-Create chunks-1,2,3 captured on sensor-1
	chk.InsertChunkEnt(dbRef, 2, 2, "Chunkone", "Avigilon", "OperatorUser", "Melonga_160x120.mp4", 1, 1234)
	// chk.InsertChunkEnt(dbRef, 2, 2, "Chunk Two", "Avigilon", "OperatorUser", "Melonga_160x120.mp4", 1, 1234)
	// chk.InsertChunkEnt(dbRef, 2, 2, "Chunk Three", "Avigilon", "OperatorUser", "Melonga_160x120.mp4", 1, 1234)

	//-Create a product-1 owned by Creator-4 that will contain the chunks-1,2,3
	prd.InsertProductEnt(dbRef, "3Chk-Product", "Avigilon", 2, 2, "ProdUser", "Melonga_160x120.mp4", 1, 1234)
	prd.InsertProductEnt(dbRef, "4Chk-Product", "SV3C", 2, 2, "ProdUser", "Melonga_160x120.mp4", 1, 4567)
	prd.InsertProductEnt(dbRef, "5Chk-Product", "SV3C", 2, 2, "ProdUser", "Melonga_160x120.mp4", 1, 4535)

	//-Create a product-1 owned by 4-Creator that will contain the chunks-1,2,3
	vid.InsertVideoEnt(dbRef, "3Chk-Video", "Avigilon", 1, 2, "2018/07/24 09:18:49", "2018/07/24 10:18:49 ", "198.162.0.10", "Tango_160x120.mp4")

	//-Make Chunks 1,2,3 be contained within Product 1
	chk.InsertChunkToProductRel(dbRef, 1, 1)
	chk.InsertChunkToProductRel(dbRef, 1, 2)
	chk.InsertChunkToProductRel(dbRef, 1, 3)

	//-Make Chunks 1,2,3 be contained within Video 3
	chk.InsertChunkToVideoRel(dbRef, 2, 1)
	chk.InsertChunkToVideoRel(dbRef, 2, 2)
	chk.InsertChunkToVideoRel(dbRef, 2, 3)
	chk.InsertChunkToVideoRel(dbRef, 2, 4)

	// Make user-3 with role reviewer authorized to see video 1, and Product 1
	vid.InsertUserToVideoRel(dbRef, 2, 1)
	vid.InsertUserToVideoRel(dbRef, 2, 2)
	vid.InsertUserToVideoRel(dbRef, 3, 2)
	vid.InsertUserToVideoRel(dbRef, 5, 1)
	vid.InsertUserToVideoRel(dbRef, 5, 2)

	// prd.InsertUserToProductRel_NOTUSED(dbRef, 2, 1)
	// prd.InsertUserToProductRel_NOTUSED(dbRef, 2, 2)
	// prd.InsertUserToProductRel_NOTUSED(dbRef, 3, 2)
	prd.InsertUserToProduct(dbRef, 2, 1)
	prd.InsertUserToProduct(dbRef, 2, 2)
	prd.InsertUserToProduct(dbRef, 3, 2)
	prd.InsertUserToProduct(dbRef, 5, 2)

	sen.InsertUserToSensorRel(dbRef, 2, 1)
	sen.InsertUserToSensorRel(dbRef, 2, 2)
	sen.InsertUserToSensorRel(dbRef, 3, 2)
	sen.InsertUserToSensorRel(dbRef, 5, 2)
}

func LoadComplexDatabase() {
	log.Printf("\n")
	log.Printf("================================================\n")
	log.Printf("    LoadComplexDatabase()\n")
	log.Printf("================================================\n")
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreatStarterEntities() {
	log.Printf("\n")
	log.Printf("================================================\n")
	log.Printf("    CreatStarterEntities()\n")
	log.Printf("================================================\n")
	dbRef := db.Connect()
	usr.InsertUserEnt(dbRef, "AdminUser", "123456", "John", "Doe", "Admin", "Token")
	chk.InsertChunkEnt(dbRef, 2, 2, "ChunkOne", "S3VC", "SomeOwner", "Melonga_160x120.mp4", 1, 1234)
	prd.InsertProductEnt(dbRef, "ProductOne", "S3VC", 2, 2, "SomeOwner", "ProductOne", 1, 1234)
	sen.InsertSensorEnt(dbRef, 3, "SensorName", "Avigilon", "Live", "Location", "SensorOwner", "rtsp", "192.168.0.2", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "SensorName", "Avigilon", "File", "Location", "SensorOwner", "rtsp", "192.168.0.2", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "SensorName", "Avigilon", "Chunk", "Location", "SensorOwner", "rtsp", "192.168.0.2", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "SensorName", "S3VC", "Live", "Location", "SensorOwner", "rtsp", "198.168.1.10", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "SensorName", "S3VC", "File", "Location", "SensorOwner", "rtsp", "198.168.1.10", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 3, "SensorName", "S3VC", "Chunk", "Location", "SensorOwner", "rtsp", "198.168.1.10", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 5, "Adm Green", "S3VC", "Live", "Location", "SensorOwner", "rtsp", "192.168.0.2", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 5, "Adm Blue", "S3VC", "File", "Location", "SensorOwner", "rtsp", "192.168.0.2", "InputURL", "Security", 1, 1)
	sen.InsertSensorEnt(dbRef, 5, "Adm Red", "S3VC", "Chunk", "Location", "SensorOwner", "rtsp", "192.168.0.2", "InputURL", "Security", 1, 1)
	vid.InsertVideoEnt(dbRef, "Melonga Dancers", "Body Cam", 1, 1, "2018/07/24 11:18:49", "2018/07/24 12:18:49 ", "198.162.0.10", "Melonga_160x120.mp4")
}

/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreatInitialRoles() {
	log.Printf("\n")
	log.Printf("================================================\n")
	log.Printf("    CreatInitialRoles()\n")
	log.Printf("================================================\n")
	dbRef := db.Connect()
	usr.InsertUserEnt(dbRef, "AdminUser", "123456", "John", "Doe", "Admin", "Token")
	usr.InsertUserEnt(dbRef, "OperateUser", "123456", "Jean", "Smith", "Operator", "Token")
	usr.InsertUserEnt(dbRef, "ReviewUser", "123456", "Mike", "Johnson", "Viewer", "Token")
	usr.InsertUserEnt(dbRef, "ProdUser", "123456", "Kim", "Hann", "Creator", "Token")
	usr.InsertUserEnt(dbRef, "AdmSki", "123456", "Adm", "Ski", "Operator", "Token")
	usr.InsertUserEnt(dbRef, "JcrSki", "123456", "Scr", "Ski", "Reviewer", "Token")
}
