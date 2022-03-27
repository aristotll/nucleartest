// Code generated by entc, DO NOT EDIT.

package student

import (
	"ent/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CarID applies equality check predicate on the "car_id" field. It's identical to CarIDEQ.
func CarID(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCarID), v))
	})
}

// CID applies equality check predicate on the "c_id" field. It's identical to CIDEQ.
func CID(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Sex applies equality check predicate on the "sex" field. It's identical to SexEQ.
func Sex(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSex), v))
	})
}

// CarIDEQ applies the EQ predicate on the "car_id" field.
func CarIDEQ(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCarID), v))
	})
}

// CarIDNEQ applies the NEQ predicate on the "car_id" field.
func CarIDNEQ(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCarID), v))
	})
}

// CarIDIn applies the In predicate on the "car_id" field.
func CarIDIn(vs ...int) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCarID), v...))
	})
}

// CarIDNotIn applies the NotIn predicate on the "car_id" field.
func CarIDNotIn(vs ...int) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCarID), v...))
	})
}

// CarIDGT applies the GT predicate on the "car_id" field.
func CarIDGT(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCarID), v))
	})
}

// CarIDGTE applies the GTE predicate on the "car_id" field.
func CarIDGTE(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCarID), v))
	})
}

// CarIDLT applies the LT predicate on the "car_id" field.
func CarIDLT(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCarID), v))
	})
}

// CarIDLTE applies the LTE predicate on the "car_id" field.
func CarIDLTE(v int) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCarID), v))
	})
}

// CIDEQ applies the EQ predicate on the "c_id" field.
func CIDEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCID), v))
	})
}

// CIDNEQ applies the NEQ predicate on the "c_id" field.
func CIDNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCID), v))
	})
}

// CIDIn applies the In predicate on the "c_id" field.
func CIDIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCID), v...))
	})
}

// CIDNotIn applies the NotIn predicate on the "c_id" field.
func CIDNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCID), v...))
	})
}

// CIDGT applies the GT predicate on the "c_id" field.
func CIDGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCID), v))
	})
}

// CIDGTE applies the GTE predicate on the "c_id" field.
func CIDGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCID), v))
	})
}

// CIDLT applies the LT predicate on the "c_id" field.
func CIDLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCID), v))
	})
}

// CIDLTE applies the LTE predicate on the "c_id" field.
func CIDLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCID), v))
	})
}

// CIDContains applies the Contains predicate on the "c_id" field.
func CIDContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCID), v))
	})
}

// CIDHasPrefix applies the HasPrefix predicate on the "c_id" field.
func CIDHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCID), v))
	})
}

// CIDHasSuffix applies the HasSuffix predicate on the "c_id" field.
func CIDHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCID), v))
	})
}

// CIDEqualFold applies the EqualFold predicate on the "c_id" field.
func CIDEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCID), v))
	})
}

// CIDContainsFold applies the ContainsFold predicate on the "c_id" field.
func CIDContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SexEQ applies the EQ predicate on the "sex" field.
func SexEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSex), v))
	})
}

// SexNEQ applies the NEQ predicate on the "sex" field.
func SexNEQ(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSex), v))
	})
}

// SexIn applies the In predicate on the "sex" field.
func SexIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSex), v...))
	})
}

// SexNotIn applies the NotIn predicate on the "sex" field.
func SexNotIn(vs ...string) predicate.Student {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Student(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSex), v...))
	})
}

// SexGT applies the GT predicate on the "sex" field.
func SexGT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSex), v))
	})
}

// SexGTE applies the GTE predicate on the "sex" field.
func SexGTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSex), v))
	})
}

// SexLT applies the LT predicate on the "sex" field.
func SexLT(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSex), v))
	})
}

// SexLTE applies the LTE predicate on the "sex" field.
func SexLTE(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSex), v))
	})
}

// SexContains applies the Contains predicate on the "sex" field.
func SexContains(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSex), v))
	})
}

// SexHasPrefix applies the HasPrefix predicate on the "sex" field.
func SexHasPrefix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSex), v))
	})
}

// SexHasSuffix applies the HasSuffix predicate on the "sex" field.
func SexHasSuffix(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSex), v))
	})
}

// SexEqualFold applies the EqualFold predicate on the "sex" field.
func SexEqualFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSex), v))
	})
}

// SexContainsFold applies the ContainsFold predicate on the "sex" field.
func SexContainsFold(v string) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSex), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
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
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(func(s *sql.Selector) {
		p(s.Not())
	})
}