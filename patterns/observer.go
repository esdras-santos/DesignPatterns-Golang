package patterns


type IObserver interface{
	Update()
}

//implementation of the IObserver interface 
type Observer struct{
	//variable for time
}

//receive the time variable as a parameter
func (obs *Observer) Update(){
	//update the time variable
}



type IObservable interface{
	Add(IObserver)
	Remove(IObserver)
	Notify()
}

//implementation of the IObservable interface 
type Observable struct{
	Observers []IObserver
}

func (obs *Observable) Add(iobs IObserver){
	obs.Observers = append(obs.Observers, iobs)
}

func (obs *Observable) Remove(iobs IObserver){
	var j int
	for i, e := range obs.Observers {
		if iobs == e {
			j = i
			break
		}
	}
	obs.Observers = append(obs.Observers[:j], obs.Observers[j+1:]...)
}

func (obs *Observable) Notify(){
	for _,observer := range obs.Observers{
		observer.Update()
	}
}