package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	"fmt"
	"encoding/json"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNilUserHandler(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping TestNilUserHandler")
	}

	initStore(context.Background())

	router := gin.Default()
	router.GET("/:id", getUserHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/123", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestUserHandler(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping TestUserHandler")
	}

	ctx := context.Background()
	initStore(ctx)

	// save
	saveRouter := gin.Default()
	testURI := "/user"
	saveRouter.POST(testURI, saveUserHandler)

	usr := getTestUserFromID("handler-123")
	b, err := json.Marshal(usr)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", testURI, bytes.NewBuffer(b))
	saveRouter.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// get
	getRouter := gin.Default()
	getRouter.GET("/user/:id", getUserHandler)

	w2 := httptest.NewRecorder()
	uri := fmt.Sprintf("/user/%s", usr.ID)
	req2, _ := http.NewRequest("GET", uri, nil)

	getRouter.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)

	//events
	eventRouter := gin.Default()
	testEventURI := "/event"
	eventRouter.POST(testEventURI, saveUserEventHandler)

	e1 := getTestEvent(usr.ID, "test-evetn-1")
	b2, err := json.Marshal(e1)
	assert.Nil(t, err)

	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", testEventURI, bytes.NewBuffer(b2))
	eventRouter.ServeHTTP(w3, req3)
	assert.Equal(t, 200, w3.Code)



}
