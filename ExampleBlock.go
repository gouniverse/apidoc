package apidoc

// == TYPE ====================================================================

// ExampleBlock represents a block with a source code example
// it is used to add an example section to the documentation
type Example struct {
	baseBlock
	title          string
	content        string
	sourceLanguage string
	sourceCode     string
}

func NewExample() *Example {
	return &Example{
		baseBlock: baseBlock{
			blockType: TYPE_EXAMPLE,
		},
	}
}

func (block *Example) GetTitle() string {
	return block.title
}

func (block *Example) SetTitle(title string) *Example {
	block.title = title
	return block
}

func (block *Example) GetContent() string {
	return block.content
}

func (block *Example) SetContent(content string) *Example {
	block.content = content
	return block
}

func (block *Example) GetSourceLanguage() string {
	return block.sourceLanguage
}

func (block *Example) SetSourceLanguage(sourceLanguage string) *Example {
	block.sourceLanguage = sourceLanguage
	return block
}

func (block *Example) GetSourceCode() string {
	return block.sourceCode
}

func (block *Example) SetSourceCode(sourceCode string) *Example {
	block.sourceCode = sourceCode
	return block
}
