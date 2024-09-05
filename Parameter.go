package apidoc

type Parameter struct {
	name     string
	title    string
	content  string
	dataType string
	location string
	required bool
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

func NewParameter() *Parameter {
	return &Parameter{}
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

func (block *Parameter) GetContent() string {
	return block.content
}

func (block *Parameter) SetContent(content string) *Parameter {
	block.content = content
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
