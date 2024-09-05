package v1

type Parameter struct {
	// Name is the name of the parameter
	Name string `json:"name"`

	// Type is the type of the parameter, e.g. string, int, bool
	Type string `json:"type"`

	// ShortDescription is the short description of the parameter
	ShortDescription string `json:"short_description"`

	// LongDescription is the long description of the parameter
	LongDescription string `json:"long_description"`

	// Location is the location of the parameter
	// Possible values: query, header, path, cookie
	Location string `json:"location"`

	// Required is true if the parameter is required
	Required bool `json:"required"`

	// Default is the default value of the parameter
	Default string `json:"default"`

	// Example is the example of the parameter (used for documentation)
	Example Example `json:"example"`

	// Enum is the enum of the parameter
	Enum []string `json:"enum"`

	// Format is the format of the parameter
	Format string `json:"format"`

	// Pattern is the pattern of the parameter
	Pattern string `json:"pattern"`

	// Minimum is the minimum value of the parameter
	Minimum string `json:"minimum"`

	// Maximum is the maximum value of the parameter
	Maximum string `json:"maximum"`

	// MinLength is the minimum length of the parameter
	MinLength string `json:"minLength"`

	// MaxLength is the maximum length of the parameter
	MaxLength string `json:"maxLength"`

	// Deprecated is true if the parameter is deprecated
	Deprecated bool `json:"deprecated"`
}
