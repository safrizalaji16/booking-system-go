package main

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/booking-system-go/controllers"
	"github.com/safrizalaji16/booking-system-go/db"
	"github.com/safrizalaji16/booking-system-go/domain/repository"
	"github.com/safrizalaji16/booking-system-go/services"
)

func main() {
	r := gin.Default()

	dbInstance := db.InitDB()
	defer dbInstance.Close()

	rp := repository.InitRepositoryInstance()

	specialitySrv := services.NewSpecialityService(rp.Speciality)
	specialityCtl := controllers.NewSpecialityController(specialitySrv)

	doctorSrv := services.NewDoctorService(rp.Doctor)
	doctorCtl := controllers.NewDoctorController(doctorSrv)

	patientSrv := services.NewPatientService(rp.Patient)
	patientCtl := controllers.NewPatientController(patientSrv)

	bookingSrv := services.NewBookingService(rp.Booking)
	bookingCtl := controllers.NewBookingController(bookingSrv)

	specialities := r.Group("/specialities")
	{
		specialities.GET("", specialityCtl.ReadAllSpecialities)
		specialities.POST("", specialityCtl.CreateSpeciality)
		specialities.GET(":id", specialityCtl.ReadDetailSpeciality)
		specialities.DELETE(":id", specialityCtl.DeleteSpeciality)
		specialities.PUT(":id", specialityCtl.EditSpeciality)
	}

	doctors := r.Group("/doctors")
	{
		doctors.GET("", doctorCtl.ReadAllDoctors)
		doctors.POST("", doctorCtl.CreateDoctor)
		doctors.GET(":id", doctorCtl.ReadDetailDoctor)
		doctors.DELETE(":id", doctorCtl.DeleteDoctor)
		doctors.PUT(":id", doctorCtl.EditDoctor)
	}

	patients := r.Group("/patients")
	{
		patients.GET("", patientCtl.ReadAllPatients)
		patients.POST("", patientCtl.CreatePatient)
		patients.GET(":id", patientCtl.ReadDetailPatient)
		patients.DELETE(":id", patientCtl.DeletePatient)
		patients.PUT(":id", patientCtl.EditPatient)
	}

	bookings := r.Group("/bookings")
	{
		bookings.GET("", bookingCtl.ReadAllBookings)
		bookings.POST("", bookingCtl.CreateBooking)
		bookings.GET(":id", bookingCtl.ReadDetailBooking)
		bookings.DELETE(":id", bookingCtl.DeleteBooking)
		bookings.PUT(":id", bookingCtl.EditBooking)
	}

	r.Run()
}
