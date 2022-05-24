package database

import "github.com/fajarnugraha37/go-rest-api/database/repository"

type Queries struct {
	*repository.BookQueries
}

func OpenConnection() (*Queries, error) {
	return &Queries{
		BookQueries: &repository.BookQueries{DB: DB},
	}, nil
}
