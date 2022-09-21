package users

type Core struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Team     string
	Role     string
}

type DataInterface interface {
	GetMyProfile(token int) (Core, error)
	SelectAll(page, token int) ([]Core, error)
	UpdateData(data Core) int
	DelData(id int) int
	InsertData(data Core) int
}

type UsecaseInterface interface {
	SelectMe(token int) (Core, error)
	GetAll(page, token int) ([]Core, error)
	PutData(data Core) int
	DeleteData(id int) int
	PostData(data Core) int
}
