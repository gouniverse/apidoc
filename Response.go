package apidoc

// == TYPE ====================================================================

// Response represents a block with a title and content
// it is used to add a response section to an endpoint
type Response struct {
	baseBlock

	// title is the title of the response
	// eg. Success, Error
	title string

	// content is the content/description of the response
	// common markdown syntax is supported
	content string

	// contentType is the content type of the response e.g. application/json
	// best use the predefined content types constants defined in this package
	contentType string

	// statusCode is the status code of the response
	// eg. 200
	statusCode string

	// example is the example of the response
	example Example
}

var _ BlockInterface = (*Response)(nil) // ensure Response implements BlockInterface

// == CONSTRUCTOR =============================================================

func NewResponse() *Response {
	return &Response{}
}

// == SETTERS & GETTERS =======================================================

func (block *Response) GetContent() string {
	return block.content
}

func (block *Response) SetContent(content string) *Response {
	block.content = content
	return block
}

// GetDescription is an alias for GetContent
func (block *Response) GetDescription() string {
	return block.GetContent()
}

// SetDescription is an alias for SetContent
func (block *Response) SetDescription(description string) *Response {
	return block.SetContent(description)
}

func (block *Response) GetTitle() string {
	return block.title
}

func (block *Response) SetTitle(title string) *Response {
	block.title = title
	return block
}

func (block *Response) GetStatusCode() string {
	return block.statusCode
}

func (block *Response) SetStatusCode(statusCode string) *Response {
	block.statusCode = statusCode
	return block
}

func (block *Response) GetContentType() string {
	return block.contentType
}

func (block *Response) SetContentType(contentType string) *Response {
	block.contentType = contentType
	return block
}

func (block *Response) GetExample() Example {
	return block.example
}

func (block *Response) SetExample(example Example) *Response {
	block.example = example
	return block
}
