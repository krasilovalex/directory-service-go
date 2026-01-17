package main

import (
	"context"
	_ "directory-service/docs"
	"directory-service/internal/delivery/http"
	"directory-service/internal/repository/postgres"
	"directory-service/internal/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	// swagger embed files
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// @title Swagger Example API
// @version 1.0
// @description Directory Service Golang
// @termsOfService http://swagger.io/terms/

// @contact.name  Alex Wayzzoo
// @contact.url https/t.me/wayzzoo

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден, берем переменные из системы")
	}

	connString := os.Getenv("DB_URL")

	if connString == "" {
		log.Fatal("ОШИБКА: Не задан DB_URL в .env ")
	}

	ctx := context.Background()
	dbPool, err := pgxpool.New(ctx, connString)

	if err != nil {
		log.Fatal("Не могу подключиться к базе", err)
	}

	defer dbPool.Close()

	deptRepo := postgres.NewDepartmentRepository(dbPool)
	deptUseCase := usecase.NewDepartmentUseCase(deptRepo)
	deptHandler := http.NewDepartmentHandler(deptUseCase)

	locRepo := postgres.NewLocationRepository(dbPool)
	locUseCase := usecase.NewLocationUseCase(locRepo)
	locHandler := http.NewLocationHandler(locUseCase)

	posRepo := postgres.NewPositionRepository(dbPool)
	posUseCase := usecase.NewPositionUseCase(posRepo)
	posHandler := http.NewPositionHandler(posUseCase)

	r := gin.Default()

	r.POST("/departments", deptHandler.Create)
	r.GET("/departments/:id", deptHandler.GetByID)
	r.GET("/departments", deptHandler.GetAll)
	r.DELETE("/departments/:id", deptHandler.Delete)
	r.PUT("/departments/:id", deptHandler.Update)

	r.POST("/locations", locHandler.Create)
	r.GET("/locations/:id", locHandler.GetByID)
	r.GET("/locations", locHandler.GetAll)
	r.DELETE("/locations/:id", locHandler.Delete)
	r.PUT("/locations/:id", locHandler.Update)

	r.POST("/positions", posHandler.Create)
	r.GET("/positions/:id", posHandler.GetByID)
	r.GET("/positions", posHandler.GetAll)
	r.DELETE("/positions/:id", posHandler.Delete)
	r.PUT("/positions/:id", posHandler.Update)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.LoadHTMLGlob("web/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")

}
