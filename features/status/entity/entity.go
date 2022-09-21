package entity

type StatusEntity struct{
	StatusID	int
	StatusName	string
}

type UsecaseStatus interface{
	GetStatus() (status []StatusEntity, err error)
}

type RepoStatus interface{
	SelectStatus() (status []StatusEntity,err error)
}