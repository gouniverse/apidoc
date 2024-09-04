package blocks

// == TYPE ====================================================================

// GroupBlock represents a block with a title and content
// it is used to group blocks of similar functionality
type GroupBlock struct {
	baseBlock
	title   string
	content string
	blocks  []BlockInterface
}

var _ BlockInterface = (*GroupBlock)(nil)

// == CONSTRUCTOR =============================================================

func NewGroupBlock() *GroupBlock {
	return &GroupBlock{
		baseBlock: baseBlock{
			blockType: "group",
		},
	}
}

// == SETTERS & GETTERS =======================================================

func (block *GroupBlock) GetTitle() string {
	return block.title
}

func (block *GroupBlock) SetTitle(title string) *GroupBlock {
	block.title = title
	return block
}

func (block *GroupBlock) GetContent() string {
	return block.content
}

func (block *GroupBlock) SetContent(content string) *GroupBlock {
	block.content = content
	return block
}

func (block *GroupBlock) GetBlocks() []BlockInterface {
	return block.blocks
}

func (block *GroupBlock) SetBlocks(blocks []BlockInterface) *GroupBlock {
	block.blocks = blocks
	return block
}

func (block *GroupBlock) AddBlock(child BlockInterface) *GroupBlock {
	block.blocks = append(block.blocks, child)
	return block
}
