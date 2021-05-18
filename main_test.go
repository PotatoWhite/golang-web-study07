package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	result, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(result), "Hello Go in Web")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	// badrequest
	resp, err := http.Post(ts.URL+"/users", "application/json", nil)
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, resp.StatusCode)

	// make a user
	reqUser, _ := json.Marshal(User{Name: "potato", Email: "potato@example.com"})

	resp, err = http.Post(ts.URL+"/users", "application/json", bytes.NewReader(reqUser))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	createdUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(createdUser)
	assert.NoError(err)
	assert.NotNil(createdUser.CreatedAt)
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	aUser := new(User)
	json.NewDecoder(resp.Body).Decode(aUser)

	assert.Equal("potato", aUser.Name)
	assert.Equal("potato@example.com", aUser.Email)
}
