package globalid

import (
	"fmt"
	"go-basic-gqlgen/ent/employee"
	"go-basic-gqlgen/ent/group"
	"go-basic-gqlgen/ent/link"
	"go-basic-gqlgen/ent/todo"
	"go-basic-gqlgen/ent/user"
	"log"
	"reflect"
)

type field struct {
	Prefix string
	Table  string
}

// GlobalIDs maps unique string to table names
type GlobalIDs struct {
	Employee field
	User field
	Link field
	Group field
	Todo field
}

// New generates a map object that is intended to be used as global identification for node interface query.
// Prefix should maintain constrained to 3 characters for encoding the entity type.
func New() GlobalIDs {
	return GlobalIDs{
		Employee: field{
			Prefix: "emp_",
			Table:  employee.Table,
		},
		User: field{
			Prefix: "usr_",
			Table: user.Table,
		},
		Link: field{
			Prefix: "lnk_",
			Table: link.Table,
		},
		Group: field{
			Prefix: "grp_",
			Table: group.Table,
		},
		Todo: field{
			Prefix: "tod_",
			Table: todo.Table,
		},
	}
}

var globalIDs = New()
var maps = structToMap(&globalIDs)

// FindTableByID returns table name by passed id
func (GlobalIDs) FindTableByID(id string) (string, error) {
	v, ok := maps[id]
	if !ok {
		return "", fmt.Errorf("could not map '%s' to a table name", id)
	}
	return v, nil
}

func structToMap(data *GlobalIDs) map[string]string {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	result := make(map[string]string, size)

	for i := 0; i < size; i++ {
		value := elem.Field(i).Interface()
		f, ok := value.(field)
		if !ok {
			log.Fatalf("Cannot convert struct to map")
		}
		result[f.Prefix] = f.Table
	}

	return result
}
