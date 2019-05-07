package fileStore

/*===================================================================================*\
	This file was taken from a sample with much more capabilities.
	If the need arises the additional functionality can be uncommented.
\*===================================================================================*/
// https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
// To run:
// go run main.go

import (
	"fmt"

	"github.com/kjk/betterguid"
	"github.com/satori/go.uuid"
)

// func genXid() {
// 	id := xid.New()
// 	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
// }

// func genKsuid() {
// 	id := ksuid.New()
// 	fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
// }

func GenBetterGUID() string {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:   %s\n", id)
	return id
}

// func genUlid() {
// 	t := time.Now().UTC()
// 	entropy := rand.New(rand.NewSource(t.UnixNano()))
// 	id := ulid.MustNew(ulid.Timestamp(t), entropy)
// 	fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
// }

// func genSonyflake() {
// 	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
// 	id, err := flake.NextID()
// 	if err != nil {
// 		log.Fatalf("flake.NextID() failed with %s\n", err)
// 	}
// 	// Note: this is base16, could shorten by encoding as base62 string
// 	fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
// }

// func genSid() {
// 	id := sid.Id()
// 	fmt.Printf("github.com/chilts/sid:       %s\n", id)
// }

func GenUUIDv4() {
	id, err := uuid.NewV4()
	fmt.Printf("github.com/satori/go.uuid:   %s,%v\n", id, err)
}

func main() {
	// genXid()
	// genKsuid()
	GenBetterGUID()
	// genUlid()
	// genSonyflake()
	// genSid()
	GenUUIDv4()
}
