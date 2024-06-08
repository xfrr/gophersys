// gopher_test package contains the unit tests for the gopher package.
package gopher_test

import (
	"testing"

	"github.com/xfrr/gophersys/internal/gopher"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		modifiers []gopher.Modifier
		wantErr   bool
	}{
		{
			name:      "valid gopher creation",
			id:        "valid_id",
			modifiers: []gopher.Modifier{gopher.WithName("Gopher"), gopher.WithUsername("gopher_username")},
			wantErr:   false,
		},
		{
			name:      "invalid gopher creation with invalid ID",
			id:        "",
			modifiers: []gopher.Modifier{gopher.WithName("Gopher"), gopher.WithUsername("gopher_username")},
			wantErr:   true,
		},
		{
			name:      "invalid gopher creation with invalid username",
			id:        "valid_id",
			modifiers: []gopher.Modifier{gopher.WithName("Gopher"), gopher.WithUsername("")},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := gopher.New(tt.id, tt.modifiers...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAggregate_Update(t *testing.T) {
	agg, err := gopher.New("valid_id", gopher.WithName("Gopher"), gopher.WithUsername("gopher_username"))
	if err != nil {
		t.Fatalf("failed to create gopher aggregate: %v", err)
	}

	tests := []struct {
		name      string
		agg       *gopher.Aggregate
		newName   gopher.Name
		newUser   gopher.Username
		newStatus gopher.Status
		newMeta   gopher.Metadata
		wantErr   bool
	}{
		{
			name:      "valid update",
			agg:       agg,
			newName:   gopher.Name("NewGopher"),
			newUser:   gopher.Username("new_username"),
			newStatus: gopher.StatusActive,
			newMeta:   gopher.Metadata{},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.agg.Update(tt.newName, tt.newUser, tt.newStatus, tt.newMeta)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Aggregate.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if tt.agg.Name() != tt.newName {
					t.Errorf("Aggregate.Update() name = %v, want %v", tt.agg.Name(), tt.newName)
				}

				if tt.agg.Username() != tt.newUser {
					t.Errorf("Aggregate.Update() username = %v, want %v", tt.agg.Username(), tt.newUser)
				}

				if tt.agg.Status() != tt.newStatus {
					t.Errorf("Aggregate.Update() status = %v, want %v", tt.agg.Status(), tt.newStatus)
				}

				if tt.agg.Metadata().Compare(tt.newMeta) != 0 {
					t.Errorf("Aggregate.Update() metadata = %v, want %v", tt.agg.Metadata(), tt.newMeta)
				}
			}
		})
	}
}

func TestAggregate_Properties(t *testing.T) {
	agg, err := gopher.New("valid_id", gopher.WithName("Gopher"), gopher.WithUsername("gopher_username"))
	if err != nil {
		t.Fatalf("failed to create gopher aggregate: %v", err)
	}

	tests := []struct {
		name     string
		method   string
		expected interface{}
	}{
		{
			name:     "ID",
			method:   "ID",
			expected: gopher.ID("valid_id"),
		},
		{
			name:     "Name",
			method:   "Name",
			expected: gopher.Name("Gopher"),
		},
		{
			name:     "Username",
			method:   "Username",
			expected: gopher.Username("gopher_username"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch tt.method {
			case "ID":
				result = agg.ID()
			case "Name":
				result = agg.Name()
			case "Username":
				result = agg.Username()
			default:
				t.Fatalf("unknown method %s", tt.method)
			}

			if result != tt.expected {
				t.Errorf("%s() = %v, want %v", tt.method, result, tt.expected)
			}
		})
	}
}
