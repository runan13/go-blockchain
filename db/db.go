package db

import (
	"awesomeProject/utils"
	"github.com/boltdb/bolt"
)

var db *bolt.DB

const (
	dbName       = "blockchain.db"
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

func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleError(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleError(err)
}
