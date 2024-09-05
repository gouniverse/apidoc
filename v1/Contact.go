package v1

type Contact struct {
	// Name is the identificating name of the contact person/organization.
	Name string `json:"name"`

	// Email is the email of the contact person/organization.
	Email string `json:"email"`

	// URL is the URL pointing to contact information of the person/organization.
	URL string `json:"url"`
}
