package utils

// https://golang.org/src/net/http/status.go

import (
	"net/http"
	"fmt"
)

const (
	HTTP_OK                    = 0
	HTTP_CREATED               = 1
	HTTP_NO_CONTENT            = 2
	HTTP_BAD_REQUEST           = 3
	HTTP_UNAUTHORIZED          = 4
	HTTP_UNPROCESSABLE_ENTITY  = 5
	HTTP_INTERNAL_SERVER_ERROR = 6
)

type Status struct {
	Status int `json:"status"`
	Description string `json:"description"`
}

var HttpStatus = []Status{
	Status{http.StatusOK, "OK"}, 					 					// 200
	Status{http.StatusCreated, "CREATED"}, 			 					// 201
	Status{http.StatusNoContent, "NO CONTENT"}, 	 					// 204
	Status{http.StatusBadRequest, "BAD REQUEST"}, 	 					// 400
	Status{http.StatusUnauthorized, "UNAUTHORIZED"}, 					// 401
	Status{http.StatusUnprocessableEntity, "UNPROCESSABLE ENTITY"},		// 422
	Status{http.StatusInternalServerError, "INTERNAL SERVER ERROR"},	// 500
}

func HttpInfo(r *http.Request) {
	fmt.Printf("%s/\t %s\t %s%s\t %s\n", r.Method, r.Proto, r.Host, r.URL, DateTimeFormat())
}