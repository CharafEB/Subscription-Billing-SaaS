package model

import "database/sql"

type Data struct {
	DB *sql.DB
}

type Store struct{}

func NewStore(db *sql.DB) Store {
	if db == nil {
		panic("nil pointer passed to NewStore")
	}
	//data := &Data{DB: db}
	return Store{}
}
