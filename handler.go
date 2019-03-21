package main

import (
	"log"
	"net/http"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/kuser/message"
)

var (
	// TODO: Parse context from the request
	ctx = context.Background()
)

func getUserHandler(c *gin.Context) {

	uid := c.Param("id")
	if uid == "" {
		log.Println("Nil user ID")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Null user ID",
			"status":  http.StatusBadRequest,
		})
		return
	}

	usr, err := getUser(ctx, uid)
	if err != nil {
		log.Printf("Error while binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	log.Printf("Request: %v", usr)
	c.JSON(http.StatusOK, usr)

}


func saveUserHandler(c *gin.Context) {

	// bind
	var usr message.KUser
	err := c.BindJSON(&usr)
	if err != nil {
		log.Printf("Error while binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// save
	log.Printf("Request: %v", usr)
	err = saveUser(ctx, &usr)
	if err != nil {
		log.Printf("Error while saving user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Error",
			"status":  http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, usr)

}

func saveUserEventHandler(c *gin.Context) {

	// bind
	var event message.KUserEvent
	err := c.BindJSON(&event)
	if err != nil {
		log.Printf("Error while binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// save
	log.Printf("Request: %v", event)
	err = saveEvent(ctx, &event)
	if err != nil {
		log.Printf("Error while saving event: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Error",
			"status":  http.StatusBadRequest,
		})
		return
	}


	c.JSON(http.StatusOK, event)

}


func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}