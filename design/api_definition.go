package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the MOEBOE_PROP_GOA_API_NAME application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("MOEBOE_PROP_GOA_API_NAME", func() {
	Title("MOEBOE_PROP_GOA_API_NAME title")
	Description("MOEBOE_PROP_GOA_API_NAME description")
	Contact(func() {
		Name("moeboe team")
		Email("chuck@moeboe.io")
		URL("http://moeboe.io")
	})
	Host("localhost:MOEBOE_PROP_GOA_API_PORT")
	Scheme("http")
	BasePath("/MOEBOE_PROP_GOA_API_NAME")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
