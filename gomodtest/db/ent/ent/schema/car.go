package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("register_at"),
	}
}

// Edges of the Car.
// 假设我们有一个汽车对象，我们希望获得其所有者，即这辆车属于的用户。
// 为此，我们有另一种称为“Inverse Edge”的逆向查询
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// 创建一个名为 owner 的 "Inverse Edge"，它的类型是 User
		//
		edge.From("owner", User.Type).
			Ref("cars").
			Unique(),
	}
}
