package gopher

type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) IsValid() bool {
	return i != ""
}
