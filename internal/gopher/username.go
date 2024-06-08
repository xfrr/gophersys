package gopher

// Username represents a gopher username
type Username string

// String returns the username as a string
func (n Username) String() string {
	return string(n)
}

// IsValid returns true if the username is valid
func (u Username) IsValid() bool {
	return len(u) >= 1 && len(u) <= 320
}
