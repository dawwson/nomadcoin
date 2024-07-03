package db

import (
	"github.com/boltdb/bolt"
	"github.com/dawwson/nomadcoin/utils"
)

const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

// NOTE: singleton pattern
func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		utils.HandleErr(err)
		db = dbPointer

		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}