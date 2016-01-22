package models

import (
  "time"
  "net/url"
  "strconv"

  "github.com/tksasha/go-date"
)

type Item struct {
  ID          int       `json:"id"`
  Date        time.Time `json:"date"`
  Sum         float32   `json:"sum"`
  Formula     string    `json:"formula"`
  Description string    `json:"description"`
  Category    Category  `json:"category"`
  CategoryID  int       `json:"category_id"`
}

func (i *Item) Build(values url.Values) {
  i.Date = date.New(values.Get("item[date]")).Time()

  if id, err := strconv.Atoi(values.Get("item[category_id]")); err == nil {
    i.CategoryID = id
  }

  if sum, err := strconv.ParseFloat(values.Get("item[sum]"), 32); err == nil {
    i.Sum = float32(sum)
  }
}
