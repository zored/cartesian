package configs

type (
	IOs []IO
)

func (o IOs) Last() IO {
	i := len(o) - 1
	if i < 0 {
		return nil
	}
	return o[i]
}

func (o IOs) WithParentIO(parent IO) IOs {
	for _, io := range o {
		io.SetParentIO(parent)
	}
	return o
}