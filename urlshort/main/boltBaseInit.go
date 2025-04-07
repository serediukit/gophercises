package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("mybolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

	bucketName := []byte("paths")

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("/war"), []byte("https://en.wikipedia.org/wiki/Russo-Ukrainian_War"))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("/ukraine"), []byte("https://en.wikipedia.org/wiki/Ukraine"))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}
