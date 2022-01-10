package thulium

type P struct {
	*Prototype
	Text	string
}

func (q *P) Render()  {
	q.Inner = q.Text
}