package client

type ClientErr struct {
	Msg string
}

func (c *ClientErr) Error() string {
	return c.Msg
}
