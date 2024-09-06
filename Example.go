package apidoc

// == TYPE ====================================================================

// Example represents an example block with a source code example
// it is used to add an example section to an endpoint
type Example struct {
	baseBlock

	// title is the title of the example block
	// eg. Source Code
	// !!! WARNING !!! do we need this?
	// review this in the future (2025-06-01) to see if we need it
	title string

	// content is the description of the example block
	// common markdown syntax is supported
	// !!! WARNING !!! do we need this?
	// review this in the future (2025-06-01) to see if we need it
	content string

	// sourceLanguage is the language of the source code
	// eg. go, javascript, python
	sourceLanguage string

	// sourceCode is the source code of the example
	// eg. func main() {
	// 		fmt.Println("Hello, World!")
	// }
	sourceCode string
}

var _ BlockInterface = (*Example)(nil) // ensure Example implements BlockInterface

func NewExample() *Example {
	return &Example{
		baseBlock: baseBlock{
			blockType: BLOCK_TYPE_EXAMPLE,
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

// GetDescription is an alias for GetContent
func (block *Example) GetDescription() string {
	return block.GetContent()
}

// SetDescription is an alias for SetContent
func (block *Example) SetDescription(description string) *Example {
	block.SetContent(description)
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
