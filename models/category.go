package models

import (
  "net/url"
  "strings"
  "time"

  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/validations"
  . "github.com/tksasha/balance/config"
)

func init() {
  DB.AutoMigrate(&Category{})
}

type Category struct {
  Model

  ID        int        `json:"id"`
  Name      string     `json:"name"`
  Income    bool       `json:"income"`
  CreatedAt time.Time  `json:"-"`
  UpdatedAt time.Time  `json:"-"`
  DeletedAt *time.Time `json:"-"`
}

type CategoryCollection []Category

func (c *Category) Build(values url.Values) {
  if name := values.Get("category[name]"); name != "" {
    c.Name = name
  }

  if income := values.Get("category[income]"); income != "" {
    switch values.Get("category[income]") {
    case "1", "true":
      c.Income = true
    default:
      c.Income = false
    }
  }
}

func (c *Category) IsValid() bool {
  ValidatePresenceOf(c, "Name")

  c.validateUniquenessOfName()

  return c.Errors().IsEmpty()
}

func (c *Category) IsCreate(values url.Values) bool {
  c.Build(values)

  if c.IsValid() {
    DB.Create(c)

    return true
  } else {
    return false
  }
}

func (c *Category) IsUpdate(values url.Values) bool {
  c.Build(values)

  if c.IsValid() {
    DB.Save(c)

    return true
  } else {
    return false
  }
}

func (c *Category) validateUniquenessOfName() {
  var count int

  query := DB.Table("categories").Where("LOWER(name)=?", strings.ToLower(c.Name))

  if c.ID != 0 {
    query = query.Not("id = ?", c.ID)
  }

  query.Count(&count)

  if count > 0 {
    c.Errors().Add("name", "already exists")
  }
}

//
// CategoryCollection.Search
//
func (c *CategoryCollection) Search(values url.Values) {
  DB.Order("income").Find(&c)
}
