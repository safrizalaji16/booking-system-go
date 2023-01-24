package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/safrizalaji16/booking-system-go/domain"
	"github.com/safrizalaji16/booking-system-go/dto/requests"
	"github.com/safrizalaji16/booking-system-go/helpers"
	"github.com/safrizalaji16/booking-system-go/services"
)

type PatientController interface {
	ReadAllPatients(c *gin.Context)
	CreatePatient(c *gin.Context)
	ReadDetailPatient(c *gin.Context)
	DeletePatient(c *gin.Context)
	EditPatient(c *gin.Context)
}

type patientController struct {
	patientSrv services.PatientService
}

func NewPatientController(patientSrv services.PatientService) PatientController {
	return &patientController{
		patientSrv: patientSrv,
	}
}

func (ctl *patientController) ReadAllPatients(c *gin.Context) {
	var patients []domain.Patient

	err := ctl.patientSrv.ReadAllPatients(c, &patients)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": patients,
	})
}

func (ctl *patientController) CreatePatient(c *gin.Context) {
	var src requests.CreateDoctorRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var patient domain.Patient

	if err := copier.Copy(&patient, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	patient.Password, _ = helpers.HashPassword(patient.Password)

	err := ctl.patientSrv.CreatePatient(c, &patient)
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

func (ctl *patientController) ReadDetailPatient(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var patient domain.Patient

	err = ctl.patientSrv.ReadDetailPatient(c, &patient, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": patient,
	})
}

func (ctl *patientController) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var patient domain.Patient

	err = ctl.patientSrv.DeletePatient(c, &patient, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": `Patient has been deleted`,
	})
}

func (ctl *patientController) EditPatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var src requests.CreatePatientRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var patient domain.Patient

	if err := copier.Copy(&patient, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	if err = ctl.patientSrv.EditPatient(c, &patient, id); err != nil {
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
