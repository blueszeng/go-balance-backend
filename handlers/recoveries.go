package handlers

import (
  "github.com/tksasha/go-rest"

  . "../models"
)

var Recoveries *rest.Handler

func init() {
  Recoveries = &rest.Handler {
    Resource: func() rest.Resource { return new(Recovery) },
  }
}