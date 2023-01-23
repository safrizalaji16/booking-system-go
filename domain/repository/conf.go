package repository

type repositoryPool struct {
	Speciality SpecialityRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		Speciality: NewSpecialityRepository(),
	}
}
