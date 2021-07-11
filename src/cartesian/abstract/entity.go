package abstract

type (
	// Entity is pointer to result value.
	Entity   interface{}
	Entities []Entity
)

func (v Entities) Each(f func(Entity)) {
	for _, o := range v {
		f(o)
	}
}

func (v Entities) AsValues() (r Values) {
	for _, e := range v {
		r = append(r, e)
	}
	return r
}
