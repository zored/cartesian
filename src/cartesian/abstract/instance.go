package abstract

type (
	// Instance is pointer to result value.
	Instance  interface{}
	Instances []Instance
)

func (v Instances) Each(f func(Instance)) {
	for _, o := range v {
		f(o)
	}
}

func (v Instances) AsValues() (r Values) {
	for _, e := range v {
		r = append(r, e)
	}
	return r
}
