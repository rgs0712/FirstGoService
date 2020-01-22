package entities

var Index int = 0

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

func NewPerson(FirstName string, LastName string) Person {
	return Person{Id: GetPersonId(), FirstName: FirstName, LastName: LastName}
}
func GetPersonId() int {
	Index++
	return Index
}
