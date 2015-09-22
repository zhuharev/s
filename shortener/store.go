package shortener

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type Store interface {
	Set(key, val string) error
	Get(key string) (string, error)

	Has(key string) (bool, error)
}

var (
	ErrNotFound = fmt.Errorf("not found")
	ErrExists   = fmt.Errorf("already exists")
)

type LevelDbStore struct {
	*leveldb.DB
}

func NewLevelDbStore(path string) (*LevelDbStore, error) {
	db, e := leveldb.OpenFile(path, nil)
	if e != nil {
		return nil, e
	}
	st := new(LevelDbStore)
	st.DB = db
	return st, nil
}

func (st *LevelDbStore) Set(key, val string) error {
	has, e := st.DB.Has([]byte(key), nil)
	if has {
		return ErrExists
	}
	if e != nil {
		return e
	}
	e = st.DB.Put([]byte(key), []byte(val), nil)
	return e
}

func (st *LevelDbStore) Has(key string) (bool, error) {
	return st.DB.Has([]byte(key), nil)
}

func (st *LevelDbStore) Get(key string) (string, error) {
	bts, e := st.DB.Get([]byte(key), nil)
	if e != nil {
		if e == errors.ErrNotFound {
			return "", ErrNotFound
		} else {
			return "", e
		}
	}
	return string(bts), nil
}
