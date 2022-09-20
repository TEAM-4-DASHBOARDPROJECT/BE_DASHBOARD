package entity

type ClassEntity struct{
	ClassID		uint
	UserID		uint
	Name		string
	MulaiKelas	string
	AkhirKelas	string
}

type UsecaseClass interface{
	Create(class ClassEntity) (err error)
	Update(class ClassEntity) (err error)
	Delete(class ClassEntity) (err error)
	GetClass() (result []ClassEntity, err error)
}

type RepoClass interface{
	Insert(class ClassEntity) (affectedRow int, err error)
	Update(class ClassEntity) (affectedRow int, err error)
	Delete(class ClassEntity) (affectedRow int, err error)
	FindAll() (result []ClassEntity, err error)
}