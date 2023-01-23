package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/db"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/uptrace/bun"
)

type BookingRepository interface {
	ReadAllBookings(c *gin.Context, src *[]domain.Booking) (err error)
	CreateBooking(c *gin.Context, src *domain.Booking) (err error)
	ReadDetailBooking(c *gin.Context, src *domain.Booking, id int64) (err error)
	DeleteBooking(c *gin.Context, src *domain.Booking, id int64) (err error)
	EditBooking(c *gin.Context, src *domain.Booking, id int64) (err error)
}

type bookingRepository struct {
	db *bun.DB
}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{
		db: db.GetConn(),
	}
}

func (r *bookingRepository) ReadAllBookings(c *gin.Context, src *[]domain.Booking) (err error) {
	if err = r.db.NewSelect().Model(src).Scan(c); err != nil {
		return err
	}

	return
}

func (r *bookingRepository) CreateBooking(c *gin.Context, src *domain.Booking) (err error) {
	if _, err = r.db.NewInsert().Model(src).Exec(c); err != nil {
		return err
	}

	return
}

func (r *bookingRepository) ReadDetailBooking(c *gin.Context, src *domain.Booking, id int64) (err error) {
	if err = r.db.NewSelect().Model(src).Where("id = ?", &id).Scan(c); err != nil {
		return err
	}

	return
}

func (r *bookingRepository) DeleteBooking(c *gin.Context, src *domain.Booking, id int64) (err error) {
	if _, err = r.db.NewDelete().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}

func (r *bookingRepository) EditBooking(c *gin.Context, src *domain.Booking, id int64) (err error) {
	if _, err = r.db.NewUpdate().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}
