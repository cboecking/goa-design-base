package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("MOEBOE_PROP_GORMA_STORAGE_GROUP", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("MOEBOE_PROP_GORMA_STORAGE_MODEL", func() {
			RendersTo(MOEBOE_PROP_GOA_MEDIA_TYPE_NAME)
			Description("MOEBOE_PROP_GOA_API_NAME MOEBOE_PROP_GORMA_STORAGE_MODEL")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("created_by", gorma.String)
		})

	})
})
