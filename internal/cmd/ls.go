package cmd

import (
	"encoding/json"
	"log"

	bolt "go.etcd.io/bbolt"
)

func List(db *bolt.DB) []map[string]any {
  var res []map[string]any

  db.View( func(tx *bolt.Tx) error{
    b := tx.Bucket([]byte("apps"))
    c := b.Stats().KeyN

    res = make([]map[string]any, c)
    i := 0

    b.ForEach(func(k, v []byte) error{
      var t map[string]any
      if err := json.Unmarshal(v, &t); err != nil {
        log.Fatal(err)
      }
      res[i] = t
      i++

      return nil
    })
    return nil
  })

  return res
}
