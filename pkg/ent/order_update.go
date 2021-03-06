// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/discount"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/order"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/predicate"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/product"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/user"
	"github.com/google/uuid"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCustomerID sets the "customer_id" field.
func (ou *OrderUpdate) SetCustomerID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetCustomerID(u)
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetCreatedAt(t)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetCreatedAt(*t)
	}
	return ou
}

// SetProcessStatus sets the "process_status" field.
func (ou *OrderUpdate) SetProcessStatus(os order.ProcessStatus) *OrderUpdate {
	ou.mutation.SetProcessStatus(os)
	return ou
}

// SetRemark sets the "remark" field.
func (ou *OrderUpdate) SetRemark(s string) *OrderUpdate {
	ou.mutation.SetRemark(s)
	return ou
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableRemark(s *string) *OrderUpdate {
	if s != nil {
		ou.SetRemark(*s)
	}
	return ou
}

// ClearRemark clears the value of the "remark" field.
func (ou *OrderUpdate) ClearRemark() *OrderUpdate {
	ou.mutation.ClearRemark()
	return ou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *OrderUpdate) SetUserID(id uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *OrderUpdate) SetUser(u *User) *OrderUpdate {
	return ou.SetUserID(u.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ou *OrderUpdate) AddProductIDs(ids ...uuid.UUID) *OrderUpdate {
	ou.mutation.AddProductIDs(ids...)
	return ou
}

// AddProducts adds the "products" edges to the Product entity.
func (ou *OrderUpdate) AddProducts(p ...*Product) *OrderUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddProductIDs(ids...)
}

// AddDiscountIDs adds the "discounts" edge to the Discount entity by IDs.
func (ou *OrderUpdate) AddDiscountIDs(ids ...uuid.UUID) *OrderUpdate {
	ou.mutation.AddDiscountIDs(ids...)
	return ou
}

// AddDiscounts adds the "discounts" edges to the Discount entity.
func (ou *OrderUpdate) AddDiscounts(d ...*Discount) *OrderUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ou.AddDiscountIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ou *OrderUpdate) ClearUser() *OrderUpdate {
	ou.mutation.ClearUser()
	return ou
}

// ClearProducts clears all "products" edges to the Product entity.
func (ou *OrderUpdate) ClearProducts() *OrderUpdate {
	ou.mutation.ClearProducts()
	return ou
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (ou *OrderUpdate) RemoveProductIDs(ids ...uuid.UUID) *OrderUpdate {
	ou.mutation.RemoveProductIDs(ids...)
	return ou
}

// RemoveProducts removes "products" edges to Product entities.
func (ou *OrderUpdate) RemoveProducts(p ...*Product) *OrderUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemoveProductIDs(ids...)
}

// ClearDiscounts clears all "discounts" edges to the Discount entity.
func (ou *OrderUpdate) ClearDiscounts() *OrderUpdate {
	ou.mutation.ClearDiscounts()
	return ou
}

// RemoveDiscountIDs removes the "discounts" edge to Discount entities by IDs.
func (ou *OrderUpdate) RemoveDiscountIDs(ids ...uuid.UUID) *OrderUpdate {
	ou.mutation.RemoveDiscountIDs(ids...)
	return ou
}

// RemoveDiscounts removes "discounts" edges to Discount entities.
func (ou *OrderUpdate) RemoveDiscounts(d ...*Discount) *OrderUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ou.RemoveDiscountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.ProcessStatus(); ok {
		if err := order.ProcessStatusValidator(v); err != nil {
			return &ValidationError{Name: "process_status", err: fmt.Errorf("ent: validator failed for field \"process_status\": %w", err)}
		}
	}
	if _, ok := ou.mutation.UserID(); ou.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.ProcessStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: order.FieldProcessStatus,
		})
	}
	if value, ok := ou.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldRemark,
		})
	}
	if ou.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldRemark,
		})
	}
	if ou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: order.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedProductsIDs(); len(nodes) > 0 && !ou.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: order.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: order.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.DiscountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.DiscountsTable,
			Columns: order.DiscountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedDiscountsIDs(); len(nodes) > 0 && !ou.mutation.DiscountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.DiscountsTable,
			Columns: order.DiscountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.DiscountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.DiscountsTable,
			Columns: order.DiscountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetCustomerID sets the "customer_id" field.
func (ouo *OrderUpdateOne) SetCustomerID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetCustomerID(u)
	return ouo
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetCreatedAt(t)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetCreatedAt(*t)
	}
	return ouo
}

// SetProcessStatus sets the "process_status" field.
func (ouo *OrderUpdateOne) SetProcessStatus(os order.ProcessStatus) *OrderUpdateOne {
	ouo.mutation.SetProcessStatus(os)
	return ouo
}

// SetRemark sets the "remark" field.
func (ouo *OrderUpdateOne) SetRemark(s string) *OrderUpdateOne {
	ouo.mutation.SetRemark(s)
	return ouo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableRemark(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetRemark(*s)
	}
	return ouo
}

// ClearRemark clears the value of the "remark" field.
func (ouo *OrderUpdateOne) ClearRemark() *OrderUpdateOne {
	ouo.mutation.ClearRemark()
	return ouo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *OrderUpdateOne) SetUserID(id uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *OrderUpdateOne) SetUser(u *User) *OrderUpdateOne {
	return ouo.SetUserID(u.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ouo *OrderUpdateOne) AddProductIDs(ids ...uuid.UUID) *OrderUpdateOne {
	ouo.mutation.AddProductIDs(ids...)
	return ouo
}

// AddProducts adds the "products" edges to the Product entity.
func (ouo *OrderUpdateOne) AddProducts(p ...*Product) *OrderUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddProductIDs(ids...)
}

// AddDiscountIDs adds the "discounts" edge to the Discount entity by IDs.
func (ouo *OrderUpdateOne) AddDiscountIDs(ids ...uuid.UUID) *OrderUpdateOne {
	ouo.mutation.AddDiscountIDs(ids...)
	return ouo
}

// AddDiscounts adds the "discounts" edges to the Discount entity.
func (ouo *OrderUpdateOne) AddDiscounts(d ...*Discount) *OrderUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ouo.AddDiscountIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *OrderUpdateOne) ClearUser() *OrderUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// ClearProducts clears all "products" edges to the Product entity.
func (ouo *OrderUpdateOne) ClearProducts() *OrderUpdateOne {
	ouo.mutation.ClearProducts()
	return ouo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (ouo *OrderUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *OrderUpdateOne {
	ouo.mutation.RemoveProductIDs(ids...)
	return ouo
}

// RemoveProducts removes "products" edges to Product entities.
func (ouo *OrderUpdateOne) RemoveProducts(p ...*Product) *OrderUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemoveProductIDs(ids...)
}

// ClearDiscounts clears all "discounts" edges to the Discount entity.
func (ouo *OrderUpdateOne) ClearDiscounts() *OrderUpdateOne {
	ouo.mutation.ClearDiscounts()
	return ouo
}

// RemoveDiscountIDs removes the "discounts" edge to Discount entities by IDs.
func (ouo *OrderUpdateOne) RemoveDiscountIDs(ids ...uuid.UUID) *OrderUpdateOne {
	ouo.mutation.RemoveDiscountIDs(ids...)
	return ouo
}

// RemoveDiscounts removes "discounts" edges to Discount entities.
func (ouo *OrderUpdateOne) RemoveDiscounts(d ...*Discount) *OrderUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ouo.RemoveDiscountIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.ProcessStatus(); ok {
		if err := order.ProcessStatusValidator(v); err != nil {
			return &ValidationError{Name: "process_status", err: fmt.Errorf("ent: validator failed for field \"process_status\": %w", err)}
		}
	}
	if _, ok := ouo.mutation.UserID(); ouo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Order.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.ProcessStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: order.FieldProcessStatus,
		})
	}
	if value, ok := ouo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldRemark,
		})
	}
	if ouo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldRemark,
		})
	}
	if ouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: order.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !ouo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: order.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.ProductsTable,
			Columns: order.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.DiscountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.DiscountsTable,
			Columns: order.DiscountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discount.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedDiscountsIDs(); len(nodes) > 0 && !ouo.mutation.DiscountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.DiscountsTable,
			Columns: order.DiscountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.DiscountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.DiscountsTable,
			Columns: order.DiscountsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: discount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
