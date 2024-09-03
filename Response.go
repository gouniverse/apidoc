package apidoc

type Response struct {
	// Name is the name of the response
	Name string `json:"name"`

	// ShortDescription is the short description of the response
	ShortDescription string `json:"short_description"`

	// LongDescription is the long description of the response
	LongDescription string `json:"long_description"`

	// StatusCode is the status code of the response, e.g. 200, 400
	StatusCode string `json:"statusCode"`

	// Type is the type of the response, i.e application/json, text/plain, etc.
	Type string `json:"type"`

	// Example is the example of the response
	Example Example `json:"example"`
}
