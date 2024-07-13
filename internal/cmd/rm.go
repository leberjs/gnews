package cmd

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

func Remove(db *bolt.DB, appId string) {
  db.Update(func(tx *bolt.Tx) error{
    b := tx.Bucket([]byte("apps"))
    if b == nil {
      log.Fatal("Bucket `apps` does not exist")
    }

    err := b.Delete([]byte(appId))

    return err
  })
}
