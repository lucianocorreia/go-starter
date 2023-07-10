package handlers

import database "github.com/lucianocorreia/go-starter/database/sqlc"

// Handlers is the application handlers struct
type Handlers struct {
	store database.Store
}

// NewHandlers returns a new Handlers struct
func NewHandlers(store database.Store) *Handlers {
	return &Handlers{
		store: store,
	}
}
