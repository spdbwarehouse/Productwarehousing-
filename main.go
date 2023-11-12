package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"wareHouse/controllers"
	"wareHouse/repository"
	"wareHouse/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	DB_PORT  = flag.String("dbport", "3306", "database port")
	DB_HOST  = flag.String("dbhost", "localhost", "db server address or domain")
	DB_NAME  = flag.String("dbname", "warehouse", "db schema name where tables present")
	DB_USER  = flag.String("dbusername", "root", "db username")
	DB_PASS  = flag.String("dbpassword", "root", "db password")
	APP_PORT = flag.String("appport", "8080", "application port")
)

func main() {
	flag.Parse()
	server := gin.Default()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", *DB_USER, *DB_PASS, *DB_HOST, *DB_PORT, *DB_NAME)
	db, err := gorm.Open(dsn)
	if err != nil {
		panic("failed to connect to database")
	}
	userrepo := repository.NewUserRepository(db)
	userService := service.NewService(userrepo)
	userController := controllers.NewController(userService)
	userController.SetupRoutes(server)

	productrepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productrepo)
	productController := controllers.NewProductController(productService)
	productController.SetupRoutes(server)

	go func() {
		if err := server.Run(":" + (*APP_PORT)); err != nil {
			fmt.Println("error running the ware house server", err)
		}
	}()
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	<-killSignal
	fmt.Println("Gracefully shutting down the server...")
	err = db.Close()
	if err != nil {
		fmt.Println("error closing the db connection", err)
	}
	fmt.Println("Server stopped")
}
