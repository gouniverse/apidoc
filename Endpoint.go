package apidoc

// == TYPE ====================================================================

// Endpoint represents an endpoint block
// it is used to describe an API endpoint to the documentation
type Endpoint struct {
	baseBlock

	// title is the title of the endpoint block
	title string

	// content is the content/description of the endpoint block
	// common markdown syntax is supported
	content string

	// method is the HTTP method of the endpoint
	// eg. GET, POST, PUT, PATCH, DELETE
	method string

	// path is the path of the endpoint
	// eg. /users
	path string

	// parameters are the parameters of the endpoint
	// eg. name, age
	parameters []Parameter

	// responses are the responses of the endpoint
	// eg. 200, 404
	responses []Response

	// 	// Consumes is the content types of the request body
	// 	// The content type is case-insensitive.
	// 	// Possible values: application/json, application/xml, text/plain, text/html, application/octet-stream
	// 	Consumes []string `json:"consumes"`

	// 	// Produces is the content types of the response body
	// 	// The content type is case-insensitive.
	// 	// Possible values: application/json, application/xml, text/plain, text/html, application/octet-stream
	// 	Produces []string `json:"produces"`

	// 	// Deprecated is true if the endpoint is deprecated
	// 	Deprecated bool `json:"deprecated"`

	// 	// Security is the security requirements of the endpoint (optional)
	// 	// Possible values: BasicAuth, OAuth, OAuth2, APIKey, JWT
	// 	Security []string `json:"security"`

	// 	Headers []Header `json:"headers"`
}

var _ BlockInterface = (*Endpoint)(nil) // ensure Endpoint implements BlockInterface

// == CONSTRUCTOR =============================================================

func NewEndpoint() *Endpoint {
	return &Endpoint{
		baseBlock: baseBlock{
			blockType: BLOCK_TYPE_ENDPOINT,
		},
	}
}

// == SETTERS & GETTERS =======================================================

func (block *Endpoint) GetTitle() string {
	return block.title
}

func (block *Endpoint) SetTitle(title string) *Endpoint {
	block.title = title
	return block
}

func (block *Endpoint) GetContent() string {
	return block.content
}

func (block *Endpoint) SetContent(content string) *Endpoint {
	block.content = content
	return block
}

// GetDescription is an alias for GetContent
func (block *Endpoint) GetDescription() string {
	return block.GetContent()
}

// SetDescription is an alias for SetContent
func (block *Endpoint) SetDescription(description string) *Endpoint {
	return block.SetContent(description)
}

func (block *Endpoint) GetMethod() string {
	return block.method
}

func (block *Endpoint) SetMethod(method string) *Endpoint {
	block.method = method
	return block
}

func (block *Endpoint) GetPath() string {
	return block.path
}

func (block *Endpoint) SetPath(path string) *Endpoint {
	block.path = path
	return block
}

func (block *Endpoint) GetParameters() []Parameter {
	return block.parameters
}

func (block *Endpoint) SetParameters(Parameters []Parameter) *Endpoint {
	block.parameters = Parameters
	return block
}

func (block *Endpoint) AddParameter(Parameter Parameter) *Endpoint {
	block.parameters = append(block.parameters, Parameter)
	return block
}

func (block *Endpoint) GetResponses() []Response {
	return block.responses
}

func (block *Endpoint) SetResponses(responses []Response) *Endpoint {
	block.responses = responses
	return block
}

func (block *Endpoint) AddResponse(response Response) *Endpoint {
	block.responses = append(block.responses, response)
	return block
}
