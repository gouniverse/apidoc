package apidoc

type BlockInterface interface {
	// GetType returns the type of the block
	GetType() string

	// GetTitle returns the title of the block
	GetTitle() string

	// GetContent returns the content of the block
	GetContent() string
}
