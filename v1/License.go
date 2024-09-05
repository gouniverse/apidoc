package v1

type License struct {
	// Name is the name of the license
	Name string `json:"name"`

	// URL is the URL of the license
	URL string `json:"url"`

	// Text is the text of the license
	Text string `json:"text"`
}
