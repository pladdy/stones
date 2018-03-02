package stones

// Bundle is a STIX bundle: a collection of STIX objects
type Bundle struct {
	Type        string   `json:"type"`
	ID          StixID   `json:"id"`
	SpecVersion string   `json:"spec_version"`
	Objects     []string `json:"objects"`
}

// NewBundle returns a STIX bundle object
func NewBundle() (Bundle, error) {
	b := Bundle{}
	id, err := NewStixID()

	b.ID = id
	return b, err
}
