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

type SpecialityController interface {
	ReadAllSpecialities(c *gin.Context)
	CreateSpeciality(c *gin.Context)
	ReadDetailSpeciality(c *gin.Context)
	DeleteSpeciality(c *gin.Context)
	EditSpeciality(c *gin.Context)
}

type specialityController struct {
	specialitySrv services.SpecialityService
}

func NewSpecialityController(specialitySrv services.SpecialityService) SpecialityController {
	return &specialityController{
		specialitySrv: specialitySrv,
	}
}

func (ctl *specialityController) ReadAllSpecialities(c *gin.Context) {
	var specialities []domain.Speciality

	err := ctl.specialitySrv.ReadAllSpecialities(c, &specialities)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": specialities,
	})
}

func (ctl *specialityController) CreateSpeciality(c *gin.Context) {
	var src requests.CreateSpecialityRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var speciality domain.Speciality

	if err := copier.Copy(&speciality, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	err := ctl.specialitySrv.CreateSpeciality(c, &speciality)
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

func (ctl *specialityController) ReadDetailSpeciality(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var speciality domain.Speciality

	err = ctl.specialitySrv.ReadDetailSpeciality(c, &speciality, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": speciality,
	})
}

func (ctl *specialityController) DeleteSpeciality(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var speciality domain.Speciality

	err = ctl.specialitySrv.DeleteSpeciality(c, &speciality, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": `Speciality has been deleted`,
	})
}

func (ctl *specialityController) EditSpeciality(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var src requests.CreateSpecialityRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var speciality domain.Speciality

	if err := copier.Copy(&speciality, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	if err = ctl.specialitySrv.EditSpeciality(c, &speciality, id); err != nil {
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
