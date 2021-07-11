package config

type Context struct {
	completeIOs IOs
}

func (c *Context) AddCompleteIO(io IO) {
	c.completeIOs = append(c.completeIOs, io)
}

func (c *Context) EachCompleteIO(ok func(v IO) (stop bool)) {
	for _, o := range c.completeIOs {
		if ok(o) {
			return
		}
	}
}

func (c *Context) GetLastCompleteIO() IO {
	i := len(c.completeIOs) - 1
	if i < 0 {
		return nil
	}
	return c.completeIOs[i]
}
