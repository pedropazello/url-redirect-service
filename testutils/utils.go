package testutils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func MakeGetRequest(path string, handler gin.HandlerFunc, requestPath string) *httptest.ResponseRecorder {
	r := gin.Default()
	r.GET(path, handler)
	req, _ := http.NewRequest("GET", requestPath, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
