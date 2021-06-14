package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// testing process develop
func TestAllHandler(t *testing.T) {
	router := SetupRouter()

	testTable := []struct {
		Name         string
		Method       string
		Path         string
		body         gin.H
		ExpectStatus int
		ExpectBody   gin.H
	}{}

	for _, tc := range testTable {
		t.Run(tc.Name, func(t *testing.T) {
			w := performRequest(router, tc.Method, tc.Path)

			fmt.Println(w)
		})
	}
}
