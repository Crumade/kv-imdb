package storage

import "errors"

func (db *Database) Get(key string) (string, error) {
	val, ok := db.Data[key]
	if ok {
		return val, nil
	}
	return val, errors.New("not found")
}

func (db *Database) Set(key, value string) {
	db.Data[key] = value
}

func (db *Database) Delete(key string) {
	delete(db.Data, key)
}
