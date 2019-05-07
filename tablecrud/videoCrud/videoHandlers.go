/*=====================================================================*\
	This is the Request/Reply handler of Backend Video
	Table Component. It is responsble for receiving RestAPI
	requests from the Frontend, interacting with the
	fileStore, constructing	and returning a response.
\*=====================================================================*/
package videoComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"
)

// CreateVideo ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>videoHandlers.CreateVideo()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)

	var video Video
	json.Unmarshal([]byte(har), &video)
	fmt.Printf("%v", video)

	myDB := db.Connect()

	log.Printf("   videoHandlers.Create():=%v\n", video)

	sqlErr := video.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created video videoHandlers.CreateVideo(): video=%v\n", video)
		json.NewEncoder(w).Encode(video)
		return
	}
	log.Printf("Error creating video in videoHandlers.CreateVideo(), Reason:%v\n", sqlErr)
}

// UpdateVideo ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func UpdateVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("===>videoHandlers.UpdateVideo()")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)

	pVideo := params["Video"]
	log.Printf("   Video=%v\n", pVideo)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var video Video
	json.Unmarshal([]byte(har), &video)
	fmt.Printf("%v", video)

	sqlErr := video.Update(myDB)
	if sqlErr == nil {
		log.Printf("Successful update in videoHandlers.UpdateVideo() video=%v\n", video)
		json.NewEncoder(w).Encode(video)
		return
	}
	log.Printf("Error updateing video in videoHandlers.UpdateVideo(), Reason:%v\n", sqlErr)
}

//DeleteVideo ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>videoHandlers.DeleteVideo()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myVideo := &Video{}
	paramInt, paramErr := strconv.Atoi(params["id"])

	if paramErr == nil {
		myVideo.ID = paramInt
		sqlErr := myVideo.Delete(myDB)
		if sqlErr == nil {
			log.Printf("Successful delete in service.DeleteVideos() video=%v\n", myVideo)
			json.NewEncoder(w).Encode(myVideo)
			return
		}
		log.Printf("Error with database delete in service.DeleteVideos(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in videoHandlers.DeleteVideos(), Reason:  %v\n", paramErr)
}

// GetVideo ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetVideo(w http.ResponseWriter, r *http.Request) {
	log.Println("===>videoHandlers.GetVideo()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myVideo := &Video{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		myVideo.ID = paramInt
		video, sqlErr := myVideo.GetByID(myDB)
		if sqlErr == nil {
			log.Printf("Successful select in service.GetVideos() video=%v\n", video)
			json.NewEncoder(w).Encode(video)
			return
		}
		log.Printf("Error with database select in service.GetVideo(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in videoHandlers.GetVideo(), Reason:  %v\n", paramErr)
}

/*=====================================================================*\
	Handlers Functions completed the END 2 END naming conversion

\*=====================================================================*/

// GetVideosFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetVideosFullList(w http.ResponseWriter, r *http.Request) {
	log.Println("===>videoHandlers.GetVideosFullList()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	myDB := db.Connect()
	myVideo := &Video{}
	videos, sqlErr := myVideo.GetVideosFullList(myDB)
	log.Printf("%d videos found inside videoHandlers.GetVideosFullList()\n", len(videos))
	// for i := 0; i < len(videos); i++ {
	// 	log.Printf("   %v\n", videos[i])
	// }

	if sqlErr == nil {
		log.Printf("Successful select of all videos in videoHandlers.GetVideosFullList()")
		json.NewEncoder(w).Encode(videos)
		return
	}
	log.Printf("Error with database select in videoHandlers.GetVideosFullList(), Reason:%v\n", sqlErr)
}

// GetVideosByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetVideosByOwnerID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.GetVideosByOwnerID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myVideo := &Video{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   UserID=%d\n", paramInt)

	if paramErr == nil {
		myVideo.ID = paramInt
		videos, sqlErr := myVideo.GetVideosByOwnerID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d videos found inside productHandlers.GetVideosByOwnerID()\n", len(videos))
			// for i := 0; i < len(videos); i++ {
			// 	log.Printf("  VideoName=%s,UserID=%d", videos[i].VideoName, videos[i].UserID)
			// }
			log.Printf("Successful select of all products in productHandlers.GetVideosByOwnerID()")
			json.NewEncoder(w).Encode(videos)
			return
		}
		log.Printf("Error with GetResultsByUserID() in productHandlers.GetVideosByOwnerID()\n")
		log.Printf("Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in productHandlers.GetVideosByOwnerID()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

// GetVideosByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetVideosByUserID(w http.ResponseWriter, r *http.Request) {
	log.Println(">===>productHandlers.GetVideosByUserID()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myVideo := &Video{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("   UserID=%d\n", paramInt)

	if paramErr == nil {
		myVideo.ID = paramInt
		videos, sqlErr := myVideo.GetVideosByUserID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d video found inside productHandlers.GetVideosByUserID()\n", len(videos))
			// for i := 0; i < len(videos); i++ {
			// 	log.Printf("  VideoName=%s,UserID=%d", videos[i].VideoName, videos[i].UserID)
			// }
			log.Printf("Successful select of all products in productHandlers.GetVideosByUserID()")
			json.NewEncoder(w).Encode(videos)
			return
		}
		log.Printf("Error with GetResultsByUserID() in productHandlers.GetVideosByUserID()\n")
		log.Printf("Reason:  %v\n", sqlErr)
		return
	}
	log.Printf("Error with Paramerer in productHandlers.GetVideosByUserID()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

// GetUsersByVideoID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetUsersByVideoID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>videoHandlers.GetUsersByVideoID()")

	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myProduct := &Video{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	log.Printf("  paramInt == %d", paramInt)
	if paramErr == nil {
		myProduct.ID = paramInt
		userList, sqlErr := myProduct.GetUsersByVideoID(myDB, paramInt)
		if sqlErr == nil {
			log.Printf("%d Users found in videoHandlers.GetUsersByVideoID()\n", len(userList))
			json.NewEncoder(w).Encode(userList)
			return
		}
		log.Printf("Error with database select in videoHandlers.GetUsersByVideoID(), Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in videoHandlers.GetUsersByVideoID(), Reason:  %v\n", paramErr)

}

// GetUsersByVideoID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetNotUsersByVideoID(w http.ResponseWriter, r *http.Request) {
	log.Println("===>sensorHandlers.GetUsersByVideoID()")

}
