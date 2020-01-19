package entities

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

func NewPerson(Id int, FirstName string, LastName string) Person {
	return Person{Id: Id, FirstName: FirstName, LastName: LastName}
}
