package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("MOEBOE_PROP_GOA_RESOURCE_NAME", func() {

	DefaultMedia(MOEBOE_PROP_GOA_MEDIA_TYPE_NAME)
	BasePath("/MOEBOE_PROP_GOA_RESOURCE_NAMEs")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all MOEBOE_PROP_GOA_RESOURCE_NAMEs.")
		Response(OK, CollectionOf(MOEBOE_PROP_GOA_MEDIA_TYPE_NAME))
	})

	Action("show", func() {
		Routing(
			GET("/:MOEBOE_PROP_GOA_RESOURCE_NAMEID"),
		)
		Description("Retrieve MOEBOE_PROP_GOA_RESOURCE_NAME with given id.")
		Params(func() {
			Param("MOEBOE_PROP_GOA_RESOURCE_NAMEID", Integer, "MOEBOE_PROP_GOA_RESOURCE_NAME ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new MOEBOE_PROP_GOA_RESOURCE_NAME")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/MOEBOE_PROP_GOA_RESOURCE_NAMEs/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:MOEBOE_PROP_GOA_RESOURCE_NAMEID"),
		)
		Description("Change MOEBOE_PROP_GOA_RESOURCE_NAME name")
		Params(func() {
			Param("MOEBOE_PROP_GOA_RESOURCE_NAMEID", Integer, "MOEBOE_PROP_GOA_RESOURCE_NAME ID")
		})
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:MOEBOE_PROP_GOA_RESOURCE_NAMEID"),
		)
		Params(func() {
			Param("MOEBOE_PROP_GOA_RESOURCE_NAMEID", Integer, "MOEBOE_PROP_GOA_RESOURCE_NAME ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
