package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/safrizalaji16/booking-system-go/dto/requests"
	"github.com/safrizalaji16/booking-system-go/services"
)

type BookingController interface {
	ReadAllBookings(c *gin.Context)
	CreateBooking(c *gin.Context)
	ReadDetailBooking(c *gin.Context)
	DeleteBooking(c *gin.Context)
	EditBooking(c *gin.Context)
}

type bookingController struct {
	bookingSrv services.BookingService
}

func NewBookingController(bookingSrv services.BookingService) BookingController {
	return &bookingController{
		bookingSrv: bookingSrv,
	}
}

func (ctl *bookingController) ReadAllBookings(c *gin.Context) {
	var bookings []domain.Booking

	err := ctl.bookingSrv.ReadAllBookings(c, &bookings)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

func (ctl *bookingController) CreateBooking(c *gin.Context) {
	var src requests.CreateBookingRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var booking domain.Booking

	if err := copier.Copy(&booking, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	err := ctl.bookingSrv.CreateBooking(c, &booking)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": src,
	})
}

func (ctl *bookingController) ReadDetailBooking(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var booking domain.Booking

	err = ctl.bookingSrv.ReadDetailBooking(c, &booking, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": booking,
	})
}

func (ctl *bookingController) DeleteBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var booking domain.Booking

	err = ctl.bookingSrv.DeleteBooking(c, &booking, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": `Booking has been deleted`,
	})
}

func (ctl *bookingController) EditBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var src requests.CreateBookingRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var booking domain.Booking

	if err := copier.Copy(&booking, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	if err = ctl.bookingSrv.EditBooking(c, &booking, id); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": src,
	})
}
