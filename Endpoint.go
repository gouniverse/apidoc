package apidoc

type Endpoint struct {
	// Method is the HTTP method of the endpoint (e.g. GET, POST, PUT, DELETE)
	// The method is case-insensitive.
	Method string `json:"method"`

	// Path is the path of the endpoint
	// The path MUST start with a slash (/).
	// The path is appended to the API URL in order to construct the full URL.
	// Path parameters are allowed, (i.e. /user/{user_id})
	Path string `json:"path"`

	// ShortDescription is the short description of the endpoint (optional)
	ShortDescription string `json:"short_description"`

	// LongDescription is the long description of the endpoint (optional)
	LongDescription string `json:"long_description"`

	// Parameters are the parameters of the endpoint (optional)
	Parameters []Parameter `json:"parameters"`

	// Responses are the responses of the endpoint
	Responses []Response `json:"responses"`

	// Consumes is the content types of the request body
	// The content type is case-insensitive.
	// Possible values: application/json, application/xml, text/plain, text/html, application/octet-stream
	Consumes []string `json:"consumes"`

	// Produces is the content types of the response body
	// The content type is case-insensitive.
	// Possible values: application/json, application/xml, text/plain, text/html, application/octet-stream
	Produces []string `json:"produces"`

	// Deprecated is true if the endpoint is deprecated
	Deprecated bool `json:"deprecated"`

	// Security is the security requirements of the endpoint (optional)
	// Possible values: BasicAuth, OAuth, OAuth2, APIKey, JWT
	Security []string `json:"security"`

	Headers []Header `json:"headers"`
}
