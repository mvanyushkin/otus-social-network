package domain

type Profile struct {
	Id        uint64
	Email     string
	FirstName string
	LastName  string
	Age       uint8
	Gender    uint8
	City      string
	Hobby     string
}
