package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStudentListHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

func TestGetStudentHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var student Student
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal(student.Name, "aaa")
}
