package services

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/safrizalaji16/booking-system-go/domain/repository"
)

type BookingService interface {
	ReadAllBookings(c *gin.Context, src *[]domain.Booking) (err error)
	CreateBooking(c *gin.Context, src *domain.Booking) (err error)
	ReadDetailBooking(c *gin.Context, src *domain.Booking, id int64) (err error)
	DeleteBooking(c *gin.Context, src *domain.Booking, id int64) (err error)
	EditBooking(c *gin.Context, src *domain.Booking, id int64) (err error)
}

type bookingService struct {
	bookingRepo repository.BookingRepository
}

func NewBookingService(bookingRepo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
	}
}

func (srv *bookingService) ReadAllBookings(c *gin.Context, src *[]domain.Booking) (err error) {
	if err = srv.bookingRepo.ReadAllBookings(c, src); err != nil {
		return err
	}

	return
}

func (srv *bookingService) CreateBooking(c *gin.Context, src *domain.Booking) (err error) {
	if err = srv.bookingRepo.CreateBooking(c, src); err != nil {
		return err
	}

	return
}

func (srv *bookingService) ReadDetailBooking(c *gin.Context, src *domain.Booking, id int64) (err error) {
	if err = srv.bookingRepo.ReadDetailBooking(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *bookingService) DeleteBooking(c *gin.Context, src *domain.Booking, id int64) (err error) {
	if err = srv.bookingRepo.DeleteBooking(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *bookingService) EditBooking(c *gin.Context, src *domain.Booking, id int64) (err error) {
	if err = srv.bookingRepo.EditBooking(c, src, id); err != nil {
		return err
	}

	return
}
