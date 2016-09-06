package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the XAPI_NAME application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("XAPI_NAME", func() {
	Title("XAPI_TITLE")
	Description("XAPI_DESC")
	Contact(func() {
		Name("moeboe team")
		Email("chuck@moeboe.io")
		URL("http://moeboe.io")
	})
	Host("localhost:8081")
	Scheme("http")
	BasePath("/XAPI_NAME")

})
