package cmd

import (
	"encoding/json"
	"log"

	"github.com/leberjs/gnews/internal/models"
	bolt "go.etcd.io/bbolt"
)

func List(db *bolt.DB) []models.App {
  var ret []models.App

  db.View( func(tx *bolt.Tx) error{
    b := tx.Bucket([]byte("apps"))
    c := b.Stats().KeyN

    ret = make([]models.App, c)
    i := 0

    b.ForEach(func(k, v []byte) error{
      var t models.App
      if err := json.Unmarshal(v, &t); err != nil {
        log.Fatal(err)
      }
      ret[i] = t
      i++

      return nil
    })
    return nil
  })

  return ret
}
