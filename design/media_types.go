package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// XMEDIATYPE_U is the XMEDIATYPE_L resource media type.
var XMEDIATYPE_U = MediaType("application/vnd.XMEDIATYPE_L+json", func() {
	Description("A XMEDIATYPE_L")
	Attributes(func() {
		Attribute("id", Integer, "ID of XMEDIATYPE_L", func() {
			Example(1)
		})
		Attribute("href", String, "API href of XMEDIATYPE_L", func() {
			Example("/XMEDIATYPE_Ls/1")
		})
		Attribute("name", String, "Name of XMEDIATYPE_U", func() {
			Example("test")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("created_by", String, "Email of XMEDIATYPE_L owner", func() {
			Format("email")
			Example("me@moeboe.io")
		})

		Required("id", "href", "name", "created_at", "created_by")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("created_at")
		Attribute("created_by")
	})

	View("tiny", func() {
		Description("tiny is the view used to list accounts")
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
