package gopher

type Name string

func (n Name) String() string {
	return string(n)
}

func (n Name) IsValid() bool {
	return n.String() != ""
}
