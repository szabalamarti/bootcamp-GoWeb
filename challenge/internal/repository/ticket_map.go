package repository

import (
	"app/internal"
	"errors"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(loader internal.TicketLoader) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     make(map[int]internal.TicketAttributes),
		lastId: 0,
		loader: loader,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// loader
	loader internal.TicketLoader

	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	// lastId represents the last id of the ticket
	lastId int
}

// Load loads the tickets from the loader
func (r *RepositoryTicketMap) Load() error {
	// load the tickets from the loader
	t, err := r.loader.Load()
	if err != nil {
		return errors.New("error loading the tickets")
	}

	lastId := 0
	// save the tickets in the database
	for k, v := range t {
		// save the ticket in the database
		r.db[k] = v
		// update the last id
		if k > lastId {
			lastId = k
		}
		r.lastId = lastId
	}

	return nil
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get() (map[int]internal.TicketAttributes, error) {
	// load the tickets from the loader
	err := r.Load()
	if err != nil {
		return nil, errors.New("error loading the tickets")
	}
	// create a copy of the map
	t := make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return t, nil
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(country string) (map[int]internal.TicketAttributes, error) {
	// load the tickets from the loader
	err := r.Load()
	if err != nil {
		return nil, errors.New("error loading the tickets")
	}
	// create a copy of the map
	t := make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}
	return t, nil
}
