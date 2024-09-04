package blocks

type Example struct {
	title          string
	content        string
	sourceLanguage string
	sourceCode     string
}

func NewExample() *Example {
	return &Example{}
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
