package main

import (
	//"go-postgres/function"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var result string = "OK response is expected"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", ReadingItem).Methods("GET")
	router.HandleFunc("/movie/3", ReadingItemid).Methods("GET")
	router.HandleFunc("/movie", CreateItem).Methods("POST")
	router.HandleFunc("/movie/2", UpdateItems).Methods("PUT")
	router.HandleFunc("/movie/1", Softdelete).Methods("DELETE")
	router.HandleFunc("/movie/4", DeleteItem).Methods("DELETE")
	return router
}

func TestReadingItem(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, result)
}

func TestCreateItem(t *testing.T) {
	request, _ := http.NewRequest("POST", "/movie", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, result)
}

func TestReadItem(t *testing.T) {
	request, _ := http.NewRequest("GET", "/movie/3", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, result)
}

func TestUpdateItem(t *testing.T) {
	request, _ := http.NewRequest("PUT", "/movie/2", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, result)
}
func TestSoftDelete(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/movie/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, result)
}

func TestDeleteItem(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/movie/4", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, result)
}
