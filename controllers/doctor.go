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

type DoctorController interface {
	ReadAllDoctors(c *gin.Context)
	CreateDoctor(c *gin.Context)
	ReadDetailDoctor(c *gin.Context)
	DeleteDoctor(c *gin.Context)
	EditDoctor(c *gin.Context)
}

type doctorController struct {
	doctorSrv services.DoctorService
}

func NewDoctorController(doctorSrv services.DoctorService) DoctorController {
	return &doctorController{
		doctorSrv: doctorSrv,
	}
}

func (ctl *doctorController) ReadAllDoctors(c *gin.Context) {
	var doctors []domain.Doctor

	err := ctl.doctorSrv.ReadAllDoctors(c, &doctors)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": doctors,
	})
}

func (ctl *doctorController) CreateDoctor(c *gin.Context) {
	var src requests.CreateDoctorRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var doctor domain.Doctor

	if err := copier.Copy(&doctor, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	err := ctl.doctorSrv.CreateDoctor(c, &doctor)
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

func (ctl *doctorController) ReadDetailDoctor(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var doctor domain.Doctor

	err = ctl.doctorSrv.ReadDetailDoctor(c, &doctor, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": doctor,
	})
}

func (ctl *doctorController) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var doctor domain.Doctor

	err = ctl.doctorSrv.DeleteDoctor(c, &doctor, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": `Doctor has been deleted`,
	})
}

func (ctl *doctorController) EditDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var src requests.CreateDoctorRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var doctor domain.Doctor

	if err := copier.Copy(&doctor, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	if err = ctl.doctorSrv.EditDoctor(c, &doctor, id); err != nil {
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
