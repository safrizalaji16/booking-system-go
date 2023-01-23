package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/db"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/uptrace/bun"
)

type DoctorRepository interface {
	ReadAllDoctors(c *gin.Context, src *[]domain.Doctor) (err error)
	CreateDoctor(c *gin.Context, src *domain.Doctor) (err error)
	ReadDetailDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error)
	DeleteDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error)
	EditDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error)
}

type doctorRepository struct {
	db *bun.DB
}

func NewDoctorRepository() DoctorRepository {
	return &doctorRepository{
		db: db.GetConn(),
	}
}

func (r *doctorRepository) ReadAllDoctors(c *gin.Context, src *[]domain.Doctor) (err error) {
	if err = r.db.NewSelect().Model(src).Scan(c); err != nil {
		return err
	}

	return
}

func (r *doctorRepository) CreateDoctor(c *gin.Context, src *domain.Doctor) (err error) {
	if _, err = r.db.NewInsert().Model(src).Exec(c); err != nil {
		return err
	}

	return
}

func (r *doctorRepository) ReadDetailDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error) {
	if err = r.db.NewSelect().Model(src).Where("id = ?", &id).Scan(c); err != nil {
		return err
	}

	return
}

func (r *doctorRepository) DeleteDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error) {
	if _, err = r.db.NewDelete().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}

func (r *doctorRepository) EditDoctor(c *gin.Context, src *domain.Doctor, id int64) (err error) {
	if _, err = r.db.NewUpdate().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}
