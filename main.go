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

	specialities := r.Group("/specialities")
	{
		specialities.GET("", specialityCtl.ReadAllSpecialities)
		specialities.POST("", specialityCtl.CreateSpeciality)
		specialities.GET(":id", specialityCtl.ReadDetailSpeciality)
		specialities.DELETE(":id", specialityCtl.DeleteSpeciality)
		specialities.PUT(":id", specialityCtl.EditSpeciality)
	}

	r.Run()
}
