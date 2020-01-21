package entities

var Index int = 0

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

func NewPerson(FirstName string, LastName string) Person {
	Index++
	return Person{Id: Index, FirstName: FirstName, LastName: LastName}
}
