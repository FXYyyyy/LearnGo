package animal

type Pet struct {
	name string
}

func NewPet(n string) Pet {
	return Pet{name: n}
}

/*
func (p Pet) GetName() string {
	return p.name
}
 */
