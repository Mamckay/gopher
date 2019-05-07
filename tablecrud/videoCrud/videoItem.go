// Package videoComp ...
/*=====================================================================*\

\*=====================================================================*/
package videoComp

import (
	sen "AdvDashBoard/goAdvDashBoard/tablecrud/sensorCrud"
	usr "AdvDashBoard/goAdvDashBoard/tablecrud/userCrud"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) Create(db *pg.DB) error {
	log.Printf("===>videoItem.Create()")
	log.Printf("   videoItem.Create():=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in videoItem.Create(), Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("Video %s inserted successfully into table", gi.VideoName)
	return nil
}

// Update ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) Update(db *pg.DB) error {
	log.Printf("===>.Update()")

	_, updateErr := db.Model(gi).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in videoItem.Update(), Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Product %s updated successfully in table", gi.VideoName)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) Delete(db *pg.DB) error {
	log.Printf("===>videoItem.Delete()")

	_, deleteErr := db.Model(gi).
		Where("videoname = ?0", gi.VideoName).
		WhereOr("id = ?0", gi.ID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item in videoItem.Delete(), Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Product %s deleted successfully from table", gi.VideoName)
	return nil
}

// GetByName ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) GetByName(db *pg.DB) error {
	log.Printf("===>videoItem.GetByName()")
	//getErr := db.Select(gi)
	getErr := db.Model(gi).
		Where("videoname = ?0", gi.VideoName).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in videoItem.GetByName(), Reason %v\n", getErr)
		return getErr
	}
	log.Printf("Select successful for ID: %v\n", *gi)
	return nil
}

// GetByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) GetByID(db *pg.DB) (Video, error) {
	log.Printf("===>videoItem.GetByID(ID=%d)", gi.ID)

	//getErr := db.Select(gi)
	getErr := db.Model(gi).Where("id = ?0", gi.ID).Select()
	if getErr != nil {
		log.Printf("Error while selecting item, Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select successful in videoItem.GetById() video=%v\n", *gi)
	return *gi, nil
}

// GetVideosBySensorID Get the Videos for the sensor owner.
/*---------------------------------------------------------------------*\
BAD BAD
\*---------------------------------------------------------------------*/
func (gi *Video) GetVideosBySensorID(db *pg.DB, sensorid int) ([]Video, error) {
	log.Printf(">===>videoItem.GetVideosBySensorID(sensorid=%d)", sensorid)

	var videos []Video
	var sensors []sen.Sensor
	var user usr.User
	var users []usr.User
	log.Printf("users= %v\n", users)
	log.Printf("sensors= %v\n", sensors)
	log.Printf("videos= %v\n", videos)
	log.Printf("Video Queary User  %v", user)

	getErr := db.Model(&videos).Column("*").
		// Offset(0).
		// Order("id asc").
		// Relation("infuser").
		//Join("inner join infuser  AS u ON infvideo.userid = 2 And infuser.userid = 2").
		//ON("sensor.sensorid = infvideo.sensorid").
		Where("sensor_id = ?0", sensorid).
		Select()

	if getErr != nil {
		log.Printf("Error while selecting all videos in videoItem.GetVideosBySensorID(), Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("%d videos found inside videoItem.GetVideosBySensorID()\n", len(videos))
	// for i := 0; i < len(videos); i++ {
	// 	log.Printf("  %v\n", videos[i])
	// }
	return videos, getErr
}

/*=====================================================================*\

	Item Functions that are using RELATIONSHIP queries

\*=====================================================================*/

// GetUsersByVideoID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) GetUsersByVideoID(db *pg.DB, videoID int) ([]User, error) {
	log.Printf("===>videoItem.GetUsersByVideoID()")

	var video Video
	getErr := db.Model(&video).
		Relation("Users", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC", videoID)
			return q, nil
		}).
		Where("id = ?0", videoID).
		First()

	if getErr != nil {
		log.Printf("Error in videoItem.GetUsersByVideoID, Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d users found inside videoItem.GetUsersByVideoID()", len(video.Users))
	// for i := 0; i < len(video.Users); i++ {
	// 	log.Printf("   %v\n", video.Users[i])
	// }
	return video.Users, getErr
}

// GetVideosByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) GetVideosByUserID(db *pg.DB, userid int) ([]Video, error) {
	log.Printf("===>videoItem.GetVideosByUserID()")

	var user User
	getErr := db.Model(&user).
		Relation("Videos", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC", userid)
			return q, nil
		}).
		Where("id = ?0", userid).
		First()

	if getErr != nil {
		// Error Handler
		log.Printf("Error in videoItem.GetVideosByUserID, Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d Videos found inside videoItem.GetVideosByUserID()\n", len(user.Videos))
	// for i := 0; i < len(user.Videos); i++ {
	// 	log.Printf("  %v\n", user.Videos[i])
	// }
	return user.Videos, getErr
}

/*=====================================================================*\
	Functions completed the END 2 END naming conversion
\*=====================================================================*/

// GetVideosByOwnerID witll Get the videos for the user owner.
// This is only for example as per desing users do not own videos.
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) GetVideosByOwnerID(db *pg.DB, userid int) ([]Video, error) {
	log.Printf(">===>videoItem.GetVideosByOwnerID(userid=%d)", userid)

	var videos []Video
	var sensors []sen.Sensor
	var user usr.User
	var users []usr.User
	log.Printf("users= %v\n", users)
	log.Printf("sensors= %v\n", sensors)
	log.Printf("videos= %v\n", videos)
	log.Printf("Video Queary User  %v", user)

	getErr := db.Model(&videos).Column("*").
		// Offset(0).
		// Order("id asc").
		// Relation("infuser").
		//Join("inner join infuser  AS u ON infvideo.userid = 2 And infuser.userid = 2").
		//ON("sensor.sensorid = infvideo.sensorid").
		Where("userid = ?0", userid).
		Select()

	if getErr != nil {
		log.Printf("Error while selecting all videos in videoItem.GetVideosByOwnerID(), Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("%d videos found inside videoItem.GetVideosByOwnerID()\n", len(videos))
	// for i := 0; i < len(videos); i++ {
	// 	log.Printf("  %v\n", videos[i])
	// }
	return videos, getErr
}

// GetVideosFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Video) GetVideosFullList(db *pg.DB) ([]Video, error) {
	log.Printf("===>videoItem.GetVideosFullList()")
	var videos []Video
	getErr := db.Model(&videos).Column("*").
		Offset(0).
		Order("id asc").
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all videos in videoItem.GetVideosFullList(), Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("%d videos found inside videoItem.GetVideosFullList()\n", len(videos))
	// for i := 0; i < len(videos); i++ {
	// 	log.Printf("  %v\n", videos[i])
	// }
	return videos, getErr
}

/*=====================================================================*\

	UNUSED VIDEO ITEM FUNCTIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/
