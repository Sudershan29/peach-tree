// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmailAddress holds the string denoting the email_address field in the database.
	FieldEmailAddress = "email_address"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldPremium holds the string denoting the premium field in the database.
	FieldPremium = "premium"
	// EdgeSkills holds the string denoting the skills edge name in mutations.
	EdgeSkills = "skills"
	// EdgePreference holds the string denoting the preference edge name in mutations.
	EdgePreference = "preference"
	// Table holds the table name of the user in the database.
	Table = "users"
	// SkillsTable is the table that holds the skills relation/edge.
	SkillsTable = "user_skills"
	// SkillsInverseTable is the table name for the UserSkill entity.
	// It exists in this package in order to avoid circular dependency with the "userskill" package.
	SkillsInverseTable = "user_skills"
	// SkillsColumn is the table column denoting the skills relation/edge.
	SkillsColumn = "user_skills"
	// PreferenceTable is the table that holds the preference relation/edge.
	PreferenceTable = "preferences"
	// PreferenceInverseTable is the table name for the Preference entity.
	// It exists in this package in order to avoid circular dependency with the "preference" package.
	PreferenceInverseTable = "preferences"
	// PreferenceColumn is the table column denoting the preference relation/edge.
	PreferenceColumn = "user_preference"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmailAddress,
	FieldCreatedAt,
	FieldUUID,
	FieldPremium,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultPremium holds the default value on creation for the "premium" field.
	DefaultPremium bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmailAddress orders the results by the email_address field.
func ByEmailAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailAddress, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByPremium orders the results by the premium field.
func ByPremium(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPremium, opts...).ToFunc()
}

// BySkillsCount orders the results by skills count.
func BySkillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSkillsStep(), opts...)
	}
}

// BySkills orders the results by skills terms.
func BySkills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSkillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPreferenceField orders the results by preference field.
func ByPreferenceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPreferenceStep(), sql.OrderByField(field, opts...))
	}
}
func newSkillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SkillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SkillsTable, SkillsColumn),
	)
}
func newPreferenceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PreferenceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, PreferenceTable, PreferenceColumn),
	)
}
