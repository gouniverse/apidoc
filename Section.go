package apidoc

// == TYPE ====================================================================

// Section represents a section block with a title and content
// it is used to add an information section to the documentation
type Section struct {
	baseBlock

	// title is the title of the info block
	// eg. Introduction, Getting Started, ...
	title string

	// content is the content/description of the info block
	// common markdown syntax is supported
	content string
}

var _ BlockInterface = (*Section)(nil) // ensure Section implements BlockInterface

// == CONSTRUCTOR =============================================================

func NewSection() *Section {
	return &Section{
		baseBlock: baseBlock{
			blockType: BLOCK_TYPE_SECTION,
		},
	}
}

// == SETTERS & GETTERS =======================================================

func (block *Section) GetTitle() string {
	return block.title
}

func (block *Section) SetTitle(title string) *Section {
	block.title = title
	return block
}

func (block *Section) GetContent() string {
	return block.content
}

func (block *Section) SetContent(content string) *Section {
	block.content = content
	return block
}

// GetDescription is an alias for GetContent
func (block *Section) GetDescription() string {
	return block.GetContent()
}

// SetDescription is an alias for SetContent
func (block *Section) SetDescription(description string) *Section {
	return block.SetContent(description)
}
