package apidoc

type Header struct {
	// Name is the name of the header
	Name string `json:"name"`

	// Value is the value of the header
	Value string `json:"value"`
}
