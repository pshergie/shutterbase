// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/shutterbase/shutterbase/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailValidated applies equality check predicate on the "email_validated" field. It's identical to EmailValidatedEQ.
func EmailValidated(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailValidated, v))
}

// ValidationKey applies equality check predicate on the "validation_key" field. It's identical to ValidationKeyEQ.
func ValidationKey(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldValidationKey, v))
}

// ValidationSentAt applies equality check predicate on the "validation_sent_at" field. It's identical to ValidationSentAtEQ.
func ValidationSentAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldValidationSentAt, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordResetKey applies equality check predicate on the "password_reset_key" field. It's identical to PasswordResetKeyEQ.
func PasswordResetKey(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordResetKey, v))
}

// PasswordResetAt applies equality check predicate on the "password_reset_at" field. It's identical to PasswordResetAtEQ.
func PasswordResetAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordResetAt, v))
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldActive, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLastName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// EmailValidatedEQ applies the EQ predicate on the "email_validated" field.
func EmailValidatedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailValidated, v))
}

// EmailValidatedNEQ applies the NEQ predicate on the "email_validated" field.
func EmailValidatedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmailValidated, v))
}

// ValidationKeyEQ applies the EQ predicate on the "validation_key" field.
func ValidationKeyEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldValidationKey, v))
}

// ValidationKeyNEQ applies the NEQ predicate on the "validation_key" field.
func ValidationKeyNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldValidationKey, v))
}

// ValidationKeyIn applies the In predicate on the "validation_key" field.
func ValidationKeyIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldValidationKey, vs...))
}

// ValidationKeyNotIn applies the NotIn predicate on the "validation_key" field.
func ValidationKeyNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldValidationKey, vs...))
}

// ValidationKeyGT applies the GT predicate on the "validation_key" field.
func ValidationKeyGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldValidationKey, v))
}

// ValidationKeyGTE applies the GTE predicate on the "validation_key" field.
func ValidationKeyGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldValidationKey, v))
}

// ValidationKeyLT applies the LT predicate on the "validation_key" field.
func ValidationKeyLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldValidationKey, v))
}

// ValidationKeyLTE applies the LTE predicate on the "validation_key" field.
func ValidationKeyLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldValidationKey, v))
}

// ValidationSentAtEQ applies the EQ predicate on the "validation_sent_at" field.
func ValidationSentAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldValidationSentAt, v))
}

// ValidationSentAtNEQ applies the NEQ predicate on the "validation_sent_at" field.
func ValidationSentAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldValidationSentAt, v))
}

// ValidationSentAtIn applies the In predicate on the "validation_sent_at" field.
func ValidationSentAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldValidationSentAt, vs...))
}

// ValidationSentAtNotIn applies the NotIn predicate on the "validation_sent_at" field.
func ValidationSentAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldValidationSentAt, vs...))
}

// ValidationSentAtGT applies the GT predicate on the "validation_sent_at" field.
func ValidationSentAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldValidationSentAt, v))
}

// ValidationSentAtGTE applies the GTE predicate on the "validation_sent_at" field.
func ValidationSentAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldValidationSentAt, v))
}

// ValidationSentAtLT applies the LT predicate on the "validation_sent_at" field.
func ValidationSentAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldValidationSentAt, v))
}

// ValidationSentAtLTE applies the LTE predicate on the "validation_sent_at" field.
func ValidationSentAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldValidationSentAt, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordResetKeyEQ applies the EQ predicate on the "password_reset_key" field.
func PasswordResetKeyEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordResetKey, v))
}

// PasswordResetKeyNEQ applies the NEQ predicate on the "password_reset_key" field.
func PasswordResetKeyNEQ(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswordResetKey, v))
}

// PasswordResetKeyIn applies the In predicate on the "password_reset_key" field.
func PasswordResetKeyIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswordResetKey, vs...))
}

// PasswordResetKeyNotIn applies the NotIn predicate on the "password_reset_key" field.
func PasswordResetKeyNotIn(vs ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswordResetKey, vs...))
}

// PasswordResetKeyGT applies the GT predicate on the "password_reset_key" field.
func PasswordResetKeyGT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswordResetKey, v))
}

// PasswordResetKeyGTE applies the GTE predicate on the "password_reset_key" field.
func PasswordResetKeyGTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswordResetKey, v))
}

// PasswordResetKeyLT applies the LT predicate on the "password_reset_key" field.
func PasswordResetKeyLT(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswordResetKey, v))
}

// PasswordResetKeyLTE applies the LTE predicate on the "password_reset_key" field.
func PasswordResetKeyLTE(v uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswordResetKey, v))
}

// PasswordResetAtEQ applies the EQ predicate on the "password_reset_at" field.
func PasswordResetAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordResetAt, v))
}

// PasswordResetAtNEQ applies the NEQ predicate on the "password_reset_at" field.
func PasswordResetAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswordResetAt, v))
}

// PasswordResetAtIn applies the In predicate on the "password_reset_at" field.
func PasswordResetAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswordResetAt, vs...))
}

// PasswordResetAtNotIn applies the NotIn predicate on the "password_reset_at" field.
func PasswordResetAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswordResetAt, vs...))
}

// PasswordResetAtGT applies the GT predicate on the "password_reset_at" field.
func PasswordResetAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswordResetAt, v))
}

// PasswordResetAtGTE applies the GTE predicate on the "password_reset_at" field.
func PasswordResetAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswordResetAt, v))
}

// PasswordResetAtLT applies the LT predicate on the "password_reset_at" field.
func PasswordResetAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswordResetAt, v))
}

// PasswordResetAtLTE applies the LTE predicate on the "password_reset_at" field.
func PasswordResetAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswordResetAt, v))
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldActive, v))
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldActive, v))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectAssignments applies the HasEdge predicate on the "projectAssignments" edge.
func HasProjectAssignments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ProjectAssignmentsTable, ProjectAssignmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectAssignmentsWith applies the HasEdge predicate on the "projectAssignments" edge with a given conditions (other predicates).
func HasProjectAssignmentsWith(preds ...predicate.ProjectAssignment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newProjectAssignmentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Image) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedUsers applies the HasEdge predicate on the "created_users" edge.
func HasCreatedUsers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CreatedUsersTable, CreatedUsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedUsersWith applies the HasEdge predicate on the "created_users" edge with a given conditions (other predicates).
func HasCreatedUsersWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCreatedUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedBy applies the HasEdge predicate on the "created_by" edge.
func HasCreatedBy() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "created_by" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCreatedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasModifiedUsers applies the HasEdge predicate on the "modified_users" edge.
func HasModifiedUsers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ModifiedUsersTable, ModifiedUsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModifiedUsersWith applies the HasEdge predicate on the "modified_users" edge with a given conditions (other predicates).
func HasModifiedUsersWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newModifiedUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasModifiedBy applies the HasEdge predicate on the "modified_by" edge.
func HasModifiedBy() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ModifiedByTable, ModifiedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModifiedByWith applies the HasEdge predicate on the "modified_by" edge with a given conditions (other predicates).
func HasModifiedByWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newModifiedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
