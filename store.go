package main

import (
	"errors"
	"log"
	"context"
	"fmt"
	"cloud.google.com/go/firestore"
	"github.com/mchmarny/kuser/message"
	"google.golang.org/api/iterator"
)

const (
	userCollectionName = "kuser"
	eventCollectionName = "kuser-event"
)

var (
	fsClient   *firestore.Client
	userColl   *firestore.CollectionRef
	eventColl   *firestore.CollectionRef
	errNilDocRef = errors.New("firestore: nil DocumentRef")
)


func initStore(ctx context.Context) {

	// in case called multiple times during test
	if eventColl != nil && userColl != nil && fsClient != nil {
		return
	}

	projectID := mustGetEnv("GCP_PROJECT_ID", "")
	log.Printf("Initiating firestore in %s project", projectID)

	// Assumes GOOGLE_APPLICATION_CREDENTIALS is set
	c, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Error while creating Firestore client: %v", err)
	}
	fsClient = c
	userColl = c.Collection(userCollectionName)
	eventColl = c.Collection(eventCollectionName)
}


func getUser(ctx context.Context, id string) (usr *message.KUser, err error) {

	if id == "" {
		return nil, errors.New("Nil job ID parameter")
	}

	d, err := userColl.Doc(id).Get(ctx)
	if err != nil {
		if err == errNilDocRef {
			return nil, fmt.Errorf("No user for ID: %s", id)
		}
		return nil, err
	}

	var u message.KUser
	if err := d.DataTo(&u); err != nil {
		return nil, fmt.Errorf("Stored data not user: %v", err)
	}

	return &u, nil

}


func saveUser(ctx context.Context, usr *message.KUser) error {

	if usr == nil || usr.ID == "" {
		log.Println("nil id on user save")
		return errors.New("Nil ID")
	}

	_, err := userColl.Doc(usr.ID).Set(ctx, usr)
	if err != nil {
		log.Printf("error on save: %v", err)
		return fmt.Errorf("Error on save: %v", err)
	}

	return nil

}

func saveEvent(ctx context.Context, event *message.KUserEvent) error {

	if event == nil || event.ID == "" {
		log.Println("nil id on event save")
		return errors.New("Nil ID")
	}

	_, err := eventColl.Doc(event.ID).Set(ctx, event)
	if err != nil {
		log.Printf("error on save: %v", err)
		return fmt.Errorf("Error on save: %v", err)
	}

	return nil

}


func deleteUser(ctx context.Context, id string) error {

	if id == "" {
		return errors.New("Nil job ID parameter")
	}

	batch := fsClient.Batch()

	doc, err := userColl.Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error on doc get: %v", err)
		return err
	}
	batch.Delete(doc.Ref)

	iter := eventColl.Where("userId", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error on iter next: %v", err)
			return err
		}
		batch.Delete(doc.Ref)
	}

	_, err = batch.Commit(ctx)
	return err

}

