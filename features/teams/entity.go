package teams

import "time"

type Core struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
	User      User
}

type User struct {
	ID   uint
	Name string
}

type RepoTeam interface {
	Insert(data Core) (affectedRow int, err error)
	Update(data Core) (affectedRow int, err error)
	Delete(data Core) (affectedRow int, err error)
	FindAll() (result []Core, err error)
}
type UsecaseTeam interface {
	Create(data Core) (err error)
	Update(data Core) (err error)
	Delete(data Core) (err error)
	GetTeam() (result []Core, err error)
}
