package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the XAPI_NAME application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("XAPI_NAME", func() {
	Title("XAPI_NAME title")
	Description("XAPI_NAME description")
	Contact(func() {
		Name("moeboe team")
		Email("chuck@moeboe.io")
		URL("http://moeboe.io")
	})
	Host("localhost:XAPI_PORT")
	Scheme("http")
	BasePath("/XAPI_NAME")

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
