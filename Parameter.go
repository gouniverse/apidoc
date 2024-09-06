package apidoc

// == TYPE ====================================================================

type Parameter struct {
	baseBlock

	// name is the name of the parameter as expected by the endpoint
	name string

	// title is the user friendly title of the parameter
	title string

	// content is the description of the parameter
	// common markdown syntax is supported
	content string

	// dataType is the data type of the parameter
	// eg. string, number, boolean, array, object
	dataType string

	// location is the location of the parameter
	// eg. path, query, header, cookie
	location string

	// required is the required status of the parameter
	// eg. true, false
	required bool

	// !!! WARNING !!! do we need this?
	// review this in the future (2025-06-01) to see if we need it
	// default_         string
	// example          Example
	// enum            []string
	// format          string
	// pattern         string
	// minimum         string
	// maximum         string
	// minLength       string
	// maxLength       string
	// deprecated      bool
	// allowEmptyValue bool
	// style           string
}

var _ BlockInterface = (*Parameter)(nil) // ensure Parameter implements BlockInterface

// == CONSTRUCTOR =============================================================

func NewParameter() *Parameter {
	return &Parameter{
		baseBlock: baseBlock{
			blockType: BLOCK_TYPE_PARAMETER,
		},
	}
}

// == SETTERS & GETTERS =======================================================

func (block *Parameter) GetContent() string {
	return block.content
}

func (block *Parameter) SetContent(content string) *Parameter {
	block.content = content
	return block
}

// GetDescription is an alias for GetContent
func (block *Parameter) GetDescription() string {
	return block.GetContent()
}

// SetDescription is an alias for SetContent
func (block *Parameter) SetDescription(content string) *Parameter {
	return block.SetContent(content)
}

func (block *Parameter) GetName() string {
	return block.name
}

func (block *Parameter) SetName(name string) *Parameter {
	block.name = name
	return block
}

func (block *Parameter) GetTitle() string {
	return block.title
}

func (block *Parameter) SetTitle(title string) *Parameter {
	block.title = title
	return block
}

func (block *Parameter) GetType() string {
	return block.dataType
}

func (block *Parameter) SetType(dataType string) *Parameter {
	block.dataType = dataType
	return block
}

func (block *Parameter) GetLocation() string {
	return block.location
}

func (block *Parameter) SetLocation(location string) *Parameter {
	block.location = location
	return block
}

func (block *Parameter) GetRequired() bool {
	return block.required
}

func (block *Parameter) SetRequired(required bool) *Parameter {
	block.required = required
	return block
}
