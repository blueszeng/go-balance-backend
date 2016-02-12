package handlers

import (
  "github.com/tksasha/go-rest"

  . "../models"
)

var Consolidates *rest.Handler

func init() {
  Consolidates = &rest.Handler {
    Collection: func() rest.Collection { return new(ConsolidateCollection) },
  }
}
