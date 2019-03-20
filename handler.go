package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/kuser/message"
)

func getUserHandler(c *gin.Context) {

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

	log.Printf("Request: %v", usr)


	c.JSON(http.StatusOK, usr)

}

func saveUserEventHandler(c *gin.Context) {

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

	log.Printf("Request: %v", usr)


	c.JSON(http.StatusOK, usr)

}


func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}