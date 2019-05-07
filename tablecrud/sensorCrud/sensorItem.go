// Package sensorComp ...
/*=====================================================================*\

\*=====================================================================*/
package sensorComp

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// GetPidByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetPidByID(db *pg.DB, sensorID int) (int, error) {
	// log.Printf(">===>sensorItem.GetPidByID()")

	getErr := db.Model(gi).
		Where("id = ?0", sensorID).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in sensorItem.GetByOwner()\n")
		log.Printf("Reason %v\n", getErr)
		return gi.SensorPid, getErr
	}
	// log.Printf("Select by Owner successful for gi: %v\n", *gi)
	return gi.SensorPid, nil
}

// UpdatePid ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) UpdatePid(db *pg.DB, pid int, state int) error {
	// log.Printf(">===>sensorItem.Update()")

	_, updateErr := db.Model(gi).
		Set("sensorpid=?0,sensorstate=?1", pid, state).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in sensorModel.Update()\n")
		log.Printf("Reason %v\n", updateErr)
		return updateErr
	}
	// log.Printf("Product %s updated successfully in table", gi.Sensorname)
	return nil
}

// Create ... insert a record into the database
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) Create(db *pg.DB) error {
	log.Printf("===>sensorItem.Create()")
	log.Printf("   sensorItem.Create():=%v\n", gi)

	insertErr := db.Insert(gi)
	if insertErr != nil {
		log.Printf("Error writing to DB in sensorItem.Create()\n")
		log.Printf("Reason:%v\n", insertErr)
		return insertErr
	}
	log.Printf("Sensor %s inserted successfully into table", gi.Sensorname)
	return nil
}

// Update ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) Update(db *pg.DB) error {
	log.Printf("===>sensorItem.Update()")

	_, updateErr := db.Model(gi).
		Where("id = ?0", gi.ID).Update()
	if updateErr != nil {
		log.Printf("Error while updating item  in sensorItem.Update()\n")
		log.Printf("Reason %v\n", updateErr)
		return updateErr
	}
	log.Printf("Product %s updated successfully in table", gi.Sensorname)
	return nil
}

// Delete ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) Delete(db *pg.DB) error {
	log.Printf("===>sensorItem.Delete()")

	_, deleteErr := db.Model(gi).
		Where("sensorname = ?0", gi.Sensorname).
		WhereOr("id = ?0", gi.ID).
		Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item in sensorItem.Delete()\n")
		log.Printf("Reason %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Sensor %s deleted successfully from table", gi.Sensorname)
	return nil
}

// GetByName ... select record with matching name
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetByName(db *pg.DB) error {
	log.Printf("===>sensorItem.GetByName()")
	//getErr := db.Select(gi)
	getErr := db.Model(gi).
		Where("sensorname = ?0", gi.Sensorname).
		Select()
	if getErr != nil {
		log.Printf("Error while selecting item in sensorItem.GetByName()\n")
		log.Printf("Reason %v\n", getErr)
		return getErr
	}
	log.Printf("Select successful for ID: %v\n", *gi)
	return nil
}

// GetByID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetByID(db *pg.DB) (Sensor, error) {
	log.Printf("===>sensorItem.GetByID(SensorID=%d)", gi.ID)

	//getErr := db.Select(gi)
	getErr := db.Model(gi).Where("id = ?0", gi.ID).Select()
	if getErr != nil {
		log.Printf("Error while selecting item, Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select successful in sensorItem.GetById() sensor=%v\n", *gi)
	return *gi, nil
}

// GetByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetByUserID(db *pg.DB) (Sensor, error) {
	log.Printf("===>sensorItem.GetByUserID(ID=%d)", gi.ID)

	//getErr := db.Select(gi)
	getErr := db.Model(gi).Where("id = ?0", gi.ID).Select()
	if getErr != nil {
		log.Printf("Error in GetByUserID while selecting item\n")
		log.Printf("Reason %v\n", getErr)
		return *gi, getErr
	}
	log.Printf("Select successful in sensorItem.GetByUserID() sensor=%v\n", *gi)
	return *gi, nil
}

/*=====================================================================*\

	Item Functions that are using RELATIONSHIP queries

\*=====================================================================*/

// GetSensorsByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetSensorsByUserID(db *pg.DB, userid int) ([]Sensor, error) {
	log.Printf("===>sensorItem.GetSensorsByUserID()")

	var user User
	getErr := db.Model(&user).
		Relation("Sensors", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC", userid)
			return q, nil
		}).
		Where("id = ?0", userid).
		First()

	if getErr != nil {
		log.Printf("Error in sensorItem.GetSensorsByUserID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d sensors found inside sensorItem.GetSensorsByUserID()\n", len(user.Sensors))
	// for i := 0; i < len(user.Sensors); i++ {
	// 	log.Printf("  %v\n", user.Sensors[i])
	// }
	return user.Sensors, getErr
}

// GetNotUsersBySensorID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetNotUsersBySensorID(db *pg.DB, sensorid int) ([]User, error) {
	log.Printf("===>sensorItem.GetNotUsersBySensorID()")
	var sensor Sensor
	var user User

	getErr := db.Model(&sensor, &user).
		Relation("Users", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC", sensorid)
			return q, nil
		}).
		Where("id != ?0", sensorid).
		First()

	if getErr != nil {
		log.Printf("Error in sensorItem.GetUsersBySensorID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	return sensor.Users, getErr
}

// GetUsersBySensorID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetUsersBySensorID(db *pg.DB, sensorid int) ([]User, error) {
	log.Printf("===>sensorItem.GetUsersBySensorID()")

	var sensor Sensor
	getErr := db.Model(&sensor).
		Relation("Users", func(q *orm.Query) (*orm.Query, error) {
			q = q.OrderExpr("id ASC", sensorid)
			return q, nil
		}).
		Where("id = ?0", sensorid).
		First()

	if getErr != nil {
		log.Printf("Error in sensorItem.GetUsersBySensorID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}

	log.Printf("%d user found inside sensorItem.GetUsersBySensorID()\n", len(sensor.Users))
	// for i := 0; i < len(sensor.Users); i++ {
	// 	log.Printf("   %v", sensor.Users[i])
	// }
	return sensor.Users, getErr
}

/*=====================================================================*\
	Functions completed the END 2 END naming conversion

\*=====================================================================*/
// GetSensorsFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetSensorsFullList(db *pg.DB) ([]Sensor, error) {
	log.Printf("===>sensorItem.GetSensorsFullList()")
	var sensors []Sensor
	getErr := db.Model(&sensors).Column("*").
		Offset(0).
		Order("id asc").
		Select()
	if getErr != nil {
		log.Printf("Error while selecting all sensors in sensorItem.GetSensorsFullList()\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("%d sensors found inside sensorItem.GetSensorsFullList()\n", len(sensors))
	// for i := 0; i < len(sensors); i++ {
	// 	log.Printf("   " + sensors[i].Sensorname)
	// }
	return sensors, getErr
}

// GetSensorsByOwnerID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetSensorsByOwnerID(db *pg.DB, userid int) ([]Sensor, error) {
	log.Printf("===>sensorItem.GetSensorsByOwnerID()")
	var sensors []Sensor

	getErr := db.Model(&sensors).Column("*").
		Offset(0).
		Order("id asc").
		Where("userid = ?0", userid).
		Select()

	if getErr != nil {
		log.Printf("Error while selecting all sensors in sensorItem.GetSensorsByOwnerID()\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}
	log.Printf("Select successful for ID: %v\n", userid)

	return sensors, getErr
}

/*=====================================================================*\

	UNUSED SENSOR ITEM FUNCTIONS, SHOULD BE REMOVED SOON

\*=====================================================================*/

// GetResultsByUserID ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func (gi *Sensor) GetResultsSensorsAccessableByUserID_NOTUSED(db *pg.DB, sensorid int) ([]Sensor, error) {
	log.Printf("===>sensorItem.GetResultsByUserID()")
	// Register many to many model so ORM can better recognize m2m relation.
	// This should be done before dependant models are used.
	//orm.RegisterTable((*SensorToUser)(nil))
	var sensors []Sensor

	log.Printf("   %v", sensors)
	// A try using Relation key word.
	getErr := db.Model(&sensors).
		Column("*", "sen_user_items").
		Join("inner join sen_user_item su on ?0 = su.sensorid", sensorid).
		Where("sensors.user_id_s = sen_user_item.user_id").Select()
	// Relation("sen_user_item", func(q *orm.Query) (*orm.Query, error) {
	// 	return q.Where("sen_user_item.sensor_id = ?0", sensorid), nil
	// }).
	//Where("sensorid = ?0", sensorid).
	//Select()

	if getErr != nil {
		log.Printf("Error in sensorModel.GetResultsSensorsAccessableByUserID\n")
		log.Printf("Reason %v\n", getErr)
		return nil, getErr
	}
	// } else {
	// getErr := db.Model(sensor).
	// 	Column("*"). // Column("sensor.*", "Users").
	// 	Where("sensor.user_id_s = ?", userid).
	// 	Select()

	//	var customers []*Customer

	// getErr := db.Model(&sensors).Column("*").
	// 	Join("inner join relsensoruser su on sensors.sensorid = su.sensor_id").
	// 	Where("user_id_s = ?", userid).Select()
	// if getErr != nil {
	// 	// Error Handler
	// 	log.Printf("Error while selecting all sensors in sensorModel.GetAllResults(), Reason %v\n", getErr)
	// 	return nil, getErr
	// } else {
	// 	for _, sensor := range sensors {
	// 		fmt.Printf("Customer -> id: %d, username:%s \n", sensor.SensorID, sensor.SensorName)
	// 	}
	// }

	log.Printf("%d sensors found inside sensorModel.GetAllResults()\n", len(sensors))
	// for i := 0; i < len(sensors); i++ {
	// 	log.Printf("   " + sensors[i].SensorName)
	// }
	return sensors, getErr
}
