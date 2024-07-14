package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/leberjs/gnews/internal/models"
	bolt "go.etcd.io/bbolt"
)

func Read(db *bolt.DB) []models.ReadCmdResponse {
  var ret []models.ReadCmdResponse
  
  u, _ := url.Parse("http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002")

  db.View( func(tx *bolt.Tx) error{
    b := tx.Bucket([]byte("apps"))
    c := b.Stats().KeyN

    ret = make([]models.ReadCmdResponse, c)
    i := 0

    b.ForEach( func(k, v []byte) error{
      var (
        news models.NewsApiResponse
        app models.App
      )

      if err := json.Unmarshal(v, &app); err != nil {
        log.Fatal(err)
      }

      query := url.Values{}
      query.Set("appid", string(k))
      query.Set("count", "1")
      query.Set("format", "json")

      u.RawQuery = query.Encode()

      res, err := http.Get(u.String())
      if err != nil {
        log.Fatal(err)
      }

      defer res.Body.Close()

      body, _ := io.ReadAll(res.Body)

      if err := json.Unmarshal(body, &news); err != nil {
        log.Fatal(err)
      }

      ret[i] = models.ReadCmdResponse{Name: app.Name, Contents: news.AppNews.NewsItems[0].Contents}
      i++

      return nil
    })

    return nil
  })

  return ret
}
