// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SchemaMigrations is the golang structure of table schema_migrations for DAO operations like Where/Data.
type SchemaMigrations struct {
	g.Meta  `orm:"table:schema_migrations, do:true"`
	Version interface{} //
	Dirty   interface{} //
}
