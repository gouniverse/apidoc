package apidoc

type baseBlock struct {
	blockType string
}

func (block *baseBlock) GetType() string {
	return block.blockType
}

func (block *baseBlock) SetType(blockType string) *baseBlock {
	block.blockType = blockType
	return block
}
