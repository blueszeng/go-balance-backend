package model

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
)

type Model struct {
  errors errors.Errors  `sql:"-" json:"-"` // Ignore this field
}

func (m *Model) Errors() *errors.Errors { return &m.errors }

func (*Model) Find(id string) error { panic("Isn't implemented!"); return nil }

func (*Model) Build(url.Values) { panic("Isn't implemented!") }

func (*Model) Create() { panic("Isn't implemented!") }

func (*Model) Update() { panic("Isn't implemented!") }

func (*Model) Delete() { panic("Isn't implemented!") }
