package gopher

import "testing"

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want string
	}{
		{name: "StatusActive", s: StatusActive, want: "active"},
		{name: "StatusInactive", s: StatusInactive, want: "inactive"},
		{name: "StatusSuspended", s: StatusSuspended, want: "suspended"},
		{name: "StatusDeleted", s: StatusDeleted, want: "deleted"},
		{name: "StatusInvalid", s: Status(0), want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Status.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsValid(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{name: "StatusActive", s: StatusActive, want: true},
		{name: "StatusInactive", s: StatusInactive, want: true},
		{name: "StatusSuspended", s: StatusSuspended, want: true},
		{name: "StatusDeleted", s: StatusDeleted, want: true},
		{name: "StatusInvalid", s: Status(0), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValid(); got != tt.want {
				t.Errorf("Status.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStatus(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Status
	}{
		{name: "StatusActive", args: args{s: "active"}, want: StatusActive},
		{name: "StatusInactive", args: args{s: "inactive"}, want: StatusInactive},
		{name: "StatusSuspended", args: args{s: "suspended"}, want: StatusSuspended},
		{name: "StatusDeleted", args: args{s: "deleted"}, want: StatusDeleted},
		{name: "StatusInvalid", args: args{s: "invalid"}, want: Status(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStatus(tt.args.s); got != tt.want {
				t.Errorf("ParseStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
