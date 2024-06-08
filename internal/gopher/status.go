package gopher

type Status int

const (
	// StatusActive represents an active user
	StatusActive Status = iota + 1
	// StatusInactive represents an inactive user
	StatusInactive
	// StatusSuspended represents a suspended user
	StatusSuspended
	// StatusDeleted represents a deleted user
	StatusDeleted
)

func (s Status) String() string {
	return [...]string{"", "active", "inactive", "suspended", "deleted"}[s]
}

func (s Status) IsValid() bool {
	return s > 0 && s < 5
}

func StatusFromString(s string) Status {
	switch s {
	case "active":
		return StatusActive
	case "inactive":
		return StatusInactive
	case "suspended":
		return StatusSuspended
	case "deleted":
		return StatusDeleted
	default:
		return 0
	}
}
