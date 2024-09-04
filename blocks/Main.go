package blocks

func Main() *Documentation {
	groupInfo := NewGroupBlock().
		SetTitle(`Introduction`).
		SetContent(`The API documentation is generated from the apidoc block generator.`)

	groupEndpoints := NewGroupBlock().
		SetTitle("Endpoints").
		AddBlock(NewEndpointBlock().
			SetTitle("Token Create").
			SetContent("Creates a secure token with a password. The token will be returned in the response. The token and its corresponding password are required to delete the token.").
			SetMethod("POST").
			SetPath("/api/token-create").
			AddParameter(*NewParameter().
				SetName("value").
				SetTitle("The value to be encrypted").
				SetContent("The value to be encrypted. This value will be encrypted with the password.").
				SetLocation(LOCATION_QUERY).
				SetRequired(true)).
			AddParameter(*NewParameter().
				SetName("password").
				SetTitle("The password used to encrypt the value").
				SetContent("The password used to encrypt the value when the secure token was created. This password is required to delete the token.").
				SetLocation(LOCATION_QUERY).
				SetRequired(true)),
		)

	doc := NewDocumentation().
		AddBlock(groupInfo).
		AddBlock(groupEndpoints)

	return doc
}
