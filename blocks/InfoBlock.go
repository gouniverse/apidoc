package blocks

// == TYPE ====================================================================

// InfoBlock represents a block with a title and content
// it is used to add an information section to the documentation
type InfoBlock struct {
	baseBlock
	title   string
	content string
}

var _ BlockInterface = (*InfoBlock)(nil)

// == CONSTRUCTOR =============================================================

func NewInfoBlock() *InfoBlock {
	return &InfoBlock{
		baseBlock: baseBlock{
			blockType: TYPE_INFO,
		},
	}
}

// == SETTERS & GETTERS =======================================================

func (block *InfoBlock) GetTitle() string {
	return block.title
}

func (block *InfoBlock) SetTitle(title string) *InfoBlock {
	block.title = title
	return block
}

func (block *InfoBlock) GetContent() string {
	return block.content
}

func (block *InfoBlock) SetContent(content string) *InfoBlock {
	block.content = content
	return block
}
