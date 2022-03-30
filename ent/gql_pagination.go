// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"go-basic-gqlgen/ent/employee"
	"go-basic-gqlgen/ent/schema/ulid"
	"io"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    ulid.ID `msgpack:"i"`
	Value Value   `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func getCollectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	field := fc.Field

walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Name == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return getCollectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

// EmployeeEdge is the edge representation of Employee.
type EmployeeEdge struct {
	Node   *Employee `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// EmployeeConnection is the connection containing edges to Employee.
type EmployeeConnection struct {
	Edges      []*EmployeeEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

// EmployeePaginateOption enables pagination customization.
type EmployeePaginateOption func(*employeePager) error

// WithEmployeeOrder configures pagination ordering.
func WithEmployeeOrder(order *EmployeeOrder) EmployeePaginateOption {
	if order == nil {
		order = DefaultEmployeeOrder
	}
	o := *order
	return func(pager *employeePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultEmployeeOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithEmployeeFilter configures pagination filter.
func WithEmployeeFilter(filter func(*EmployeeQuery) (*EmployeeQuery, error)) EmployeePaginateOption {
	return func(pager *employeePager) error {
		if filter == nil {
			return errors.New("EmployeeQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type employeePager struct {
	order  *EmployeeOrder
	filter func(*EmployeeQuery) (*EmployeeQuery, error)
}

func newEmployeePager(opts []EmployeePaginateOption) (*employeePager, error) {
	pager := &employeePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultEmployeeOrder
	}
	return pager, nil
}

func (p *employeePager) applyFilter(query *EmployeeQuery) (*EmployeeQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *employeePager) toCursor(e *Employee) Cursor {
	return p.order.Field.toCursor(e)
}

func (p *employeePager) applyCursors(query *EmployeeQuery, after, before *Cursor) *EmployeeQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultEmployeeOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *employeePager) applyOrder(query *EmployeeQuery, reverse bool) *EmployeeQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultEmployeeOrder.Field {
		query = query.Order(direction.orderFunc(DefaultEmployeeOrder.Field.field))
	}
	return query
}

// Paginate executes the query and returns a relay based cursor connection to Employee.
func (e *EmployeeQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...EmployeePaginateOption,
) (*EmployeeConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newEmployeePager(opts)
	if err != nil {
		return nil, err
	}

	if e, err = pager.applyFilter(e); err != nil {
		return nil, err
	}

	conn := &EmployeeConnection{Edges: []*EmployeeEdge{}}
	if !hasCollectedField(ctx, edgesField) || first != nil && *first == 0 || last != nil && *last == 0 {
		if hasCollectedField(ctx, totalCountField) ||
			hasCollectedField(ctx, pageInfoField) {
			count, err := e.Count(ctx)
			if err != nil {
				return nil, err
			}
			conn.TotalCount = count
			conn.PageInfo.HasNextPage = first != nil && count > 0
			conn.PageInfo.HasPreviousPage = last != nil && count > 0
		}
		return conn, nil
	}

	if (after != nil || first != nil || before != nil || last != nil) && hasCollectedField(ctx, totalCountField) {
		count, err := e.Clone().Count(ctx)
		if err != nil {
			return nil, err
		}
		conn.TotalCount = count
	}

	e = pager.applyCursors(e, after, before)
	e = pager.applyOrder(e, last != nil)
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	if limit > 0 {
		e = e.Limit(limit)
	}

	if field := getCollectedField(ctx, edgesField, nodeField); field != nil {
		e = e.collectField(graphql.GetOperationContext(ctx), *field)
	}

	nodes, err := e.All(ctx)
	if err != nil || len(nodes) == 0 {
		return conn, err
	}

	if len(nodes) == limit {
		conn.PageInfo.HasNextPage = first != nil
		conn.PageInfo.HasPreviousPage = last != nil
		nodes = nodes[:len(nodes)-1]
	}

	var nodeAt func(int) *Employee
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Employee {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Employee {
			return nodes[i]
		}
	}

	conn.Edges = make([]*EmployeeEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		conn.Edges[i] = &EmployeeEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor
	if conn.TotalCount == 0 {
		conn.TotalCount = len(nodes)
	}

	return conn, nil
}

// EmployeeOrderField defines the ordering field of Employee.
type EmployeeOrderField struct {
	field    string
	toCursor func(*Employee) Cursor
}

// EmployeeOrder defines the ordering of Employee.
type EmployeeOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *EmployeeOrderField `json:"field"`
}

// DefaultEmployeeOrder is the default ordering of Employee.
var DefaultEmployeeOrder = &EmployeeOrder{
	Direction: OrderDirectionAsc,
	Field: &EmployeeOrderField{
		field: employee.FieldID,
		toCursor: func(e *Employee) Cursor {
			return Cursor{ID: e.ID}
		},
	},
}

// ToEdge converts Employee into EmployeeEdge.
func (e *Employee) ToEdge(order *EmployeeOrder) *EmployeeEdge {
	if order == nil {
		order = DefaultEmployeeOrder
	}
	return &EmployeeEdge{
		Node:   e,
		Cursor: order.Field.toCursor(e),
	}
}
