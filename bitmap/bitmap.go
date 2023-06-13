package bitmap

type Container struct {
	Bitmap []uint64 `json:"bitmap"`
}

func NewContainer() *Container {
	return new(Container)
}

func (c *Container) Insert(val int) {
	idx := val / 64
	offset := val % 64

	containerDiff := idx + 1 - len(c.Bitmap)
	for i := 0; i < containerDiff; i++ {
		c.Bitmap = append(c.Bitmap, 0)
	}

	c.Bitmap[idx] = c.Bitmap[idx] | (1 << offset)
}

func (c *Container) GetAll() []int {
	var res []int
	for ci, c := range c.Bitmap {
		for i := 0; i < 64; i++ {
			if (c >> i & 1) == 0 {
				continue
			}
			res = append(res, ci*64+i)
		}
	}
	return res
}

func (c *Container) Copy() *Container {
	n := NewContainer()
	n.Bitmap = make([]uint64, len(c.Bitmap))
	copy(n.Bitmap, c.Bitmap)
	return n
}

func (c *Container) Or(a *Container) {
	if a == nil {
		return
	}
	if len(c.Bitmap) > len(a.Bitmap) {
		for i := 0; i < len(a.Bitmap); i++ {
			c.Bitmap[i] = a.Bitmap[i] | c.Bitmap[i]
		}
		return
	}

	for i := 0; i < len(c.Bitmap); i++ {
		c.Bitmap[i] = a.Bitmap[i] | c.Bitmap[i]
	}
	c.Bitmap = append(c.Bitmap, a.Bitmap[len(c.Bitmap):]...)
}

func (c *Container) And(a *Container) {
	if a == nil {
		return
	}

	nextLen := len(c.Bitmap)
	if len(c.Bitmap) > len(a.Bitmap) {
		nextLen = len(a.Bitmap)
	}

	for i := 0; i < nextLen; i++ {
		c.Bitmap[i] = a.Bitmap[i] & c.Bitmap[i]
	}
	c.Bitmap = c.Bitmap[:nextLen]
}
