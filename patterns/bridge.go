package patterns

type InterfaceClass interface{
	DogName() string
}

type AbstractClass struct{
	IC InterfaceClass
	Show func() string
}

//////////////////////////////////////////////////////////////////////////////////
//IMPLEMENTATION
//////////////////////////////////////////////////////////////////////////////////

//abstract class implementation
type View struct{
	AbstractClass
}

func InitView(ic InterfaceClass) View{
	
	v := View{}
	v.IC = ic
	//abstract class method
	v.Show = func() string{
		return "do you take "+v.IC.DogName()+" or leave "+v.IC.DogName()+"?"
	}

	return v
}

//class that implements the InterfaceClass
type Resource struct{
	dogName string
}

//interface method
func (r Resource) DogName() string{
	return r.dogName
}

func InitResource(dogname string) Resource{
	r := Resource{}
	r.dogName = dogname
	InitView(r)
	return r
} 
