package main

import (
	"log"

	"github.com/gouniverse/apidoc"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
	"github.com/gouniverse/webserver"
)

// main is the entry point of the application
// it will start the web server and serve the documentation
func main() {
	utils.EnvInitialize()

	host := utils.Env("SERVER_HOST")
	port := utils.Env("SERVER_PORT")

	if host == "" || port == "" {
		host = "127.0.0.1"
		port = "32080"
	}

	address := host + ":" + port

	log.Println(`Starting server on ` + address)

	server := webserver.New(address, apidoc.NewWebController(doc()).IdiomaticHandler)
	err := server.Start()

	if err != nil {
		panic(err)
	}
}

// doc returns the documentation to be served
//
// it sets the properties of the documentation
// - title
// - homeURL
// - handlerURL
// - description
// - logo
// - theme
// - themesDark
// - themesLight
// - info sections
// - custom group with sections
// - endpoints
func doc() apidoc.Documentation {
	return *apidoc.NewDocumentation().
		SetTitle("Example API Documentation").
		SetHomeURL("https://lesichkov.co.uk/api-doc").
		SetDescription("This is an example API documentation. It is not intended to be used in production.").
		SetLogo(hb.NewI().Class("bi bi-patch-check-fill").ToHTML()).
		SetLogo(`https://upload.wikimedia.org/wikipedia/commons/a/a6/Api_logo.jpg`).
		SetHandlerURL("/docs").
		SetTheme(apidoc.THEME_DARKLY).
		SetThemesDark([]struct {
			Key  string
			Name string
		}{
			{Key: apidoc.THEME_DARKLY, Name: "Dark"},
			{Key: apidoc.THEME_VAPOR, Name: "Dark 2"},
			{Key: apidoc.THEME_SOLAR, Name: "Dark 3"},
			{Key: apidoc.THEME_SLATE, Name: "Dark 4"},
			{Key: apidoc.THEME_CYBORG, Name: "Dark 5"},
			{Key: apidoc.THEME_SUPERHERO, Name: "Dark 6"},
		}).
		SetThemesLight([]struct {
			Key  string
			Name string
		}{
			{Key: apidoc.THEME_DEFAULT, Name: "Default"},
			{Key: apidoc.THEME_CERULEAN, Name: "Light 1"},
			{Key: apidoc.THEME_COSMO, Name: "Light 2"},
			{Key: apidoc.THEME_FLATLY, Name: "Light 3"},
			{Key: apidoc.THEME_JOURNAL, Name: "Light 4"},
			{Key: apidoc.THEME_LITERA, Name: "Light 5"},
			{Key: apidoc.THEME_LUMEN, Name: "Light 6"},
			{Key: apidoc.THEME_LUX, Name: "Light 7"},
			{Key: apidoc.THEME_MATERIA, Name: "Light 8"},
			{Key: apidoc.THEME_MINTY, Name: "Light 9"},
			{Key: apidoc.THEME_MORPH, Name: "Light 10"},
			{Key: apidoc.THEME_PULSE, Name: "Light 11"},
			{Key: apidoc.THEME_QUARTZ, Name: "Light 12"},
			{Key: apidoc.THEME_SANDSTONE, Name: "Light 13"},
			{Key: apidoc.THEME_SIMPLEX, Name: "Light 14"},
			{Key: apidoc.THEME_SKETCHY, Name: "Light 15"},
			{Key: apidoc.THEME_SPACELAB, Name: "Light 16"},
			{Key: apidoc.THEME_UNITED, Name: "Light 17"},
			{Key: apidoc.THEME_YETI, Name: "Light 18"},
			{Key: apidoc.THEME_ZEPHYR, Name: "Light 19"},
		}).
		AddBlock(sectionVersion()).
		AddBlock(sectionContact()).
		AddBlock(sectionLicense()).
		AddBlock(groupAuthentication()).
		AddBlock(groupEndpoints())
}

func sectionContact() apidoc.BlockInterface {
	return apidoc.NewSection().
		SetTitle("Contact").
		SetContent(`
This is an example contact section. If you have any technical questions or suggestions, please contact us at:

- [Github](https://github.com/gouniverse/apidoc)
- [Website](https://lesichkov.co.uk/contact)
		`)
}

func sectionLicense() apidoc.BlockInterface {
	return apidoc.NewSection().
		SetTitle("License").
		SetContent(`
This is an example license section. Please see the [license](https://github.com/gouniverse/apidoc/blob/master/LICENSE) for more information.
		`)
}

func sectionVersion() apidoc.BlockInterface {
	return apidoc.NewSection().
		SetTitle("Version").
		SetContent(`
This is an example version section. The current version of the API is 2.0.0.

The previous version of the API is 1.0.0. Documentation for the previous version can be found at:

- [https://lesichkov.co.uk/apidoc](https://lesichkov.co.uk/apidoc)
		`)
}

func groupAuthentication() apidoc.BlockInterface {
	return apidoc.NewGroup().
		SetTitle("Authentication").
		SetDescription(`This is an example authentication section`).
		AddBlock(apidoc.NewSection().
			SetTitle("API Key").
			SetContent(`
This is an example API key authentication section with code snippet.
The API key must be provided in the Authorization header of the request as follows:

~~~http	
Authorization: ApiKey <your_api_key>
~~~
			`))
}

func groupEndpoints() apidoc.BlockInterface {
	return apidoc.NewGroup().
		SetTitle("Endpoints").
		SetDescription(`This is an example endpoints section`).
		AddBlock(endpointCreateUser()).
		AddBlock(endpointGetUser()).
		AddBlock(endpointUpdateUser()).
		AddBlock(endpointDeleteUser()).
		AddBlock(endpointListUsers())
}

func endpointGetUser() apidoc.BlockInterface {
	return apidoc.NewEndpoint().
		SetTitle("Get User").
		SetMethod("GET").
		SetPath("/api/users/{id}").
		AddParameter(*apidoc.NewParameter().
			SetName("id").
			SetType("string").
			SetRequired(true).
			SetDescription(`The user ID`)).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Successful response").
			SetDescription(`The user was successfully retrieved`).
			SetStatusCode("200").
			SetExample(*apidoc.NewExample().
				SetTitle(`Successful response`).
				SetDescription(`The user was successfully retrieved. The message will contain more details about the user`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User retrieved successfully",
	"data": {
		"id": "123",
		"name": "John Doe"
	}
}
				`))).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Error response").
			SetDescription(`The user was not found`).
			SetStatusCode("404").
			SetExample(*apidoc.NewExample().
				SetTitle(`Error response`).
				SetDescription(`The user was not found. The message will contain more details about the error`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User not found"
}
				`)))
}

func endpointListUsers() apidoc.BlockInterface {
	return apidoc.NewEndpoint().
		SetTitle("List Users").
		SetMethod("GET").
		SetPath("/api/users").
		AddResponse(*apidoc.NewResponse().
			SetTitle("Successful response").
			SetDescription(`The users were successfully retrieved`).
			SetStatusCode("200").
			SetExample(*apidoc.NewExample().
				SetTitle(`Successful response`).
				SetDescription(`The users were successfully retrieved. The message will contain more details about the users`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "Users retrieved successfully",
	"data": [{
		"id": "123",
		"name": "John Doe"
	}, {
		"id": "456",
		"name": "Jane Doe"
	}]
}
				`))).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Error response").
			SetDescription(`The users were not found`).
			SetStatusCode("404").
			SetExample(*apidoc.NewExample().
				SetTitle(`Error response`).
				SetDescription(`The users were not found. The message will contain more details about the error`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "Users not found"
}
				`)))
}

func endpointCreateUser() apidoc.BlockInterface {
	return apidoc.NewEndpoint().
		SetTitle("Create User").
		SetMethod("POST").
		SetPath("/api/users").
		AddParameter(*apidoc.NewParameter().
			SetName("name").
			SetType("string").
			SetRequired(true).
			SetDescription(`The name of the user`)).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Successful response").
			SetDescription(`The user was successfully created`).
			SetStatusCode("201").
			SetExample(*apidoc.NewExample().
				SetTitle(`Successful response`).
				SetDescription(`The user was successfully created. The message will contain more details about the user`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User created successfully",
	"data": {
		"id": "123",
		"name": "John Doe"
	}
}
				`))).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Error response").
			SetDescription(`The user was not created`).
			SetStatusCode("400").
			SetExample(*apidoc.NewExample().
				SetTitle(`Error response`).
				SetDescription(`The user was not created. The message will contain more details about the error`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User not created"
}
				`)))
}

func endpointUpdateUser() apidoc.BlockInterface {
	return apidoc.NewEndpoint().
		SetTitle("Update User").
		SetMethod("PUT").
		SetPath("/api/users/{id}").
		AddParameter(*apidoc.NewParameter().
			SetName("id").
			SetType("string").
			SetRequired(true).
			SetDescription(`The user ID`)).
		AddParameter(*apidoc.NewParameter().
			SetName("name").
			SetType("string").
			SetRequired(true).
			SetDescription(`The name of the user`)).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Successful response").
			SetDescription(`The user was successfully updated`).
			SetStatusCode("200").
			SetExample(*apidoc.NewExample().
				SetTitle(`Successful response`).
				SetDescription(`The user was successfully updated. The message will contain more details about the user`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User updated successfully",
	"data": {
		"id": "123",
		"name": "John Doe"
	}
}
				`))).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Error response").
			SetDescription(`The user was not updated`).
			SetStatusCode("400").
			SetExample(*apidoc.NewExample().
				SetTitle(`Error response`).
				SetDescription(`The user was not updated. The message will contain more details about the error`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User not updated"
}
				`)))
}

func endpointDeleteUser() apidoc.BlockInterface {
	return apidoc.NewEndpoint().
		SetTitle("Delete User").
		SetMethod("DELETE").
		SetPath("/api/users/{id}").
		AddParameter(*apidoc.NewParameter().
			SetName("id").
			SetType("string").
			SetRequired(true).
			SetDescription(`The user ID`)).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Successful response").
			SetDescription(`The user was successfully deleted`).
			SetStatusCode("200").
			SetExample(*apidoc.NewExample().
				SetTitle(`Successful response`).
				SetDescription(`The user was successfully deleted. The message will contain more details about the user`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User deleted successfully",
	"data": {
		"id": "123",
		"name": "John Doe"
	}
}
				`))).
		AddResponse(*apidoc.NewResponse().
			SetTitle("Error response").
			SetDescription(`The user was not deleted`).
			SetStatusCode("400").
			SetExample(*apidoc.NewExample().
				SetTitle(`Error response`).
				SetDescription(`The user was not deleted. The message will contain more details about the error`).
				SetSourceLanguage(`json`).
				SetSourceCode(`
{
	"message": "User not deleted"
}
				`)))
}
