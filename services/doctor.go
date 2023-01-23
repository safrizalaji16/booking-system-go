package services

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/safrizalaji16/booking-system-go/domain/repository"
)

type DoctorService interface {
	ReadAllDoctors(c *gin.Context, src *[]domain.Doctor) (err error)
	CreateDoctor(c *gin.Context, src *domain.Doctor) (err error)
	ReadDetailDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error)
	DeleteDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error)
	EditDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error)
}

type doctorService struct {
	doctorRepo repository.DoctorRepository
}

func NewDoctorService(doctorRepo repository.DoctorRepository) DoctorService {
	return &doctorService{
		doctorRepo: doctorRepo,
	}
}

func (srv *doctorService) ReadAllDoctors(c *gin.Context, src *[]domain.Doctor) (err error) {
	if err = srv.doctorRepo.ReadAllDoctors(c, src); err != nil {
		return err
	}

	return
}

func (srv *doctorService) CreateDoctor(c *gin.Context, src *domain.Doctor) (err error) {
	if err = srv.doctorRepo.CreateDoctor(c, src); err != nil {
		return err
	}

	return
}

func (srv *doctorService) ReadDetailDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error) {
	if err = srv.doctorRepo.ReadDetailDoctor(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *doctorService) DeleteDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error) {
	if err = srv.doctorRepo.DeleteDoctor(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *doctorService) EditDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error) {
	if err = srv.doctorRepo.EditDoctor(c, src, id); err != nil {
		return err
	}

	return
}
