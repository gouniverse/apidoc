package apidoc

// == TYPE ====================================================================

// Group represents a block with a title and content
// it is used to group blocks of similar functionality
type Group struct {
	baseBlock

	// title is the title of the group block
	title string

	// content is the content/description of the group block
	// common markdown syntax is supported
	content string

	// blocks are the blocks of the group block
	// can be one of the following:
	// - InfoBlock (individual block with a title and content)
	blocks []BlockInterface
}

var _ BlockInterface = (*Group)(nil) // ensure Group implements BlockInterface

// == CONSTRUCTOR =============================================================

func NewGroup() *Group {
	return &Group{
		baseBlock: baseBlock{
			blockType: "group",
		},
	}
}

// == SETTERS & GETTERS =======================================================

func (block *Group) GetTitle() string {
	return block.title
}

func (block *Group) SetTitle(title string) *Group {
	block.title = title
	return block
}

func (block *Group) GetContent() string {
	return block.content
}

func (block *Group) SetContent(content string) *Group {
	block.content = content
	return block
}

// GetDescription is an alias for GetContent
func (block *Group) GetDescription() string {
	return block.GetContent()
}

// SetDescription is an alias for SetContent
func (block *Group) SetDescription(content string) *Group {
	return block.SetContent(content)
}

func (block *Group) GetBlocks() []BlockInterface {
	return block.blocks
}

func (block *Group) SetBlocks(blocks []BlockInterface) *Group {
	block.blocks = blocks
	return block
}

func (block *Group) AddBlock(child BlockInterface) *Group {
	block.blocks = append(block.blocks, child)
	return block
}
