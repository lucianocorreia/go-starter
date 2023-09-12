package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewStore creates a new store
func NewStore(pool *pgxpool.Pool) Store {
	return SQLStore{
		Queries:  New(pool),
		connPool: pool,
	}
}

// execTx executes a function within a database transaction
// func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
