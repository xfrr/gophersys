package gopher

type Metadata struct {
	m map[string]interface{}
}

func ParseMetadata(m map[string]interface{}) Metadata {
	return Metadata{m}
}

func (m Metadata) Get(key string) interface{} {
	if m.m == nil {
		return nil
	}

	return m.m[key]
}

func (m Metadata) AsMap() map[string]interface{} {
	if m.m == nil {
		m.m = make(map[string]interface{})
	}
	return m.m
}

func (m Metadata) IsEmpty() bool {
	if m.m == nil {
		return true
	}
	return len(m.m) == 0
}

func (m Metadata) merge(metadata Metadata) {
	if m.m == nil {
		m.m = make(map[string]interface{})
	}
	for k, v := range metadata.m {
		m.m[k] = v
	}
}

func (m Metadata) Compare(metadata Metadata) int {
	if m.IsEmpty() && metadata.IsEmpty() {
		return 0
	}

	if m.IsEmpty() {
		return -1
	}

	if metadata.IsEmpty() {
		return 1
	}

	if len(m.m) != len(metadata.m) {
		return -1
	}

	for k, v := range m.m {
		if metadata.Get(k) != v {
			return -1
		}
	}

	return 0
}
