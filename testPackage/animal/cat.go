package animal

type Cat struct {
	Animal
	Pet
}

func NewCat(a Animal, p Pet) Cat {
	return Cat{a, p}
}

func (c *Cat) GetName() string {
	return c.Animal.GetName() + c.Pet.GetName()
}


func (c Cat) Call() string {
	return "喵喵喵"
}

func (c Cat) FavorFood() string {
	return "Fish"
}
