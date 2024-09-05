package v1

type Example struct {
	// Name is the name of the example
	Name string `json:"name"`

	// ShortDescription is the short description of the example
	ShortDescription string `json:"short_description"`

	// LongDescription is the long description of the example
	LongDescription string `json:"long_description"`

	// Value is the example value
	Value string `json:"value"`
}
