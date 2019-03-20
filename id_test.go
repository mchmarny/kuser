package main

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {

	testEmail := "Test@Chmarny.com"

	id1 := makeID(testEmail)
	log.Printf("ID1: %s", id1)

	testEmail2 := strings.ToLower(testEmail)

	id2 := makeID(testEmail2)
	log.Printf("ID2: %s", id2)
	assert.Equalf(t, id1, id2, "IDs don't equal %s != %s", id1, id2)


	email, err := parseID(id2)
	log.Printf("Email: %s", email)
	assert.Nil(t, err)
	assert.Equalf(t, email, testEmail2, "Emails don't equal %s != %s", email, testEmail2)

}
