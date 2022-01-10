package thulium



type Quark interface {
	Render()
}

type Prototype struct {
	Inner		string
	Children 	[]*Prototype
	Parent		*Prototype
}

func (q *Prototype) T(children ...*Prototype) *Prototype {
	for _,c := range children {
		c.Parent = q
	}
	q.Children = append(q.Children, children...)

	return q
}