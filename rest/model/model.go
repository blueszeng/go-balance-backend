package model

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
)

type Model struct {
  errors errors.Errors  `sql:"-" json:"-"` // Ignore this field
}

func (m *Model) Errors() *errors.Errors { return &m.errors }

func (*Model) Build(url.Values) { panic("Isn't implemented!") }
