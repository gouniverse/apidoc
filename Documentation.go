package apidoc

import (
	"slices"

	"github.com/samber/lo"
)

// == TYPE ====================================================================

type Documentation struct {
	// title is the title of the documentation
	title string

	// logo is the logo of the documentation, can be a local or remote URL or HTML
	logo string

	// homeURL is the URL to the home page of the web application (eg. /)
	// if not set, the handlerURL will be used
	homeURL string

	// handlerURL is the URL to the handler serving the documentation (eg. /docs)
	handlerURL string

	theme string

	themesLight []struct {
		Key  string
		Name string
	}

	themesDark []struct {
		Key  string
		Name string
	}

	// blocks are the blocks of the documentation
	blocks []BlockInterface
}

// var _ BlockInterface = (*Documentation)(nil)

// == CONSTRUCTOR =============================================================

func NewDocumentation() *Documentation {
	return &Documentation{}
}

// == METHODS =================================================================

func (doc *Documentation) IsThemeDark() bool {
	keys := lo.Map(doc.themesDark, func(t struct{ Key, Name string }, _ int) string { return t.Key })
	return slices.Contains(keys, doc.theme)
}

func (doc *Documentation) IsThemeLight() bool {
	keys := lo.Map(doc.themesLight, func(t struct{ Key, Name string }, _ int) string { return t.Key })
	return slices.Contains(keys, doc.theme)
}

// == SETTERS & GETTERS =======================================================

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

// GetHandlerURL returns the URL to the handler serving the documentation
func (doc *Documentation) GetHandlerURL() string {
	return doc.handlerURL
}

// SetHandlerURL sets the URL to the handler serving the documentation
func (doc *Documentation) SetHandlerURL(handlerURL string) *Documentation {
	doc.handlerURL = handlerURL
	return doc
}

func (doc *Documentation) GetHomeURL() string {
	return doc.homeURL
}

func (doc *Documentation) SetHomeURL(homeURL string) *Documentation {
	doc.homeURL = homeURL
	return doc
}

func (doc *Documentation) GetLogo() string {
	return doc.logo
}

func (doc *Documentation) SetLogo(logo string) *Documentation {
	doc.logo = logo
	return doc
}

func (doc *Documentation) GetTheme() string {
	return doc.theme
}

func (doc *Documentation) SetTheme(theme string) *Documentation {
	doc.theme = theme
	return doc
}

func (doc *Documentation) GetThemesLight() []struct {
	Key  string
	Name string
} {
	return doc.themesLight
}

func (doc *Documentation) SetThemesLight(themesLight []struct {
	Key  string
	Name string
}) *Documentation {
	doc.themesLight = themesLight
	return doc
}

func (doc *Documentation) GetThemesDark() []struct {
	Key  string
	Name string
} {
	return doc.themesDark
}

func (doc *Documentation) SetThemesDark(themesDark []struct {
	Key  string
	Name string
}) *Documentation {
	doc.themesDark = themesDark
	return doc
}

func (doc *Documentation) GetTitle() string {
	return doc.title
}

func (doc *Documentation) SetTitle(title string) *Documentation {
	doc.title = title
	return doc
}
