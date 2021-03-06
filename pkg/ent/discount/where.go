// Code generated by entc, DO NOT EDIT.

package discount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PeriodStart applies equality check predicate on the "period_start" field. It's identical to PeriodStartEQ.
func PeriodStart(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriodStart), v))
	})
}

// PeriodEnd applies equality check predicate on the "period_end" field. It's identical to PeriodEndEQ.
func PeriodEnd(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriodEnd), v))
	})
}

// DiscountPrice applies equality check predicate on the "discount_price" field. It's identical to DiscountPriceEQ.
func DiscountPrice(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPercentage applies equality check predicate on the "discount_percentage" field. It's identical to DiscountPercentageEQ.
func DiscountPercentage(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountPercentage), v))
	})
}

// PeriodStartEQ applies the EQ predicate on the "period_start" field.
func PeriodStartEQ(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriodStart), v))
	})
}

// PeriodStartNEQ applies the NEQ predicate on the "period_start" field.
func PeriodStartNEQ(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPeriodStart), v))
	})
}

// PeriodStartIn applies the In predicate on the "period_start" field.
func PeriodStartIn(vs ...time.Time) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPeriodStart), v...))
	})
}

// PeriodStartNotIn applies the NotIn predicate on the "period_start" field.
func PeriodStartNotIn(vs ...time.Time) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPeriodStart), v...))
	})
}

// PeriodStartGT applies the GT predicate on the "period_start" field.
func PeriodStartGT(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPeriodStart), v))
	})
}

// PeriodStartGTE applies the GTE predicate on the "period_start" field.
func PeriodStartGTE(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPeriodStart), v))
	})
}

// PeriodStartLT applies the LT predicate on the "period_start" field.
func PeriodStartLT(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPeriodStart), v))
	})
}

// PeriodStartLTE applies the LTE predicate on the "period_start" field.
func PeriodStartLTE(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPeriodStart), v))
	})
}

// PeriodStartIsNil applies the IsNil predicate on the "period_start" field.
func PeriodStartIsNil() predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPeriodStart)))
	})
}

// PeriodStartNotNil applies the NotNil predicate on the "period_start" field.
func PeriodStartNotNil() predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPeriodStart)))
	})
}

// PeriodEndEQ applies the EQ predicate on the "period_end" field.
func PeriodEndEQ(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriodEnd), v))
	})
}

// PeriodEndNEQ applies the NEQ predicate on the "period_end" field.
func PeriodEndNEQ(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPeriodEnd), v))
	})
}

// PeriodEndIn applies the In predicate on the "period_end" field.
func PeriodEndIn(vs ...time.Time) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPeriodEnd), v...))
	})
}

// PeriodEndNotIn applies the NotIn predicate on the "period_end" field.
func PeriodEndNotIn(vs ...time.Time) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPeriodEnd), v...))
	})
}

// PeriodEndGT applies the GT predicate on the "period_end" field.
func PeriodEndGT(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPeriodEnd), v))
	})
}

// PeriodEndGTE applies the GTE predicate on the "period_end" field.
func PeriodEndGTE(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPeriodEnd), v))
	})
}

// PeriodEndLT applies the LT predicate on the "period_end" field.
func PeriodEndLT(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPeriodEnd), v))
	})
}

// PeriodEndLTE applies the LTE predicate on the "period_end" field.
func PeriodEndLTE(v time.Time) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPeriodEnd), v))
	})
}

// PeriodEndIsNil applies the IsNil predicate on the "period_end" field.
func PeriodEndIsNil() predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPeriodEnd)))
	})
}

// PeriodEndNotNil applies the NotNil predicate on the "period_end" field.
func PeriodEndNotNil() predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPeriodEnd)))
	})
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v Method) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMethod), v))
	})
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v Method) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMethod), v))
	})
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...Method) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMethod), v...))
	})
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...Method) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMethod), v...))
	})
}

// DiscountPriceEQ applies the EQ predicate on the "discount_price" field.
func DiscountPriceEQ(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPriceNEQ applies the NEQ predicate on the "discount_price" field.
func DiscountPriceNEQ(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPriceIn applies the In predicate on the "discount_price" field.
func DiscountPriceIn(vs ...int) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscountPrice), v...))
	})
}

// DiscountPriceNotIn applies the NotIn predicate on the "discount_price" field.
func DiscountPriceNotIn(vs ...int) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscountPrice), v...))
	})
}

// DiscountPriceGT applies the GT predicate on the "discount_price" field.
func DiscountPriceGT(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPriceGTE applies the GTE predicate on the "discount_price" field.
func DiscountPriceGTE(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPriceLT applies the LT predicate on the "discount_price" field.
func DiscountPriceLT(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPriceLTE applies the LTE predicate on the "discount_price" field.
func DiscountPriceLTE(v int) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscountPrice), v))
	})
}

// DiscountPercentageEQ applies the EQ predicate on the "discount_percentage" field.
func DiscountPercentageEQ(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountPercentage), v))
	})
}

// DiscountPercentageNEQ applies the NEQ predicate on the "discount_percentage" field.
func DiscountPercentageNEQ(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscountPercentage), v))
	})
}

// DiscountPercentageIn applies the In predicate on the "discount_percentage" field.
func DiscountPercentageIn(vs ...float64) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscountPercentage), v...))
	})
}

// DiscountPercentageNotIn applies the NotIn predicate on the "discount_percentage" field.
func DiscountPercentageNotIn(vs ...float64) predicate.Discount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discount(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscountPercentage), v...))
	})
}

// DiscountPercentageGT applies the GT predicate on the "discount_percentage" field.
func DiscountPercentageGT(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscountPercentage), v))
	})
}

// DiscountPercentageGTE applies the GTE predicate on the "discount_percentage" field.
func DiscountPercentageGTE(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscountPercentage), v))
	})
}

// DiscountPercentageLT applies the LT predicate on the "discount_percentage" field.
func DiscountPercentageLT(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscountPercentage), v))
	})
}

// DiscountPercentageLTE applies the LTE predicate on the "discount_percentage" field.
func DiscountPercentageLTE(v float64) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscountPercentage), v))
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrdersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrdersTable, OrdersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrdersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrdersTable, OrdersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Discount) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Discount) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Discount) predicate.Discount {
	return predicate.Discount(func(s *sql.Selector) {
		p(s.Not())
	})
}
