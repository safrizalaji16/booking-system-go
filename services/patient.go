package services

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/safrizalaji16/booking-system-go/domain/repository"
)

type PatientService interface {
	ReadAllPatients(c *gin.Context, src *[]domain.Patient) (err error)
	CreatePatient(c *gin.Context, src *domain.Patient) (err error)
	ReadDetailPatient(c *gin.Context, src *domain.Patient, id int64) (err error)
	DeletePatient(c *gin.Context, src *domain.Patient, id int64) (err error)
	EditPatient(c *gin.Context, src *domain.Patient, id int64) (err error)
}

type patientService struct {
	patientRepo repository.PatientRepository
}

func NewPatientService(patientRepo repository.PatientRepository) PatientService {
	return &patientService{
		patientRepo: patientRepo,
	}
}

func (srv *patientService) ReadAllPatients(c *gin.Context, src *[]domain.Patient) (err error) {
	if err = srv.patientRepo.ReadAllPatients(c, src); err != nil {
		return err
	}

	return
}

func (srv *patientService) CreatePatient(c *gin.Context, src *domain.Patient) (err error) {
	if err = srv.patientRepo.CreatePatient(c, src); err != nil {
		return err
	}

	return
}

func (srv *patientService) ReadDetailPatient(c *gin.Context, src *domain.Patient, id int64) (err error) {
	if err = srv.patientRepo.ReadDetailPatient(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *patientService) DeletePatient(c *gin.Context, src *domain.Patient, id int64) (err error) {
	if err = srv.patientRepo.DeletePatient(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *patientService) EditPatient(c *gin.Context, src *domain.Patient, id int64) (err error) {
	if err = srv.patientRepo.EditPatient(c, src, id); err != nil {
		return err
	}

	return
}
