package gopher

import "context"

// Repository manages the persistence of the Gopher aggregate
type Repository interface {
	// Save persists the Gopher entity in the repository
	//
	// - When already exists, it updates the Gopher properties.
	//
	// - When does not exist, it creates a new Gopher entity.
	//
	// - When any other error occurs, it returns the error.
	Save(ctx context.Context, g *Aggregate) error

	// Exists checks if the Gopher entity exists
	//
	// - When filters are provided, it checks if the Gopher entity matches
	// all the given properties.
	//
	// - When no filters are provided, it returns an error.
	//
	// - When the Gopher entity exists, it returns true.
	//
	// - When the Gopher entity does not exist, it returns false.
	//
	// - When any other error occurs, it returns the error.
	Exists(ctx context.Context, filters ...Filter) (bool, error)

	// Get retrieves the Gopher entity that matches the given filters
	//
	// - When filters are provided, it retrieves the Gopher entity that matches
	// all the given properties.
	//
	// - When no filters are provided, it returns an error.
	//
	// - When the Gopher entity exists, it returns the Gopher entity.
	//
	// - When the Gopher entity does not exist, it returns an error.
	//
	// - When any other error occurs, it returns the error.
	Get(ctx context.Context, filters ...Filter) (*Aggregate, error)

	// Search retrieves all Gopher entities that match the given query
	//
	// - When the query is empty, it retrieves all Gopher entities.
	//
	// - When the query is not empty, it retrieves all Gopher entities that match
	// the query properties.
	//
	// - When any other error occurs, it returns the error.
	Search(ctx context.Context, query SearchQuery) ([]*Aggregate, error)

	// Delete deletes the Gopher entity that matches the given ID
	//
	// - When the Gopher entity exists, it deletes the Gopher entity.
	//
	// - When the Gopher entity does not exist, does nothing.
	//
	// - When any other error occurs, it returns the error.
	Delete(ctx context.Context, id ID) error
}

// SearchQuery represents a query for Gopher entities
type SearchQuery struct {
	// Cursor is the unique identifier where the search starts
	// If empty, the search starts from the beginning
	Cursor string

	// PerPage is the number of Gophers to retrieve per page
	// If zero, the default value is used
	// If negative, all Gophers are retrieved
	PerPage int

	// Names are the Gopher names to filter
	Names []string

	// Statuses are the Gopher statuses to filter
	Statuses []Status
}

// Filters contains a set of properties to filter when
// getting gophers
type Filters struct {
	// ID is the Gopher ID
	ID ID

	// Username is the Gopher username
	Username Username
}

// Filter is a function to filter Gophers
type Filter func(*Filters)

// ByID filters Gophers by ID
func ByID(id ID) Filter {
	return func(f *Filters) {
		f.ID = id
	}
}

// ByUsername filters Gophers by username
func ByUsername(username Username) Filter {
	return func(f *Filters) {
		f.Username = username
	}
}
