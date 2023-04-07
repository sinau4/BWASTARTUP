package main

import (
	"BWASTARTUP/handler"
	"BWASTARTUP/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	// user := user.User{
	// 	Name: "Test simpan",
	// }
	// userRepository.Save(user)
	userService := user.NewService(userRepository)
	// userInput := user.RegisterUserInput{
	// 	Name:       "Test dari Service",
	// 	Occupation: "anak band",
	// 	Email:      "contoh@example.com",
	// 	Password:   "password",
	// }
	// userService.RegisterUser(userInput)

	// userByEmail, err := userRepository.FindByEmail("aa@emal.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// if userByEmail.ID == 0 {
	// 	fmt.Println("user tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

	// input := user.LoginInput{
	// 	Email:    "aa@emal.com",
	// 	Password: "passworda",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// userService.SaveAvatar(1, "images/1-profile.png")

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()

	// fmt.Println("Connection success")

	// var users []user.User
	// fmt.Println(len(users))

	// db.Find(&users)
	// fmt.Println(len(users))

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=====")
	// }

	// r := gin.Default()
	// r.GET("/", handler)
	// r.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	fmt.Println("Connection success")

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
