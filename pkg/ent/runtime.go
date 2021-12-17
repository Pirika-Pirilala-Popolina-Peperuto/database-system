// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/order"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderFields[2].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
}