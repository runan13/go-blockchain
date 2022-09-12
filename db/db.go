package db

import (
	"awesomeProject/utils"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

const (
	dbName       = "blockcahin.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

func DB() *bolt.DB {
	if db == nil {
		boltDB, err := bolt.Open(dbName, 0600, nil)
		utils.HandleError(err)
		db = boltDB

		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleError(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			utils.HandleError(err)
			return err
		})
		utils.HandleError(err)
	}
	return db
}
