package database

import (
	"encoding/binary"
	"time"

	"go.etcd.io/bbolt"
)

var taskBucket = []byte("tasks")
var db *bbolt.DB

type Task struct {
	Key   int
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

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(int(id64))
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ReadAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
