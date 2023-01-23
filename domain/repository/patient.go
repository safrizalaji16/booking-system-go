package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/db"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/uptrace/bun"
)

type PatientRepository interface {
	ReadAllPatients(c *gin.Context, src *[]domain.Patient) (err error)
	CreatePatient(c *gin.Context, src *domain.Patient) (err error)
	ReadDetailPatient(c *gin.Context, src *domain.Patient, id int64) (err error)
	DeletePatient(c *gin.Context, src *domain.Patient, id int64) (err error)
	EditPatient(c *gin.Context, src *domain.Patient, id int64) (err error)
}

type patientRepository struct {
	db *bun.DB
}

func NewPatientRepository() PatientRepository {
	return &patientRepository{
		db: db.GetConn(),
	}
}

func (r *patientRepository) ReadAllPatients(c *gin.Context, src *[]domain.Patient) (err error) {
	if err = r.db.NewSelect().Model(src).Scan(c); err != nil {
		return err
	}

	return
}

func (r *patientRepository) CreatePatient(c *gin.Context, src *domain.Patient) (err error) {
	if _, err = r.db.NewInsert().Model(src).Exec(c); err != nil {
		return err
	}

	return
}

func (r *patientRepository) ReadDetailPatient(c *gin.Context, src *domain.Patient, id int64) (err error) {
	if err = r.db.NewSelect().Model(src).Where("id = ?", &id).Scan(c); err != nil {
		return err
	}

	return
}

func (r *patientRepository) DeletePatient(c *gin.Context, src *domain.Patient, id int64) (err error) {
	if _, err = r.db.NewDelete().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}

func (r *patientRepository) EditPatient(c *gin.Context, src *domain.Patient, id int64) (err error) {
	if _, err = r.db.NewUpdate().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}
