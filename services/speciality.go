package services

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/safrizalaji16/booking-system-go/domain/repository"
)

type SpecialityService interface {
	ReadAllSpecialities(c *gin.Context, src *[]domain.Speciality) (err error)
	CreateSpeciality(c *gin.Context, src *domain.Speciality) (err error)
	ReadDetailSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error)
	DeleteSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error)
	EditSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error)
}

type specialityService struct {
	specialityRepo repository.SpecialityRepository
}

func NewSpecialityService(specialityRepo repository.SpecialityRepository) SpecialityService {
	return &specialityService{
		specialityRepo: specialityRepo,
	}
}

func (srv *specialityService) ReadAllSpecialities(c *gin.Context, src *[]domain.Speciality) (err error) {
	if err = srv.specialityRepo.ReadAllSpecialities(c, src); err != nil {
		return err
	}

	return
}

func (srv *specialityService) CreateSpeciality(c *gin.Context, src *domain.Speciality) (err error) {
	if err = srv.specialityRepo.CreateSpeciality(c, src); err != nil {
		return err
	}

	return
}

func (srv *specialityService) ReadDetailSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error) {
	if err = srv.specialityRepo.ReadDetailSpeciality(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *specialityService) DeleteSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error) {
	if err = srv.specialityRepo.DeleteSpeciality(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *specialityService) EditSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error) {
	if err = srv.specialityRepo.EditSpeciality(c, src, id); err != nil {
		return err
	}

	return
}
