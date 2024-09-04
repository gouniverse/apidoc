package blocks

type BlockInterface interface {
	GetType() string
	GetTitle() string
	GetContent() string
}
