package patterns

import "fmt"

type InterfaceClass interface {
	DogName() string
}

type AbstractClass struct {
	IC   InterfaceClass
	Show func() string
}

//////////////////////////////////////////////////////////////////////////////////
//IMPLEMENTATION
//////////////////////////////////////////////////////////////////////////////////

//abstract class(struct) implementation
type View struct {
	//inherite the Abstract class(struct)
	AbstractClass
}

func InitView(ic InterfaceClass) View {

	v := View{}
	v.IC = ic
	//abstract class(struct) method(function)
	v.Show = func() string {
		return "do you take " + v.IC.DogName() + " or leave " + v.IC.DogName() + "?"
	}

	return v
}

//class(struct) that implements the InterfaceClass
type Resource struct {
	dogName string
}

//interface method(function)
func (r Resource) DogName() string {
	return r.dogName
}

func InitResource(dogname string) Resource {
	r := Resource{}
	r.dogName = dogname
	InitView(r)
	return r
}

func Bridge() {
	resource := InitResource("nabunda")
	view := InitView(resource)
	fmt.Println(view.Show())
}