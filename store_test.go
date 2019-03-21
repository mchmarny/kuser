package main

import (
	"context"
	"testing"
	"time"
	"fmt"

	"github.com/mchmarny/kuser/message"

	"github.com/stretchr/testify/assert"

)

func getTestUserFromID(id string) *message.KUser {
	return &message.KUser{
		ID:      id,
		Email:   fmt.Sprintf("id-%s@domain.com", id),
		Name:    "Test User",
		Created: time.Now(),
		Updated: time.Now(),
		Picture: "http://invalid.domain.com/pic1",
	}
}

func getTestEvent(userID, id string) *message.KUserEvent {
	return &message.KUserEvent{
		ID:      id,
		UserID: userID,
		On: time.Now(),
		Data: []*message.KDataItem{
			&message.KDataItem{ Key: "d1", Value: "v1" },
			&message.KDataItem{ Key: "d2", Value: "v2" },
			&message.KDataItem{ Key: "d3", Value: "v3" },
		},
	}
}

func TestUser(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping TestSaveUser")
	}

	ctx := context.Background()
	initStore(ctx)

	// create
	usr := getTestUserFromID("store-123")
	err := saveUser(ctx, usr)
	assert.Nil(t, err)

	// get
	usr2, err := getUser(ctx, usr.ID)
	assert.Nil(t, err)
	assert.NotNil(t, usr2)
	assert.Equalf(t, usr.ID, usr2.ID, "Users' ID don't equal %s != %s", usr.ID, usr2.ID)

	// create events for user
	event1 := getTestEvent(usr2.ID, "e1")
	err = saveEvent(ctx, event1)
	assert.Nil(t, err)

	event2 := getTestEvent(usr2.ID, "e2")
	err = saveEvent(ctx, event2)
	assert.Nil(t, err)

	// delete user and its events
	err = deleteUser(ctx, usr2.ID)
	assert.Nil(t, err)

}
