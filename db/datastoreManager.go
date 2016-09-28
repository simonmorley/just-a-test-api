package db

import (
	// "database/sql"
	// "flag"
	"golang.org/x/net/context"
	// "google.golang.org/cloud/datastore"
	"cloud.google.com/go/datastore"
	"os"
)

type DatastoreManager struct {
	db *datastore.Client
}

func NewDatastore() (*DatastoreManager, error) {

	ctx := context.Background()
	d, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		return nil, err
	}
	return &DatastoreManager{d}, nil
}

func (m *DatastoreManager) RegisterUser(user string) error {
	// err := m.db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("users"))
	// 	err := b.Put([]byte(user), hash)
	// 	return err
	// })
	return nil
}

// func (m *DatastoreManager) Close() {
// 	m.db.Close()
// }
