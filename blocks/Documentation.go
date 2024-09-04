package blocks

type Documentation struct {
	// Title is the title of the documentation
	title string

	// Blocks are the blocks of the documentation
	blocks []BlockInterface
}

func NewDocumentation() *Documentation {
	return &Documentation{}
}

func (doc *Documentation) GetTitle() string {
	return doc.title
}

func (doc *Documentation) SetTitle(title string) *Documentation {
	doc.title = title
	return doc
}

func (doc *Documentation) GetBlocks() []BlockInterface {
	return doc.blocks
}

func (doc *Documentation) SetBlocks(blocks []BlockInterface) *Documentation {
	doc.blocks = blocks
	return doc
}

func (doc *Documentation) AddBlock(block BlockInterface) *Documentation {
	doc.blocks = append(doc.blocks, block)
	return doc
}
