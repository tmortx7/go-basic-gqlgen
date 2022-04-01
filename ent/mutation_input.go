// Code generated by entc, DO NOT EDIT.

package ent

import (
	"go-basic-gqlgen/ent/schema/ulid"
	"time"
)

// CreateEmployeeInput represents a mutation input for creating employees.
type CreateEmployeeInput struct {
	FirstName string
	LastName  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// Mutate applies the CreateEmployeeInput on the EmployeeCreate builder.
func (i *CreateEmployeeInput) Mutate(m *EmployeeCreate) {
	m.SetFirstName(i.FirstName)
	m.SetLastName(i.LastName)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateEmployeeInput on the create builder.
func (c *EmployeeCreate) SetInput(i CreateEmployeeInput) *EmployeeCreate {
	i.Mutate(c)
	return c
}

// UpdateEmployeeInput represents a mutation input for updating employees.
type UpdateEmployeeInput struct {
	ID        ulid.ID
	FirstName *string
	LastName  *string
	UpdatedAt *time.Time
}

// Mutate applies the UpdateEmployeeInput on the EmployeeMutation.
func (i *UpdateEmployeeInput) Mutate(m *EmployeeMutation) {
	if v := i.FirstName; v != nil {
		m.SetFirstName(*v)
	}
	if v := i.LastName; v != nil {
		m.SetLastName(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the UpdateEmployeeInput on the update builder.
func (u *EmployeeUpdate) SetInput(i UpdateEmployeeInput) *EmployeeUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateEmployeeInput on the update-one builder.
func (u *EmployeeUpdateOne) SetInput(i UpdateEmployeeInput) *EmployeeUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateGroupInput represents a mutation input for creating groups.
type CreateGroupInput struct {
	Name    string
	UserIDs []ulid.ID
}

// Mutate applies the CreateGroupInput on the GroupCreate builder.
func (i *CreateGroupInput) Mutate(m *GroupCreate) {
	m.SetName(i.Name)
	if ids := i.UserIDs; len(ids) > 0 {
		m.AddUserIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateGroupInput on the create builder.
func (c *GroupCreate) SetInput(i CreateGroupInput) *GroupCreate {
	i.Mutate(c)
	return c
}

// UpdateGroupInput represents a mutation input for updating groups.
type UpdateGroupInput struct {
	ID            ulid.ID
	Name          *string
	AddUserIDs    []ulid.ID
	RemoveUserIDs []ulid.ID
}

// Mutate applies the UpdateGroupInput on the GroupMutation.
func (i *UpdateGroupInput) Mutate(m *GroupMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if ids := i.AddUserIDs; len(ids) > 0 {
		m.AddUserIDs(ids...)
	}
	if ids := i.RemoveUserIDs; len(ids) > 0 {
		m.RemoveUserIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateGroupInput on the update builder.
func (u *GroupUpdate) SetInput(i UpdateGroupInput) *GroupUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateGroupInput on the update-one builder.
func (u *GroupUpdateOne) SetInput(i UpdateGroupInput) *GroupUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateLinkInput represents a mutation input for creating links.
type CreateLinkInput struct {
	Title     string
	Address   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	UserID    *ulid.ID
}

// Mutate applies the CreateLinkInput on the LinkCreate builder.
func (i *CreateLinkInput) Mutate(m *LinkCreate) {
	m.SetTitle(i.Title)
	m.SetAddress(i.Address)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateLinkInput on the create builder.
func (c *LinkCreate) SetInput(i CreateLinkInput) *LinkCreate {
	i.Mutate(c)
	return c
}

// UpdateLinkInput represents a mutation input for updating links.
type UpdateLinkInput struct {
	ID        ulid.ID
	Title     *string
	Address   *string
	UpdatedAt *time.Time
	UserID    *ulid.ID
	ClearUser bool
}

// Mutate applies the UpdateLinkInput on the LinkMutation.
func (i *UpdateLinkInput) Mutate(m *LinkMutation) {
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Address; v != nil {
		m.SetAddress(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateLinkInput on the update builder.
func (u *LinkUpdate) SetInput(i UpdateLinkInput) *LinkUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateLinkInput on the update-one builder.
func (u *LinkUpdateOne) SetInput(i UpdateLinkInput) *LinkUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	LinkIDs   []ulid.ID
	GroupIDs  []ulid.ID
}

// Mutate applies the CreateUserInput on the UserCreate builder.
func (i *CreateUserInput) Mutate(m *UserCreate) {
	m.SetName(i.Name)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if ids := i.LinkIDs; len(ids) > 0 {
		m.AddLinkIDs(ids...)
	}
	if ids := i.GroupIDs; len(ids) > 0 {
		m.AddGroupIDs(ids...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the create builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c)
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	ID             ulid.ID
	Name           *string
	UpdatedAt      *time.Time
	AddLinkIDs     []ulid.ID
	RemoveLinkIDs  []ulid.ID
	AddGroupIDs    []ulid.ID
	RemoveGroupIDs []ulid.ID
}

// Mutate applies the UpdateUserInput on the UserMutation.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if ids := i.AddLinkIDs; len(ids) > 0 {
		m.AddLinkIDs(ids...)
	}
	if ids := i.RemoveLinkIDs; len(ids) > 0 {
		m.RemoveLinkIDs(ids...)
	}
	if ids := i.AddGroupIDs; len(ids) > 0 {
		m.AddGroupIDs(ids...)
	}
	if ids := i.RemoveGroupIDs; len(ids) > 0 {
		m.RemoveGroupIDs(ids...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the update builder.
func (u *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateUserInput on the update-one builder.
func (u *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
