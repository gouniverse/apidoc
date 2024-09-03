package apidoc

type Definition struct {
	// Title is the title of the API
	Title string `json:"title"`

	// ShortDescription is the short description of the API
	ShortDescription string `json:"short_description"`

	// LongDescription is the long description of the API
	LongDescription string `json:"long_description"`

	// ServerURL is the base URL of the server (e.g. https://api.example.com)
	ServerURL string `json:"serverURL"`

	// BasePath is the base path of the API (e.g. /api)
	BasePath string `json:"basePath"`

	// Host is the hostname of the server (eg. api.example.com)
	Host string `json:"host"`

	// Schemes are the schemes supported by the API (e.g. http, https)
	Schemes []string `json:"schemes"`

	// Endpoints are the endpoints of the API
	Endpoints []Endpoint `json:"endpoints"`

	// Version is the version of the API
	Version string `json:"version"`

	// Author is the author of the API
	Author string `json:"author"`

	// Date is the date of the API
	Date string `json:"date"`

	// License is the license of the API (text or URL)
	License []License `json:"license"`

	// TermsOfService is the terms of service of the API (text or URL)
	TermsOfService string `json:"termsOfService"`

	// Contact is the contact of the API
	Contact Contact `json:"contact"`
}
