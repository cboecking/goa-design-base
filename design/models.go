package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("XSTORGROUP_U", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("XSTORMODEL_U", func() {
			RendersTo(XSTORMODEL_U)
			Description("XSTORGROUP_U XSTORMODEL_U")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
		})

	})
})
