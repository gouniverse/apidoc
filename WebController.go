package apidoc

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// == CONTROLLER ===============================================================

// webController is the web controller that serves the documentation as a web page
type webController struct {
	documentation Documentation
	theme         string
}

var _ router.ControllerInterface = (*webController)(nil)

// == CONSTRUCTOR ==============================================================

// NewwebController creates a new web controller
func NewWebController(documentation Documentation) *webController {
	return &webController{
		documentation: documentation,
		theme:         documentation.GetTheme(), // default theme
	}
}

// == PUBLIC METHODS ===========================================================

// IdiomaticHandler is an idiomatic HTTP handler, which can be used to be served
// directly by most of the Go routers (i.e. standard mux, chi, etc.)
func (controller *webController) IdiomaticHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(controller.Handler(w, r)))
}

// Handler is a custom HTTP handler, which can be used by a specialized
// declarative router which expects a string as the return value
// (i.e. github.com/gouniverse/router)
func (controller *webController) Handler(w http.ResponseWriter, r *http.Request) string {
	if utils.Req(r, "theme", "") != "" {
		utils.CookieSet(w, r, "theme", utils.Req(r, "theme", ""), 60*60)
		// redirect to avoid form resubmission
		http.Redirect(w, r, controller.documentation.GetHandlerURL(), http.StatusFound)
	}

	userTheme := utils.CookieGet(r, "theme")

	if userTheme == "" {
		// if not set, the default theme will be used
		userTheme = controller.documentation.GetTheme()
	}

	controller.theme = userTheme

	return controller.webpage()
}

// == PRIVATE METHODS ==========================================================

// webpage generates the final web page with the documentation content
//
// 1. Sets the title, or defaults to `Api Documentation`
// 2. Generates the page layout (sidebar, main content, return to top button)
// 3. Identifies the theme and loads the corresponding CSS
// 4. Builds the final web page
func (controller *webController) webpage() string {
	title := controller.documentation.GetTitle()

	if title == "" {
		title = "Api Documentation"
	}

	pageLayout := hb.NewDiv().
		Class("container").
		Child(hb.NewDiv().
			Class("row").
			Child(controller.sidebar()).
			Child(controller.content()).
			Child(controller.buttonReturnToTop()))

	themeURL := lo.IfF(controller.theme == THEME_CERULEAN, func() string {
		return cdn.BootstrapCeruleanCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_COSMO, func() string {
		return cdn.BootstrapCosmoCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_CYBORG, func() string {
		return cdn.BootstrapCyborgCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_DARKLY, func() string {
		return cdn.BootstrapDarklyCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_FLATLY, func() string {
		return cdn.BootstrapFlatlyCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_JOURNAL, func() string {
		return cdn.BootstrapJournalCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_LITERA, func() string {
		return cdn.BootstrapLiteraCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_LUMEN, func() string {
		return cdn.BootstrapLumenCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_MINTY, func() string {
		return cdn.BootstrapMintyCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_LUX, func() string {
		return cdn.BootstrapLuxCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_MATERIA, func() string {
		return cdn.BootstrapMateriaCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_MORPH, func() string {
		return cdn.BootstrapMorphCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_PULSE, func() string {
		return cdn.BootstrapPulseCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_QUARTZ, func() string {
		return cdn.BootstrapQuartzCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SANDSTONE, func() string {
		return cdn.BootstrapSandstoneCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SIMPLEX, func() string {
		return cdn.BootstrapSimplexCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SKETCHY, func() string {
		return cdn.BootstrapSketchyCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SLATE, func() string {
		return cdn.BootstrapSlateCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SOLAR, func() string {
		return cdn.BootstrapSolarCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SPACELAB, func() string {
		return cdn.BootstrapSpacelabCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_SUPERHERO, func() string {
		return cdn.BootstrapSuperheroCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_UNITED, func() string {
		return cdn.BootstrapUnitedCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_VAPOR, func() string {
		return cdn.BootstrapVaporCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_YETI, func() string {
		return cdn.BootstrapYetiCss_5_3_3()
	}).ElseIfF(controller.theme == THEME_ZEPHYR, func() string {
		return cdn.BootstrapZephyrCss_5_3_3()
	}).
		Else("")

	highlightCss := "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/default.min.css"
	highlightJs := "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"
	highlightGoJs := "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/go.min.js"
	highlightJsScript := `hljs.highlightAll();`

	return hb.NewWebpage().
		SetTitle(title).
		Meta(hb.NewMeta().
			Name("viewport").
			Attr("content", "width=device-width, initial-scale=1")).
		AddStyleURL(cdn.BootstrapIconsCss_1_11_3()).
		AddStyleURL(cdn.BootstrapCss_5_3_3()).
		AddStyleURL(themeURL).
		AddStyleURL(highlightCss).
		AddScriptURL(cdn.BootstrapJs_5_3_3()).
		AddScriptURL(cdn.Htmx_2_0_0()).
		AddScriptURL(highlightJs).
		AddScriptURL(highlightGoJs).
		AddScript(highlightJsScript).
		Child(controller.navbar()).
		Child(pageLayout).
		ToHTML()
}

func (controller *webController) content() hb.TagInterface {
	title := strings.TrimSpace(controller.documentation.GetTitle())
	description := controller.markdownToHtml(controller.documentation.GetDescription())

	if title == "" {
		title = "API Documentation"
	}

	return hb.NewDiv().
		Class("col-9").
		Style("padding-top: 20px; padding-bottom: 20px;").
		Child(hb.NewHeading1().
			Style("margin-bottom: 20px;").
			HTML(title)).
		ChildIf(description != "", hb.NewHR()).
		ChildIf(description != "", hb.NewParagraph().
			Style("margin-bottom: 20px;").
			HTML(description)).
		Children(lo.Map(controller.documentation.GetBlocks(), func(block BlockInterface, _ int) hb.TagInterface {
			return controller.blockToTag(block)
		}))
}

func (controller *webController) blockToTag(block BlockInterface) hb.TagInterface {
	if block == nil {
		return nil
	}

	if block.GetType() == BLOCK_TYPE_GROUP {
		return controller.contentBlockGroup(block)
	}

	if block.GetType() == BLOCK_TYPE_ENDPOINT {
		return controller.contentBlockEndpoint(block)
	}

	if block.GetType() == BLOCK_TYPE_SECTION {
		return controller.contentBlockSection(block)
	}

	return hb.NewDiv().
		Style("margin-bottom: 20px;").
		HTML(`Unknown block: `).
		HTML(block.GetType())
}

func (controller *webController) contentBlockEndpoint(block BlockInterface) hb.TagInterface {
	endpoint := block.(*Endpoint)
	title := strings.TrimSpace(endpoint.GetTitle())
	content := controller.markdownToHtml(endpoint.GetContent())
	method := endpoint.GetMethod()
	path := endpoint.GetPath()
	parameters := endpoint.GetParameters()
	responses := endpoint.GetResponses()

	heading := hb.NewHeading5().
		Class("card-title").
		Child(hb.NewSpan().
			Class("badge text-bg-primary rounded-pill").
			HTML(method)).
		HTML(" ").
		HTML(path)

	card := hb.NewDiv().
		Class("card").
		Style("margin-bottom: 20px; margin-top: 20px;").
		Child(hb.NewDiv().
			Class("card-header").
			Child(heading)).
		Child(hb.NewDiv().
			Class("card-body").
			ChildIf(title != "", hb.NewHeading3().
				Style("margin-bottom: 20px;").
				HTML(title)).
			ChildIf(content != "", hb.NewDiv().
				Style("margin-bottom: 20px;").
				HTML(content)).
			Child(controller.contentEndpointParameters(parameters)).
			Child(controller.contentEndpointResponses(responses)))

	anchor := hb.NewHyperlink().
		Name(controller.buildAnchorEndpoint(method, path)).
		Href("#" + controller.buildAnchorEndpoint(method, path))

	return hb.NewWrap().
		Child(anchor).
		Child(card)
}

func (controller *webController) contentBlockGroup(block BlockInterface) hb.TagInterface {
	groupBlock := block.(*Group)

	title := strings.TrimSpace(groupBlock.GetTitle())
	description := controller.markdownToHtml(groupBlock.GetDescription())

	anchor := hb.NewHyperlink().
		Name(controller.buildAnchorGroup(groupBlock.GetTitle())).
		Href("#" + controller.buildAnchorGroup(groupBlock.GetTitle()))

	return hb.NewWrap().
		Child(hb.NewHR()).
		Child(anchor).
		ChildIf(title != "", hb.NewHeading2().
			HTML(title)).
		ChildIf(description != "", hb.NewDiv().
			Style("margin-bottom: 20px;").
			HTML(description)).
		Children(lo.Map(groupBlock.GetBlocks(), func(block BlockInterface, _ int) hb.TagInterface {
			return controller.blockToTag(block)
		}))
}

func (controller *webController) contentBlockSection(block BlockInterface) hb.TagInterface {
	sectionBlock := block.(*Section)

	title := strings.TrimSpace(sectionBlock.GetTitle())
	content := controller.markdownToHtml(sectionBlock.GetContent())

	anchor := hb.NewHyperlink().
		Name(controller.buildAnchorSection(title)).
		Href("#" + controller.buildAnchorSection(title))

	return hb.NewDiv().
		Style("margin-bottom: 20px; margin-top: 20px;").
		Child(hb.NewHR()).
		Child(anchor).
		ChildIf(title != "", hb.NewHeading3().
			HTML(title)).
		ChildIf(content != "", hb.NewDiv().
			Style("margin-bottom: 20px;").
			HTML(content))
}

func (controller *webController) contentEndpointParameters(parameters []Parameter) hb.TagInterface {
	parametersList := lo.Map(parameters, func(param Parameter, _ int) hb.TagInterface {
		return controller.contentEndpointParameter(param)
	})

	return hb.NewWrap().
		Child(hb.NewHeading4().
			HTML("Parameters")).
		ChildIfElse(len(parameters) > 0,
			hb.NewDiv().
				Class("accordion").
				ID("accordionExample").
				Style("margin-bottom: 20px;").
				Children(parametersList),
			hb.NewDiv().
				Style("margin-bottom: 20px;").
				HTML("No parameters"))
}

func (controller *webController) contentEndpointParameter(param Parameter) hb.TagInterface {
	name := param.GetName()
	title := strings.TrimSpace(param.GetTitle())
	description := controller.markdownToHtml(param.GetDescription())
	paramType := param.GetType()
	location := param.GetLocation()
	badgeType := hb.NewSpan().
		Class("badge text-bg-primary rounded-pill").
		HTMLIf(paramType != "", paramType)
	badgeRequired := hb.NewSpan().
		Class("badge rounded-pill").
		ClassIfElse(param.GetRequired(), "text-bg-danger", "text-bg-info").
		HTMLIfElse(param.GetRequired(), "required", "optional")

	id := "param-" + utils.StrSlugify(name, '-')

	return hb.NewDiv().
		Class("accordion-item").
		Child(hb.NewDiv().
			Class("accordion-header").
			Child(hb.NewButton().
				Class("accordion-button collapsed").
				Attr("type", "button").
				Attr("data-bs-toggle", "collapse").
				Attr("data-bs-target", "#"+id).
				Attr("aria-expanded", "false").
				Attr("aria-controls", id).
				Child(badgeType.Style("margin-right: 10px;")).
				HTML(name).
				Child(hb.NewSup().
					Child(badgeRequired.Style("margin-left: 10px;"))))).
		Child(hb.NewDiv().
			Class("accordion-collapse collapse").
			Attr("id", id).
			Attr("data-bs-parent", "#accordionExample").
			Attr("aria-labelledby", id).
			Child(hb.NewDiv().
				Class("accordion-body").
				Child(hb.NewDiv().
					Style("margin-bottom: 20px;").
					ChildIf(title != "", hb.NewHeading5().
						HTML(title)).
					ChildIf(description != "", hb.NewDiv().
						Style("margin-bottom: 20px;").
						HTML(description)).
					ChildIf(location != "", hb.NewDiv().
						Child(hb.NewSpan().
							Style("font-weight: bold;").
							HTML("Location: ")).
						Child(hb.NewSpan().
							HTML(location))))))
}

func (controller *webController) contentEndpointResponses(responses []Response) hb.TagInterface {
	responsesList := lo.Map(responses, func(resp Response, _ int) hb.TagInterface {
		return controller.contentEndpointResponse(resp)
	})

	return hb.NewWrap().
		Child(hb.NewHeading4().
			HTML("Responses")).
		ChildIfElse(len(responses) > 0,
			hb.NewDiv().
				Class("accordion").
				ID("accordionExample").
				Style("margin-bottom: 20px;").
				Children(responsesList),
			hb.NewDiv().
				Style("margin-bottom: 20px;").
				HTML("No responses"))
}

func (controller *webController) contentEndpointResponse(response Response) hb.TagInterface {
	statusCode := response.GetStatusCode()
	title := strings.TrimSpace(response.GetTitle())
	description := controller.markdownToHtml(response.GetDescription())
	example := response.GetExample()

	if title == "" {
		title = "Response " + statusCode
	}

	badgeStatusCode := hb.NewSpan().
		Class("badge text-bg-primary rounded-pill").
		HTML(lo.If(statusCode != "", statusCode).Else("200"))

	id := "response-" + utils.StrSlugify(title, '-')

	return hb.NewDiv().
		Class("accordion-item").
		Child(hb.NewDiv().
			Class("accordion-header").
			Child(hb.NewButton().
				Class("accordion-button collapsed").
				Attr("type", "button").
				Attr("data-bs-toggle", "collapse").
				Attr("data-bs-target", "#"+id).
				Attr("aria-expanded", "false").
				Attr("aria-controls", id).
				Child(badgeStatusCode.Style("margin-right: 10px;")).
				HTML(title))).
		Child(hb.NewDiv().
			Class("accordion-collapse collapse").
			Attr("id", id).
			Attr("data-bs-parent", "#accordionExample").
			Attr("aria-labelledby", id).
			Child(hb.NewDiv().
				Class("accordion-body").
				Child(hb.NewDiv().
					Style("margin-bottom: 20px;").
					ChildIf(description != "", hb.NewDiv().
						Style("margin-bottom: 20px;").
						HTML(description)).
					Child(hb.NewPRE().
						Child(hb.NewCode().
							Class(`language-` + example.GetSourceLanguage()).
							HTML(example.GetSourceCode()))))))
}

func (controller *webController) navbar() hb.TagInterface {
	title := controller.documentation.GetTitle()

	if title == "" {
		title = "API Documentation"
	}

	homeURL := controller.documentation.GetHomeURL()
	handlerURL := controller.documentation.GetHandlerURL()

	if handlerURL == "" {
		handlerURL = "#"
	}

	if homeURL == "" {
		homeURL = handlerURL
	}

	themeSwitch := lo.If(controller.documentation.GetHandlerURL() != "",
		// hb.NewI().
		// 	Class("bi bi-moon-fill").
		// 	Style("margin-left: 10px;"),
		hb.NewDiv().Class("float-end").
			Style("margin-left:10px;").
			Child(controller.navbarDropdownThemeSwitch()),
	).Else(
		hb.NewI().
			Class("bi bi-sun-fill").
			Style("margin-left: 10px;"),
	)

	homeLink := hb.NewHyperlink().
		Class("btn btn-primary float-end").
		Style("margin-left: 10px;").
		Href(homeURL).
		Title("Back to home").
		Child(hb.NewI().
			Class("bi bi-house"))

	bar := hb.NewNav().
		Class("navbar navbar-expand-lg navbar-dark bg-dark").
		Child(hb.NewDiv().Class("container d-block").
			Child(hb.NewHyperlink().
				Href(handlerURL).
				Class("navbar-brand ms-3").
				ChildIf(controller.documentation.GetLogo() != "", hb.NewImage().
					Class("d-inline-block align-text-top").
					Style("height: 24px; margin-right: 10px;").
					Src(controller.documentation.GetLogo()).
					Alt("Logo")).
				HTML(title)).
			Child(homeLink).
			ChildIf(themeSwitch != nil, hb.NewDiv().Class("float-end").
				Style("margin-left:10px;").
				Child(themeSwitch)))

	return bar
}

// themeButton generates a dropdown menu with light and dark themes.
//
// It checks if the current theme is dark and creates dropdown items for both light and dark themes.
// The dropdown items are created dynamically based on the themesLight and themesDark maps.
// The function returns a *hb.Tag that represents the generated dropdown menu.
func (controller *webController) navbarDropdownThemeSwitch() *hb.Tag {
	navbarTextColor := ""
	hasNavbarTextColor := lo.Ternary(navbarTextColor == "", false, true)
	buttonTheme := "" //d.navbarButtonThemeClass()

	handlerURL := controller.documentation.GetHandlerURL()
	themesLight := controller.documentation.GetThemesLight()
	themesDark := controller.documentation.GetThemesDark()
	isDark := controller.documentation.IsThemeDark(controller.theme)

	// Light Themes
	lightDropdownItems := lo.Map(themesLight, func(themeLight struct{ Key, Name string }, index int) hb.TagInterface {
		name := themeLight.Name
		active := lo.Ternary(themeLight.Key == controller.theme, " active", "")
		url := lo.Ternary(strings.Contains(handlerURL, "?"), handlerURL+"&theme="+themeLight.Key, handlerURL+"?theme="+themeLight.Key)

		return hb.NewLI().Children([]hb.TagInterface{
			hb.NewHyperlink().
				Class("dropdown-item"+active).
				Child(hb.NewI().Class("bi bi-sun").Style("margin-right:5px;")).
				HTML(name).
				Href(url).
				Attr("ref", "nofollow"),
		})
	})

	// Dark Themes
	darkDropdownItems := lo.Map(themesDark, func(themeDark struct{ Key, Name string }, index int) hb.TagInterface {
		name := themeDark.Name
		active := lo.Ternary(themeDark.Key == controller.theme, " active", "")
		url := lo.Ternary(strings.Contains(handlerURL, "?"), handlerURL+"&theme="+themeDark.Key, handlerURL+"?theme="+themeDark.Key)

		return hb.NewLI().Children([]hb.TagInterface{
			hb.NewHyperlink().
				Class("dropdown-item"+active).
				Child(hb.NewI().Class("bi bi-moon-stars-fill").Style("margin-right:5px;")).
				HTML(name).
				Href(url).
				Attr("ref", "nofollow"),
		})
	})

	return hb.NewDiv().
		Class("dropdown").
		Children([]hb.TagInterface{
			bs.Button().
				ID("buttonTheme").
				Title("Switch theme").
				Class(buttonTheme+" btn-warning dropdown-toggle").
				// Style("background:none;border:0px;").
				StyleIf(hasNavbarTextColor, "color:"+navbarTextColor).
				Data("bs-toggle", "dropdown").
				Children([]hb.TagInterface{
					lo.Ternary(isDark, hb.NewI().Class("bi bi-moon-stars-fill"), hb.NewI().Class("bi bi-sun")),
				}),
			hb.NewUL().Class(buttonTheme+" dropdown-menu dropdown-menu-dark").
				Children(lightDropdownItems).
				ChildIf(
					len(lo.Filter(darkDropdownItems, func(item hb.TagInterface, _ int) bool { return item != nil })) > 0 && len(lo.Filter(lightDropdownItems, func(item hb.TagInterface, _ int) bool { return item != nil })) > 0,
					hb.NewLI().Children([]hb.TagInterface{
						hb.NewHR().Class("dropdown-divider"),
					}),
				).
				Children(darkDropdownItems),
		})
}

// sidebar returns the sidebar of the documentation
func (controller *webController) sidebar() hb.TagInterface {
	groupBlocks := lo.Filter(controller.documentation.GetBlocks(), func(block BlockInterface, _ int) bool {
		return block.GetType() == BLOCK_TYPE_GROUP
	})

	sectionsBlocks := lo.Filter(controller.documentation.GetBlocks(), func(block BlockInterface, _ int) bool {
		return block.GetType() == BLOCK_TYPE_SECTION
	})

	groupCards := controller.sidebarGroupCards(groupBlocks)
	apiInfoCard := controller.sidebarApiInfoCard(sectionsBlocks)

	return hb.NewDiv().
		Class("col-3").
		Style("padding-top: 20px;").
		Child(apiInfoCard).
		Children(groupCards)
}

func (controller *webController) sidebarApiInfoCard(sections []BlockInterface) hb.TagInterface {
	listInfos := hb.NewUL().
		Class("list-group list-group-flush").
		Child(hb.NewLI().
			Class("list-group-item").
			Child(hb.NewHyperlink().
				Href("#").
				HTML("Home"))).
		Children(lo.Map(sections, func(info BlockInterface, _ int) hb.TagInterface {
			return hb.NewLI().
				Class("list-group-item").
				Child(hb.NewHyperlink().
					Href("#" + controller.buildAnchorSection(info.GetTitle())).
					HTML(info.GetTitle()))
		}))

	return hb.NewDiv().
		Class("card").
		Style("margin-bottom: 20px;").
		Child(hb.NewDiv().
			Class("card-header").
			Child(hb.NewHeading5().
				Class("card-title").
				Style("margin-bottom: 0;").
				HTML("API Info"))).
		Child(hb.NewDiv().
			Class("card-body").
			Child(listInfos))
}

// sidebarGroupCards returns the group cards of the documentation
func (controller *webController) sidebarGroupCards(groupBlocks []BlockInterface) []hb.TagInterface {
	return lo.Map(groupBlocks, func(block BlockInterface, _ int) hb.TagInterface {
		groupBlock := block.(*Group)

		title := strings.TrimSpace(block.GetTitle())

		endpoints := lo.Filter(groupBlock.GetBlocks(), func(block BlockInterface, _ int) bool {
			return block.GetType() == BLOCK_TYPE_ENDPOINT
		})

		listEndpoints := hb.NewUL().
			Class("list-group list-group-flush").
			Children(lo.Map(endpoints, func(endpoint BlockInterface, _ int) hb.TagInterface {
				Endpoint := endpoint.(*Endpoint)
				return hb.NewLI().
					Class("list-group-item").
					Child(hb.NewHyperlink().
						Href("#" + controller.buildAnchorEndpoint(Endpoint.GetMethod(), Endpoint.GetPath())).
						HTML(endpoint.GetTitle()))
			}))

		sections := lo.Filter(groupBlock.GetBlocks(), func(block BlockInterface, _ int) bool {
			return block.GetType() == BLOCK_TYPE_SECTION
		})

		listSections := hb.NewUL().
			Class("list-group list-group-flush").
			Children(lo.Map(sections, func(info BlockInterface, _ int) hb.TagInterface {
				return hb.NewLI().
					Class("list-group-item").
					Child(hb.NewHyperlink().
						Href("#" + controller.buildAnchorSection(info.GetTitle())).
						HTML(info.GetTitle()))
			}))

		return hb.NewDiv().
			Class("card").
			Style("margin-bottom: 20px;").
			Child(hb.NewDiv().
				Class("card-header").
				Child(hb.NewHeading5().
					Class("card-title").
					Style("margin-bottom: 0;").
					HTML(title))).
			Child(hb.NewDiv().
				Class("card-body").
				ChildIf(len(endpoints) > 0, listEndpoints).
				ChildIf(len(sections) > 0, listSections))

	})
}

func (controller *webController) buttonReturnToTop() hb.TagInterface {
	script := hb.NewScript(`
window.onscroll = function() {scrollFunction()};

function scrollFunction() {
	if (document.body.scrollTop > 20 || document.documentElement.scrollTop > 20) {
		document.getElementById("btn-back-to-top").style.display = "block";
	} else {
		document.getElementById("btn-back-to-top").style.display = "none";
	}
}

document.getElementById("btn-back-to-top").addEventListener("click", function() {
	document.body.scrollTop = 0;
	document.documentElement.scrollTop = 0;
});
		`)

	return hb.NewWrap().
		Child(hb.NewButton().
			ID("btn-back-to-top").
			Type(hb.TYPE_BUTTON).
			Class("btn btn-danger btn-floating btn-lg rounded-circle").
			Style("position: fixed; bottom: 20px; right: 20px; display: none; width: 50px; height: 50px; padding: 0px; box-shadow: 0 8px 9px -4px rgba(209, 72, 95, 0.3), 0 4px 18px 0 rgba(209, 72, 95, 0.2);").
			Child(hb.NewI().
				Class("bi bi-arrow-up-short").
				Style("font-size: 24px;")).
			Title("Back to top")).
		Child(script)

}

// markdownToHtml converts a markdown text to html
//
// 1. the text is trimmed of any white spaces
// 2. if the text is empty, it returns an empty string
// 3. the text is converted to html using the goldmark library
func (controller *webController) markdownToHtml(text string) string {
	text = strings.TrimSpace(text)

	if text == "" {
		return ""
	}

	var buf bytes.Buffer
	md := goldmark.New(
		// goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			// html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert([]byte(text), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}

func (controller *webController) buildAnchorEndpoint(method string, path string) string {
	slug := utils.StrSlugify(path, '-')
	return "#endpoint-" + method + "-" + slug
}

func (controller *webController) buildAnchorGroup(title string) string {
	slug := utils.StrSlugify(title, '-')
	return "#group-" + slug
}

func (controller *webController) buildAnchorSection(title string) string {
	slug := utils.StrSlugify(title, '-')
	return "#section-" + slug
}
