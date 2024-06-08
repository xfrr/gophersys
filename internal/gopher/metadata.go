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
