package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// MOEBOE_PROP_GOA_MEDIA_TYPE_NAME is the MOEBOE_PROP_GOA_RESOURCE_NAME resource media type.
var MOEBOE_PROP_GOA_MEDIA_TYPE_NAME = MediaType("application/vnd.MOEBOE_PROP_GOA_RESOURCE_NAME+json", func() {
	Description("A MOEBOE_PROP_GOA_MEDIA_TYPE_NAME")
	Attributes(func() {
		Attribute("id", Integer, "ID of MOEBOE_PROP_GOA_RESOURCE_NAME", func() {
			Example(1)
		})
		Attribute("href", String, "API href of MOEBOE_PROP_GOA_RESOURCE_NAME", func() {
			Example("/MOEBOE_PROP_GOA_RESOURCE_NAME/1")
		})
		Attribute("name", String, "Name of MOEBOE_PROP_GOA_RESOURCE_NAME", func() {
			Example("test")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("created_by", String, "Email of MOEBOE_PROP_GOA_RESOURCE_NAME owner", func() {
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
