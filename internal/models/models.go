package models

type App struct {
  Id string `json:"id"`
  Name string `json:"name"`
}

type ReadCmdResponse struct {
  Name string
  Contents string
}

type NewsApiResponse struct {
  AppNews AppNews `json:"appnews"`
}

type AppNews struct {
  NewsItems []NewsItems `json:"newsitems"`
}

type NewsItems struct {
  Contents string `json:"contents"`
}

