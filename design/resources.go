package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("XRESOURCE_L", func() {

	DefaultMedia(XRESOURCE_U)
	BasePath("/XRESOURCE_Ls")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all XRESOURCE_Ls.")
		Response(OK, CollectionOf(XRESOURCE_U))
	})

	Action("show", func() {
		Routing(
			GET("/:XRESOURCE_LID"),
		)
		Description("Retrieve XRESOURCE_L with given id.")
		Params(func() {
			Param("XRESOURCE_LID", Integer, "XRESOURCE_U ID", func() {
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
		Description("Create new XRESOURCE_L")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/XRESOURCE_Ls/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:XRESOURCE_LID"),
		)
		Description("Change XRESOURCE_L name")
		Params(func() {
			Param("XRESOURCE_LID", Integer, "XRESOURCE_U ID")
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
			DELETE("/:XRESOURCE_LID"),
		)
		Params(func() {
			Param("XRESOURCE_LID", Integer, "XRESOURCE_U ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
