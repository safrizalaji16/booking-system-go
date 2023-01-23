package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/db"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/uptrace/bun"
)

type SpecialityRepository interface {
	ReadAllSpecialities(c *gin.Context, src *[]domain.Speciality) (err error)
	CreateSpeciality(c *gin.Context, src *domain.Speciality) (err error)
	ReadDetailSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error)
	DeleteSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error)
	EditSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error)
}

type specialityRepository struct {
	db *bun.DB
}

func NewSpecialityRepository() SpecialityRepository {
	return &specialityRepository{
		db: db.GetConn(),
	}
}

func (r *specialityRepository) ReadAllSpecialities(c *gin.Context, src *[]domain.Speciality) (err error) {
	if err = r.db.NewSelect().Model(src).Scan(c); err != nil {
		return err
	}

	return
}

func (r *specialityRepository) CreateSpeciality(c *gin.Context, src *domain.Speciality) (err error) {
	if _, err = r.db.NewInsert().Model(src).Exec(c); err != nil {
		return err
	}

	return
}

func (r *specialityRepository) ReadDetailSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error) {
	if err = r.db.NewSelect().Model(src).Where("id = ?", &id).Scan(c); err != nil {
		return err
	}

	return
}

func (r *specialityRepository) DeleteSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error) {
	if _, err = r.db.NewDelete().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}

func (r *specialityRepository) EditSpeciality(c *gin.Context, src *domain.Speciality, id int64) (err error) {
	if _, err = r.db.NewUpdate().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}
