package handlers

import (
  "net/http"
)

func ContentTypeHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    h.ServeHTTP(w, r)
  })
}
