package blocks

type EndpointBlock struct {
	baseBlock
	title      string
	content    string
	method     string
	path       string
	parameters []Parameter
	responses  []Response
}

func NewEndpointBlock() *EndpointBlock {
	return &EndpointBlock{
		baseBlock: baseBlock{
			blockType: TYPE_ENDPOINT,
		},
	}
}

func (block *EndpointBlock) GetTitle() string {
	return block.title
}

func (block *EndpointBlock) SetTitle(title string) *EndpointBlock {
	block.title = title
	return block
}

func (block *EndpointBlock) GetContent() string {
	return block.content
}

func (block *EndpointBlock) SetContent(content string) *EndpointBlock {
	block.content = content
	return block
}

func (block *EndpointBlock) GetMethod() string {
	return block.method
}

func (block *EndpointBlock) SetMethod(method string) *EndpointBlock {
	block.method = method
	return block
}

func (block *EndpointBlock) GetPath() string {
	return block.path
}

func (block *EndpointBlock) SetPath(path string) *EndpointBlock {
	block.path = path
	return block
}

func (block *EndpointBlock) GetParameters() []Parameter {
	return block.parameters
}

func (block *EndpointBlock) SetParameters(Parameters []Parameter) *EndpointBlock {
	block.parameters = Parameters
	return block
}

func (block *EndpointBlock) AddParameter(Parameter Parameter) *EndpointBlock {
	block.parameters = append(block.parameters, Parameter)
	return block
}

func (block *EndpointBlock) GetResponses() []Response {
	return block.responses
}

func (block *EndpointBlock) SetResponses(responses []Response) *EndpointBlock {
	block.responses = responses
	return block
}

func (block *EndpointBlock) AddResponse(response Response) *EndpointBlock {
	block.responses = append(block.responses, response)
	return block
}
