package database

import (
	"time"

	"go.etcd.io/bbolt"
)

var taskBucket = []byte("tasks")
var db *bbolt.DB

type Task struct {
	Key   string
	Value string
}

func OpenDB(dbPath string) error {
	var err error
	db, err = bbolt.Open(dbPath, 0600, &bbolt.Options{
		Timeout: 1 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
