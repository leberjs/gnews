package cmd

import (
	"encoding/json"
	"log"

	"github.com/leberjs/gnews/internal/models"

	bolt "go.etcd.io/bbolt"
)

func Add(db *bolt.DB, appId, appName string) {
  db.Update( func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("apps"))
    if b == nil {
      log.Fatal("Bucket `apps` does not exist")
    }

    if appName == "" {
      appName = appId
    }
    
    app := models.App{
      Id: appId,
      Name: appName,
    }

    if buf, err := json.Marshal(app); err != nil {
      return err
    } else if err := b.Put([]byte(appId), buf); err != nil {
      return err
    }

    return nil
  })
}
