package repository

type repositoryPool struct {
	Speciality SpecialityRepository
	Doctor     DoctorRepository
	Patient    PatientRepository
	Booking    BookingRepository
}

func InitRepositoryInstance() *repositoryPool {
	return &repositoryPool{
		Speciality: NewSpecialityRepository(),
		Doctor:     NewDoctorRepository(),
		Patient:    NewPatientRepository(),
		Booking:    NewBookingRepository(),
	}
}
