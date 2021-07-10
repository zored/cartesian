package abstract

type (
	// Entity is pointer to result value.
	Entity   interface{}
	Entities []Entity
)

func (v Entities) EachEntity(f func(Entity)) {
	for _, o := range v {
		f(o)
	}
}
