package cmd

import (
	"encoding/json"
	"log"

	// "github.com/leberjs/gnews/internal/models"

	bolt "go.etcd.io/bbolt"
)

// func List(db *bolt.DB) []models.App {
func List(db *bolt.DB) []map[string]any {
  // var res []models.App
  var res []map[string]any

  db.View( func(tx *bolt.Tx) error{
    b := tx.Bucket([]byte("apps"))
    c := b.Stats().KeyN

    // res = make([]models.App, c)
    res = make([]map[string]any, c)
    i := 0

    b.ForEach(func(k, v []byte) error{
      var t map[string]any
      // var t models.App
      if err := json.Unmarshal(v, &t); err != nil {
        log.Fatal(err)
      }
      // res[i] = models.App{Id: string(k), Name: string(v)}
      res[i] = t
      i++

      return nil
    })

    return nil
  })

  return res
}
