package senCapComp

import (
	ds "AdvDashBoard/goAdvDashBoard/datastore"
	chk "AdvDashBoard/goAdvDashBoard/tablecrud/chunkCrud"
	"log"
	"testing"

	"github.com/fsnotify/fsnotify"
)

// func TestDeleteProductWithName(t *testing.T) {
// 	TestCaptureWatcher()
// }

// TestCaptureWatcher
func TestCaptureWatcher(t *testing.T) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					relpath := ds.GetNgCapturePath()
					chk.CreateChunkTest(dbRef, 2, 2, "Chunk One", "S3VC", "SomeOwner", relpath, 1, 1234)
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

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
