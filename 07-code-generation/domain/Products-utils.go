package domain

func (p *Products) Print() {
	for _, i := range *p {
		i.Print()
	}
}
