/*=====================================================================*\



\*=====================================================================*/
package dataBase

import (
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

/*---------------------------------------------------------------------*\

\*---------------------------------------------------------------------*/

// Connect ...
func Connect() *pg.DB {
	//log.Printf(">===>pgDBMain.Connect()")
	opts := &pg.Options{
		User:         "postgres",
		Password:     "pg",
		Addr:         "localhost:5432",
		Database:     "AdvDashBoard",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		MaxAge:       1 * time.Minute, //do a ping to db every minute
		PoolSize:     20,
	}
	var db = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connection to database [%s] successful.\n", opts.Database)
	return db
}
