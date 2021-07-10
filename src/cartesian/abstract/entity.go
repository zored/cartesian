package abstract

type (
	Entity   interface{}
	Entities []Entity
)

func (v Entities) EachEntity(f func(Entity)) {
	for _, o := range v {
		f(o)
	}
}
