package kvstore

import (
	"github.com/ckm8token/ckm8/common"
	"github.com/ckm8token/ckm8/rlp"
	"github.com/ckm8token/ckm8/store"
	"github.com/ckm8token/ckm8/store/database"
)

// NewKVStore create a new instance of KVStore.
func NewKVStore(db database.Database) store.Store {
	return &KVStore{db}
}

// KVStore a Database wrapped object.
type KVStore struct {
	db database.Database
}

// Put upserts key/value into DB
func (store *KVStore) Put(key common.Bytes, value interface{}) error {
	encodedValue, err := rlp.EncodeToBytes(value)
	if err != nil {
		return err
	}
	return store.db.Put(key, encodedValue)
}

// Delete deletes key entry from DB
func (store *KVStore) Delete(key common.Bytes) error {
	return store.db.Delete(key)
}

// Get looks up DB with key and returns result into value (passed by reference)
func (store *KVStore) Get(key common.Bytes, value interface{}) error {
	encodedValue, err := store.db.Get(key)
	if err != nil {
		return err
	}
	return rlp.DecodeBytes(encodedValue, value)
}
