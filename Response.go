package apidoc

type Response struct {
	title   string
	content string

	// ContentType is the content type of the response e.g. application/json
	contentType string

	statusCode string

	example Example
}

func NewResponse() *Response {
	return &Response{}
}

func (block *Response) GetTitle() string {
	return block.title
}

func (block *Response) SetTitle(title string) *Response {
	block.title = title
	return block
}

func (block *Response) GetContent() string {
	return block.content
}

func (block *Response) SetContent(content string) *Response {
	block.content = content
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
