/*=====================================================================*\
	This is the UnitTest for the Backend Video Component


\*=====================================================================*/
package videoComp

import (
	"log"

	"github.com/go-pg/pg"
)

// TestGetVideosFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func TestGetVideosFullList(dbRef *pg.DB) {
	log.Printf("===>videoTester.TestGetVideosFullList()")

	myVideo := &Video{}
	myVideo.GetVideosFullList(dbRef)
}
func GetVideoByName(dbRef *pg.DB, videoName string) {
	log.Printf("===>videoTester.GetVideoByName()")

	myVideo := &Video{VideoName: videoName}
	myVideo.GetByName(dbRef)
}
func GetVideoById(dbRef *pg.DB, id int) {
	log.Printf("===>videoTester.GetVideoById()")

	myVideo := &Video{ID: id}
	myVideo.GetByID(dbRef)
}
func DeleteVideoWithId(dbRef *pg.DB, id int) {
	log.Printf("===>videoTester.DeleteVideoWithId()")

	myVideo := &Video{ID: id}
	myVideo.Delete(dbRef)
}
func DeleteVideoWithName(dbRef *pg.DB, videoName string) {
	log.Printf("===>videoTester.DeleteVideoWithName()")

	myVideo := &Video{VideoName: videoName}
	myVideo.Delete(dbRef)
}
