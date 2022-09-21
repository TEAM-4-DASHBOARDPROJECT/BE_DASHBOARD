package mentee

type Core struct {
	ID                uint
	FullName          string
	StatusID          uint
	ClassID           uint
	Category          string
	Address           string
	HomeAddress       string
	Email             string
	Gender            string
	Telegram          string
	Phone             string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	EducationGraduate string
}

type Status struct {
	ID   uint
	Name string
}

type Class struct {
	ID   uint
	Name string
}

type UsecaseInterface interface {
	PostMentee(data Core) (int, error)
	GetMentee(get string) (data []Core, err error)
	PutMentee(id int, newData Core) (row int, err error)
	// DeleteData(token int) (int, error)
}

type DataInterface interface {
	AddMentee(data Core) (int, error)
	SelectMentee(get string) (data []Core, err error)
	UpdateMentee(id int, newData Core) (row int, err error)
	// DeleteByToken(token int) (int, error)
}
