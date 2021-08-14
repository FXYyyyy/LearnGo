package interfacerealize

import "fmt"

type Integer int

/*
func (i Integer) Add(x Integer) (ret Integer) {
	ret = i+x
	i++
	fmt.Println("realize, i = ", i)
	return
}
 */

func (i *Integer) Add(x Integer) (ret Integer) {
	ret = (*i)+x
	*i++
	fmt.Println("realize, i = ", *i)
	return
}

func (i Integer) Multiply(x Integer) Integer {
	return i*x
}

func (i Integer) Subtract(x Integer) Integer {
	return i-x
}

func (i Integer) PrintI() {
	fmt.Println("getI = ", i)
}
