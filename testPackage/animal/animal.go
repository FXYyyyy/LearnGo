package animal

type Animal struct {
	name string
}

func NewAnimal(n string) Animal {
	return Animal{name: n}
}

func (a Animal) GetName() string {
	return a.name
}

func (a Animal) Call() string {
	return "这是一个动物的叫声：==="
}

func (a Animal) FavorFood() string {
	return "最爱吃的东西：==="
}
