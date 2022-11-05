package main

import (
	"cosmart/controller"
	"cosmart/entity/model"
	"cosmart/repository"
	"cosmart/usecase"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {

	godotenv.Load(".env")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dsn := fmt.Sprintf(`host=%s user=postgres password=postgres dbname=book_schedule port=5432 sslmode=disable TimeZone=Asia/Shanghai`, os.Getenv("HOST"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DbInit(db)
	if err != nil {
		panic(err)
	}

	rp := repository.New(db)
	uc := usecase.New(rp)
	ctrl := controller.New(uc)

	// Routes
	e.GET("/books", ctrl.GetListOfBook)
	e.POST("/schedule-pickup", ctrl.SchedulePickup)
	e.GET("/schedule-pickup", ctrl.ScheduleList)

	port := ":8080"

	e.Logger.Fatal(e.Start(port))
}

func DbInit(db *gorm.DB) error {

	tableFound := db.Migrator().HasTable(&model.BookSchedule{})

	if !tableFound {
		err := db.Migrator().CreateTable(&model.BookSchedule{})
		if err != nil {
			return err
		}
	}

	return nil
}
