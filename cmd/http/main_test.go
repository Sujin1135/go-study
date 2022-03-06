package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	method := "GET"

	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, "/", nil)

	mux := NewWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(fmt.Sprintf("Requested method is %s\n", method), string(data))
}

func TestIndexPathHandlerWithPostMethod(t *testing.T) {
	assert := assert.New(t)
	method := "POST"

	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, "/", nil)

	mux := NewWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(fmt.Sprintf("Requested method is %s\n", method), string(data))
}

func TestStudentPostHandler(t *testing.T) {
	assert := assert.New(t)
	method := "POST"

	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, "/students", nil)

	mux := NewWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	body := new(student)
	err := json.NewDecoder(res.Body).Decode(body)
	assert.Nil(err)
	expected := newStudent("철수", 59, 100)

	assert.Equal(body.Name, expected.Name)
	assert.Equal(body.Age, expected.Age)
	assert.Equal(body.Score, expected.Score)
}
