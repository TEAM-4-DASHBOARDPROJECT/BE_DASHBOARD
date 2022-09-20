package mentee

type Core struct {
	ID                int
	FullName          string
	Status            string
	Class             string
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

type UsecaseInterface interface {
	PostMentee(data Core) (int, error)
	// GetByToken(token int) (data Core, err error)
	// PutData(newData Core) (row int, err error)
	// DeleteData(token int) (int, error)
}

type DataInterface interface {
	AddMentee(data Core) (int, error)
	// SelectByToken(token int) (data Core, err error)
	// UpdateData(newData Core) (row int, err error)
	// DeleteByToken(token int) (int, error)
}
