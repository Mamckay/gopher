/*=====================================================================*\


\*=====================================================================*/
package videoComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
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
func TestGetResultsByUserID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/videoCrud.TestGetResultsByUserID()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myVideo := &Video{}
	videos, err := myVideo.GetResultsByUserID(myDB, 2)

	if err != nil {
		log.Printf("Test Failed %s", err)
		return
	}
	log.Printf("Select successful for ID: %v\n", videos)

	log.Printf("videos found inside sensorModel.GetResultsByUserID()")
	for i := 0; i < len(videos); i++ {
		log.Printf("   " + videos[i].VideoName)
	}

}

// TestGetUserById ...
func TestGetVideoById(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/userCrud.TestGetUserById()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Video{ID: 1}
	myUser.GetByID(myDB)
}

// TestGetUserByName ...
func TestGetVideoByName(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/userCrud.TestGetUserByName()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Video{VideoName: "outside1"}
	myUser.GetByName(myDB)
}

func TestDeleteVideoWithId(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/userCrud.TestDeleteUserWithId()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Video{ID: 3}
	myUser.Delete(myDB)
}

func TestDeleteVideoWithName(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/userCrud.TestDeleteUserWithName()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Video{VideoName: "front"}
	myUser.Delete(myDB)
}

/*=====================================================================*\
	Functions below this point have been prepared for integrations

\*=====================================================================*/

// Test GetUsersByAccessID
func TestGetUsersByAccessID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Video{}
	videoUsersAccess, errTest := myUser.GetUsersByAccessID(myDB, 2)
	if errTest != nil {
		log.Printf("   Test Failed in TestGetUsersBySensorID: %v", errTest)
		log.Printf("   %v", videoUsersAccess)
		return
	}
	log.Printf("%d users found inside TestGetUsersByAccessID()", len(videoUsersAccess))
	for i := 0; i < len(videoUsersAccess); i++ {
		log.Printf("   %v\n", videoUsersAccess[i])
	}

}

// TestGetUsersByVideoAccess ...
func TestGetVideosByAccessUserID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	trace2()
	fmt.Println(nameOf((*A).Method))
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myUser := &Video{}
	userVideoAccess, errTest := myUser.GetVideosByAccessUserID(myDB, 2)
	if errTest != nil {
		log.Printf("   Test Failed in TestGetUsersBySensorID: %v", errTest)
		log.Printf("   %v", userVideoAccess)
		return
	}
	fmt.Println("User ", userVideoAccess.ID)
	for i := 0; i < len(userVideoAccess.Videos); i++ {
		// log.Printf("   %v %v", sensor, sensor.Users[i])
		log.Printf("   %v\n", userVideoAccess.Videos[i])
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
func TestGetVideosBySensorID(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/videoCrud.TestGetVideosBySensorID()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myVideo := &Video{}
	videos, err := myVideo.GetVideosBySensorID(myDB, 4)

	if err != nil {
		log.Printf("Test Failed %s", err)
		return
	}
	log.Printf("Select successful for ID: %v\n", videos)

	log.Printf("videos found inside sensorModel.GetResultsByUserID()")
	for i := 0; i < len(videos); i++ {
		log.Printf("   " + videos[i].VideoName)
	}

}

/*=====================================================================*\
	Functions completed the END 2 END naming conversion

\*=====================================================================*/
// TestGetVideosFullList ...
func TestGetVideosFullList(t *testing.T) {
	log.Printf("--------------------------------------------------------")
	log.Printf("RUNNING: ===>tablecrud/videoCrud.TestGetVideosFullList()")
	log.Printf("--------------------------------------------------------")
	myDB := db.Connect()
	myVideo := &Video{}
	myVideo.GetVideosFullList(myDB)
}
