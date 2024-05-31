package config

import (
	"library/app/controller/member"
	"library/app/controller/book"
	"library/app/controller/root"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {
	rootController := root.NewController()
	router := gin.Default()
	router.Use(corsConfig)
	router.GET("/", rootController.Index)
	v1 := router.Group("v1")
	membersGroup := v1.Group("members")
	{
		memberController := member.NewController(db)
		membersGroup.GET("list", memberController.List)
	}
	booksGroup := v1.Group("books")
	{
		bookController := book.NewController(db)
		booksGroup.GET("list", bookController.List)		
		booksGroup.POST("transaction", bookController.Order)		
	}
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
