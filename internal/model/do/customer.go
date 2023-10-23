// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Customer is the golang structure of table customer for DAO operations like Where/Data.
type Customer struct {
	g.Meta            `orm:"table:customer, do:true"`
	Id                interface{} //
	CustNumbering     interface{} //
	MembershipCardNum interface{} //
	PhoneNumber       interface{} //
	Name              interface{} //
}
