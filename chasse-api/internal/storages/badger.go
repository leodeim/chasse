package storages

import (
	e "chasse-api/internal/error"
	"chasse-api/internal/logger"

	badger "github.com/dgraph-io/badger/v3"
)

var log = logger.New("BADGER")

type BadgerDB struct {
	client *badger.DB
}

func NewBadgerDB(mem bool, path string) (*BadgerDB, error) {
	b := BadgerDB{}
	var err error

	if mem {
		path = ""
	}

	b.client, err = badger.Open(
		badger.DefaultOptions(path).
			WithInMemory(mem).
			WithLogger(log),
	)
	if err != nil {
		return nil, e.BuildErrorf(e.INTERNAL, "failed at badger connection: %s", err.Error())
	}

	return &b, nil
}

func (b *BadgerDB) Get(key string) ([]byte, error) {
	var v []byte

	e := b.client.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		v, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	return v, e
}

func (b *BadgerDB) Set(key string, value []byte) error {
	return b.client.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})
}

func (b *BadgerDB) Close() error {
	return b.client.Close()
}

func (b *BadgerDB) Status() bool {
	return !b.client.IsClosed()
}
