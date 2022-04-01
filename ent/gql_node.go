// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"go-basic-gqlgen/ent/employee"
	"go-basic-gqlgen/ent/group"
	"go-basic-gqlgen/ent/link"
	"go-basic-gqlgen/ent/schema/ulid"
	"go-basic-gqlgen/ent/user"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     ulid.ID  `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string    `json:"type,omitempty"` // edge type.
	Name string    `json:"name,omitempty"` // edge name.
	IDs  []ulid.ID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (e *Employee) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     e.ID,
		Type:   "Employee",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(e.FirstName); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "firstName",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.LastName); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "lastName",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(e.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	return node, nil
}

func (gr *Group) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "Group",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "users",
	}
	err = gr.QueryUsers().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (l *Link) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     l.ID,
		Type:   "Link",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(l.Title); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "title",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.Address); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "address",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(l.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = l.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Link",
		Name: "links",
	}
	err = u.QueryLinks().
		Select(link.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Group",
		Name: "groups",
	}
	err = u.QueryGroups().
		Select(group.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id ulid.ID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, ulid.ID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, ulid.ID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, ulid.ID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id ulid.ID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id ulid.ID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id ulid.ID) (Noder, error) {
	switch table {
	case employee.Table:
		n, err := c.Employee.Query().
			Where(employee.ID(id)).
			CollectFields(ctx, "Employee").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		n, err := c.Group.Query().
			Where(group.ID(id)).
			CollectFields(ctx, "Group").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case link.Table:
		n, err := c.Link.Query().
			Where(link.ID(id)).
			CollectFields(ctx, "Link").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		n, err := c.User.Query().
			Where(user.ID(id)).
			CollectFields(ctx, "User").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []ulid.ID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]ulid.ID)
	id2idx := make(map[ulid.ID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []ulid.ID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[ulid.ID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case employee.Table:
		nodes, err := c.Employee.Query().
			Where(employee.IDIn(ids...)).
			CollectFields(ctx, "Employee").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		nodes, err := c.Group.Query().
			Where(group.IDIn(ids...)).
			CollectFields(ctx, "Group").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case link.Table:
		nodes, err := c.Link.Query().
			Where(link.IDIn(ids...)).
			CollectFields(ctx, "Link").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		nodes, err := c.User.Query().
			Where(user.IDIn(ids...)).
			CollectFields(ctx, "User").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
