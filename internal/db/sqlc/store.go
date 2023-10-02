package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Querier {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
