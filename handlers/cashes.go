package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"

  . "../config"
  . "../models"
)

type Cashes struct {
  BaseHandler
}

//
// POST /cashes
//
func (c *Cashes) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  cash := new(Cash)

  if cash.IsCreate(r.Form) {
    c.render(w, cash, 200)
  } else {
    c.render(w, cash.Errors, 422)
  }
}

//
// GET /cashes
//
func (c *Cashes) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var cashes []Cash

  DB.Find(&cashes)

  c.render(w, cashes, 200)
}
