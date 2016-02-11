package handlers

import (
  . "./rest"
  . "../models"
)

var Cashes *RESTHandler

func init() {
  Cashes = &RESTHandler {
    Collection: func()  Collection  { return new(CashCollection) },
    Resource:   func()  Resource    { return new(Cash) },
  }
}
