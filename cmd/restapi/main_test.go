package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateStudentHandler(t *testing.T) {
	assert := assert.New(t)
	name := "test"

	studentStr := `{"Id": 3, "Name": "test", "Age": 30, "Score": 85}`
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students", strings.NewReader(studentStr))
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)
	var body Student
	err := json.NewDecoder(res.Body).Decode(&body)

	assert.Nil(err)
	assert.Equal(http.StatusCreated, res.Code)
	assert.Equal(body.Name, name)
	assert.Equal(body.Age, 30)
	assert.Equal(body.Score, 85)
}
