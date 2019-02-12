package utils

import (
  "net/http"
  "encoding/json"
)

type DefaultResponse struct {
  Data    interface{} `json:"data"`
  Status  int         `json:"status"`
}

type DefaultError struct{
  Message string `json:"message"`
  Status  int    `json:"status"`
}

func ToJson(w http.ResponseWriter, data interface{}) error {
  w.Header().Set("Content-Type", "application/json")
  return json.NewEncoder(w).Encode(data)
}